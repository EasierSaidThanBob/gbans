package srcds

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"strings"
	"time"

	"github.com/leighmacdonald/gbans/internal/discord"
	"github.com/leighmacdonald/gbans/internal/domain"
	"github.com/leighmacdonald/gbans/internal/thirdparty"
	"github.com/leighmacdonald/gbans/pkg/log"
	"github.com/leighmacdonald/steamid/v4/steamid"
	"golang.org/x/exp/slices"
)

type srcdsUsecase struct {
	configUsecase   domain.ConfigUsecase
	serversUsecase  domain.ServersUsecase
	srcdsRepository domain.SRCDSRepository
	personUsecase   domain.PersonUsecase
	reportUsecase   domain.ReportUsecase
	discordUsecase  domain.DiscordUsecase
	cookie          string
}

func (h srcdsUsecase) GetOverride(ctx context.Context, overrideID int) (domain.SMOverrides, error) {
	return h.srcdsRepository.GetOverride(ctx, overrideID)
}

func NewSrcdsUsecase(srcdsRepository domain.SRCDSRepository, configUsecase domain.ConfigUsecase, serversUsecase domain.ServersUsecase,
	personUsecase domain.PersonUsecase, reportUsecase domain.ReportUsecase, discordUsecase domain.DiscordUsecase,
) domain.SRCDSUsecase {
	return &srcdsUsecase{
		configUsecase:   configUsecase,
		serversUsecase:  serversUsecase,
		personUsecase:   personUsecase,
		reportUsecase:   reportUsecase,
		discordUsecase:  discordUsecase,
		srcdsRepository: srcdsRepository,
		cookie:          configUsecase.Config().HTTP.CookieKey,
	}
}

func (h srcdsUsecase) GroupOverrides(ctx context.Context, groupID int) ([]domain.SMGroupOverrides, error) {
	group, errGroup := h.GetGroupByID(ctx, groupID)
	if errGroup != nil {
		return []domain.SMGroupOverrides{}, errGroup
	}

	return h.srcdsRepository.GroupOverrides(ctx, group)
}

func (h srcdsUsecase) Overrides(ctx context.Context) ([]domain.SMOverrides, error) {
	return h.srcdsRepository.Overrides(ctx)
}

func (h srcdsUsecase) SaveOverride(ctx context.Context, override domain.SMOverrides) (domain.SMOverrides, error) {
	if override.Name == "" || override.Flags == "" || override.Type != domain.OverrideTypeCommand && override.Type != domain.OverrideTypeGroup {
		return domain.SMOverrides{}, domain.ErrInvalidParameter
	}

	return h.srcdsRepository.SaveOverride(ctx, override)
}

func (h srcdsUsecase) AddOverride(ctx context.Context, name string, overrideType domain.OverrideType, flags string) (domain.SMOverrides, error) {
	if name == "" || flags == "" || overrideType != domain.OverrideTypeCommand && overrideType != domain.OverrideTypeGroup {
		return domain.SMOverrides{}, domain.ErrInvalidParameter
	}

	now := time.Now()

	return h.srcdsRepository.AddOverride(ctx, domain.SMOverrides{
		Type:  overrideType,
		Name:  name,
		Flags: flags,
		TimeStamped: domain.TimeStamped{
			CreatedOn: now,
			UpdatedOn: now,
		},
	})
}

func (h srcdsUsecase) DelOverride(ctx context.Context, overrideID int) error {
	override, errOverride := h.srcdsRepository.GetOverride(ctx, overrideID)
	if errOverride != nil {
		return errOverride
	}

	return h.srcdsRepository.DelOverride(ctx, override)
}

func (h srcdsUsecase) DelAdminGroup(ctx context.Context, adminID int, groupID int) (domain.SMAdmin, error) {
	admin, errAdmin := h.GetAdminByID(ctx, adminID)
	if errAdmin != nil {
		return domain.SMAdmin{}, errAdmin
	}

	group, errGroup := h.GetGroupByID(ctx, groupID)
	if errGroup != nil {
		return domain.SMAdmin{}, errGroup
	}

	existing, errExisting := h.GetAdminGroups(ctx, admin)
	if errExisting != nil && !errors.Is(errExisting, domain.ErrNoResult) {
		return admin, errExisting
	}

	if !slices.Contains(existing, group) {
		return admin, domain.ErrSMAdminGroupExists
	}

	if err := h.srcdsRepository.DeleteAdminGroup(ctx, admin, group); err != nil {
		return domain.SMAdmin{}, err
	}

	admin.Groups = slices.DeleteFunc(admin.Groups, func(g domain.SMGroups) bool {
		return g.GroupID == groupID
	})

	return admin, nil
}

func (h srcdsUsecase) AddAdminGroup(ctx context.Context, adminID int, groupID int) (domain.SMAdmin, error) {
	admin, errAdmin := h.GetAdminByID(ctx, adminID)
	if errAdmin != nil {
		return domain.SMAdmin{}, errAdmin
	}

	group, errGroup := h.GetGroupByID(ctx, groupID)
	if errGroup != nil {
		return domain.SMAdmin{}, errGroup
	}

	existing, errExisting := h.GetAdminGroups(ctx, admin)
	if errExisting != nil && !errors.Is(errExisting, domain.ErrNoResult) {
		return admin, errExisting
	}

	if slices.Contains(existing, group) {
		return admin, domain.ErrSMAdminGroupExists
	}

	if err := h.srcdsRepository.InsertAdminGroup(ctx, admin, group, len(existing)+1); err != nil {
		return domain.SMAdmin{}, err
	}

	admin.Groups = append(admin.Groups, group)

	return admin, nil
}

func (h srcdsUsecase) GetAdminGroups(ctx context.Context, admin domain.SMAdmin) ([]domain.SMGroups, error) {
	return h.srcdsRepository.GetAdminGroups(ctx, admin)
}

func (h srcdsUsecase) ServerAuth(ctx context.Context, req domain.ServerAuthReq) (string, error) {
	var server domain.Server

	errGetServer := h.serversUsecase.GetServerByPassword(ctx, req.Key, &server, true, false)
	if errGetServer != nil {
		return "", errGetServer
	}

	if server.Password != req.Key {
		return "", domain.ErrPermissionDenied
	}

	accessToken, errToken := newServerToken(server.ServerID, h.cookie)
	if errToken != nil {
		return "", errToken
	}

	server.TokenCreatedOn = time.Now()
	if errSaveServer := h.serversUsecase.SaveServer(ctx, &server); errSaveServer != nil {
		return "", errSaveServer
	}

	return accessToken, nil
}

func (h srcdsUsecase) Report(ctx context.Context, currentUser domain.UserProfile, req domain.CreateReportReq) (*domain.Report, error) {
	if req.Description == "" || len(req.Description) < 10 {
		return nil, fmt.Errorf("%w: description", domain.ErrParamInvalid)
	}

	// ServerStore initiated requests will have a sourceID set by the server
	// Web based reports the source should not be set, the reporter will be taken from the
	// current session information instead
	if !req.SourceID.Valid() {
		req.SourceID = currentUser.SteamID
	}

	if !req.SourceID.Valid() {
		return nil, domain.ErrSourceID
	}

	if !req.TargetID.Valid() {
		return nil, domain.ErrTargetID
	}

	if req.SourceID.Int64() == req.TargetID.Int64() {
		return nil, domain.ErrSelfReport
	}

	personSource, errCreatePerson := h.personUsecase.GetPersonBySteamID(ctx, req.SourceID)
	if errCreatePerson != nil {
		return nil, domain.ErrInternal
	}

	personTarget, errCreatePerson := h.personUsecase.GetOrCreatePersonBySteamID(ctx, req.TargetID)
	if errCreatePerson != nil {
		return nil, domain.ErrInternal
	}

	if personTarget.Expired() {
		if err := thirdparty.UpdatePlayerSummary(ctx, &personTarget); err != nil {
			slog.Error("Failed to update target player", log.ErrAttr(err))
		} else {
			if errSave := h.personUsecase.SavePerson(ctx, &personTarget); errSave != nil {
				slog.Error("Failed to save target player update", log.ErrAttr(err))
			}
		}
	}

	// Ensure the user doesn't already have an open report against the user
	existing, errReports := h.reportUsecase.GetReportBySteamID(ctx, personSource.SteamID, req.TargetID)
	if errReports != nil {
		if !errors.Is(errReports, domain.ErrNoResult) {
			return nil, errReports
		}
	}

	if existing.ReportID > 0 {
		return nil, domain.ErrReportExists
	}

	// TODO encapsulate all operations in single tx
	report := domain.NewReport()
	report.SourceID = req.SourceID
	report.ReportStatus = domain.Opened
	report.Description = req.Description
	report.TargetID = req.TargetID
	report.Reason = req.Reason
	report.ReasonText = req.ReasonText
	report.DemoTick = req.DemoTick
	report.PersonMessageID = req.PersonMessageID

	if errReportSave := h.reportUsecase.SaveReport(ctx, &report); errReportSave != nil {
		return nil, errReportSave
	}

	slog.Info("New report created successfully", slog.Int64("report_id", report.ReportID))

	conf := h.configUsecase.Config()

	demoURL := ""

	msg := discord.NewInGameReportResponse(report, conf.ExtURL(report), currentUser, conf.ExtURL(currentUser), demoURL)

	h.discordUsecase.SendPayload(domain.ChannelModLog, msg)

	return &report, nil
}

func (h srcdsUsecase) SetAdminGroups(ctx context.Context, authType domain.AuthType, identity string, groups ...domain.SMGroups) error {
	admin, errAdmin := h.srcdsRepository.GetAdminByIdentity(ctx, authType, identity)
	if errAdmin != nil {
		return errAdmin
	}

	// Delete existing groups.
	if errDelete := h.srcdsRepository.DeleteAdminGroups(ctx, admin); errDelete != nil && !errors.Is(errDelete, domain.ErrNoResult) {
		return errDelete
	}

	// If no groups are given to add, this is treated purely as a delete function
	if len(groups) == 0 {
		return nil
	}

	for i := range groups {
		if errInsert := h.srcdsRepository.InsertAdminGroup(ctx, admin, groups[i], i); errInsert != nil {
			return errInsert
		}
	}

	return nil
}

func (h srcdsUsecase) DelGroup(ctx context.Context, groupID int) error {
	group, errGroup := h.srcdsRepository.GetGroupByID(ctx, groupID)
	if errGroup != nil {
		return errGroup
	}

	return h.srcdsRepository.DeleteGroup(ctx, group)
}

const validFlags = "zabcdefghijklmnopqrst"

func (h srcdsUsecase) AddGroup(ctx context.Context, name string, flags string, immunityLevel int) (domain.SMGroups, error) {
	if name == "" {
		return domain.SMGroups{}, domain.ErrSMGroupName
	}

	if immunityLevel > 100 || immunityLevel < 0 {
		return domain.SMGroups{}, domain.ErrSMImmunity
	}

	for _, flag := range flags {
		if !strings.ContainsRune(validFlags, flag) {
			return domain.SMGroups{}, domain.ErrSMAdminFlagInvalid
		}
	}

	return h.srcdsRepository.AddGroup(ctx, domain.SMGroups{
		Flags:         flags,
		Name:          name,
		ImmunityLevel: immunityLevel,
	})
}

func validateAuthIdentity(ctx context.Context, authType domain.AuthType, identity string, password string) (string, error) {
	switch {
	case authType == domain.AuthTypeSteam:
		steamID, errSteamID := steamid.Resolve(ctx, identity)
		if errSteamID != nil {
			return "", domain.ErrInvalidSID
		}

		identity = steamID.String()
	case authType == domain.AuthTypeIP:
		if ip := net.ParseIP(identity); ip == nil || ip.To4() != nil {
			return "", domain.ErrInvalidIP
		}
	case authType == domain.AuthTypeName:
		if identity == "" {
			return "", domain.ErrSMInvalidAuthName
		}

		if password == "" {
			return "", domain.ErrSMRequirePassword
		}
	}

	return identity, nil
}

func (h srcdsUsecase) DelAdmin(ctx context.Context, adminID int) error {
	admin, errAdmin := h.srcdsRepository.GetAdminByID(ctx, adminID)
	if errAdmin != nil {
		return errAdmin
	}

	return h.srcdsRepository.DelAdmin(ctx, admin)
}

func (h srcdsUsecase) GetAdminByID(ctx context.Context, adminID int) (domain.SMAdmin, error) {
	return h.srcdsRepository.GetAdminByID(ctx, adminID)
}

func (h srcdsUsecase) SaveAdmin(ctx context.Context, admin domain.SMAdmin) (domain.SMAdmin, error) {
	realIdentity, errValidate := validateAuthIdentity(ctx, admin.AuthType, admin.Identity, admin.Password)
	if errValidate != nil {
		return domain.SMAdmin{}, errValidate
	}

	if admin.Immunity < 0 || admin.Immunity > 100 {
		return domain.SMAdmin{}, domain.ErrSMImmunity
	}

	var steamID steamid.SteamID
	if admin.AuthType == domain.AuthTypeSteam {
		steamID = steamid.New(realIdentity)
		if _, err := h.personUsecase.GetOrCreatePersonBySteamID(ctx, steamID); err != nil {
			return domain.SMAdmin{}, domain.ErrGetPerson
		}

		admin.Identity = string(steamID.Steam3())
		admin.SteamID = steamID
	}

	return h.srcdsRepository.SaveAdmin(ctx, admin)
}

func (h srcdsUsecase) AddAdmin(ctx context.Context, alias string, authType domain.AuthType, identity string, flags string, immunity int, password string) (domain.SMAdmin, error) {
	realIdentity, errValidate := validateAuthIdentity(ctx, authType, identity, password)
	if errValidate != nil {
		return domain.SMAdmin{}, errValidate
	}

	if immunity < 0 || immunity > 100 {
		return domain.SMAdmin{}, domain.ErrSMImmunity
	}

	admin, errAdmin := h.srcdsRepository.GetAdminByIdentity(ctx, authType, realIdentity)
	if errAdmin != nil && !errors.Is(errAdmin, domain.ErrNoResult) {
		return domain.SMAdmin{}, errAdmin
	}

	if errAdmin == nil {
		return admin, domain.ErrSMAdminExists
	}

	var steamID steamid.SteamID
	if authType == domain.AuthTypeSteam {
		steamID = steamid.New(realIdentity)
		if _, err := h.personUsecase.GetOrCreatePersonBySteamID(ctx, steamID); err != nil {
			return domain.SMAdmin{}, domain.ErrGetPerson
		}

		identity = string(steamID.Steam3())
	}

	return h.srcdsRepository.AddAdmin(ctx, domain.SMAdmin{
		SteamID:  steamID,
		AuthType: authType,
		Identity: identity,
		Password: password,
		Flags:    flags,
		Name:     alias,
		Immunity: immunity,
	})
}

func (h srcdsUsecase) Admins(ctx context.Context) ([]domain.SMAdmin, error) {
	return h.srcdsRepository.Admins(ctx)
}

func (h srcdsUsecase) Groups(ctx context.Context) ([]domain.SMGroups, error) {
	return h.srcdsRepository.Groups(ctx)
}

func (h srcdsUsecase) GetGroupByID(ctx context.Context, groupID int) (domain.SMGroups, error) {
	return h.srcdsRepository.GetGroupByID(ctx, groupID)
}

func (h srcdsUsecase) SaveGroup(ctx context.Context, group domain.SMGroups) (domain.SMGroups, error) {
	if group.Name == "" {
		return domain.SMGroups{}, domain.ErrSMGroupName
	}

	if group.ImmunityLevel > 100 || group.ImmunityLevel < 0 {
		return domain.SMGroups{}, domain.ErrSMImmunity
	}

	for _, flag := range group.Flags {
		if !strings.ContainsRune(validFlags, flag) {
			return domain.SMGroups{}, domain.ErrSMAdminFlagInvalid
		}
	}

	return h.srcdsRepository.SaveGroup(ctx, group)
}
