package srcds

import (
	"context"
	"errors"
	"log/slog"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/leighmacdonald/gbans/internal/database"
	"github.com/leighmacdonald/gbans/internal/domain"
	"github.com/leighmacdonald/steamid/v4/steamid"
)

type srcdsRepository struct {
	database database.Database
}

func NewRepository(database database.Database) domain.SRCDSRepository {
	return srcdsRepository{database: database}
}

func (r srcdsRepository) GroupOverrides(ctx context.Context) ([]domain.SMGroupOverrides, error) {
	rows, errRows := r.database.QueryBuilder(ctx, r.database.Builder().
		Select("group_id", "type", "name", "access", "created_on", "updated_on").
		From("sm_group_overrides"))
	if errRows != nil {
		return nil, r.database.DBErr(errRows)
	}

	var overrides []domain.SMGroupOverrides

	for rows.Next() {
		var override domain.SMGroupOverrides
		if errScan := rows.Scan(&override.GroupID, &override.Type, &override.Name, &override.Access,
			&override.CreatedOn, &override.UpdatedOn); errScan != nil {
			return nil, r.database.DBErr(errScan)
		}

		overrides = append(overrides, override)
	}

	return overrides, nil
}

func (r srcdsRepository) Overrides(ctx context.Context) ([]domain.SMOverrides, error) {
	rows, errRows := r.database.QueryBuilder(ctx, r.database.Builder().
		Select("type", "name", "flags", "created_on", "updated_on").
		From("sm_overrides"))
	if errRows != nil {
		return nil, r.database.DBErr(errRows)
	}

	var overrides []domain.SMOverrides

	for rows.Next() {
		var override domain.SMOverrides
		if errScan := rows.Scan(&override.Type, &override.Name, &override.Flags,
			&override.CreatedOn, &override.UpdatedOn); errScan != nil {
			return nil, r.database.DBErr(errScan)
		}

		overrides = append(overrides, override)
	}

	return overrides, nil
}

func (r srcdsRepository) AddGroupOverride(ctx context.Context, override domain.SMGroupOverrides) error {
	return r.database.DBErr(r.database.ExecInsertBuilder(ctx, r.database.Builder().
		Insert("sm_group_overrides").
		SetMap(map[string]interface{}{
			"group_id":   override.GroupID,
			"type":       override.Type,
			"name":       override.Name,
			"access":     override.Access,
			"created_in": override.CreatedOn,
			"updated_on": override.UpdatedOn,
		})))
}

func (r srcdsRepository) AddOverride(ctx context.Context, overrides domain.SMOverrides) error {
	return r.database.DBErr(r.database.ExecInsertBuilder(ctx, r.database.Builder().
		Insert("sm_overrides").
		SetMap(map[string]interface{}{
			"type":       overrides.Type,
			"name":       overrides.Name,
			"flags":      overrides.Flags,
			"created_on": overrides.CreatedOn,
			"updated_on": overrides.UpdatedOn,
		})))
}

func (r srcdsRepository) GetAdminGroups(ctx context.Context, admin domain.SMAdmin) ([]domain.SMGroups, error) {
	rows, errRows := r.database.QueryBuilder(ctx, r.database.Builder().
		Select("g.id", "g.flags", "g.name", "g.immunity_level", "g.created_on", "g.updated_on").
		From("sm_groups g").
		LeftJoin("sm_admins_groups ag ON ag.group_id = g.id").
		Where(sq.Eq{"ag.admin_id": admin.AdminID}))
	if errRows != nil {
		return nil, errRows
	}

	var groups []domain.SMGroups

	for rows.Next() {
		var group domain.SMGroups
		if errScan := rows.Scan(&group.GroupID, &group.Flags, &group.Name, &group.ImmunityLevel,
			&group.CreatedOn, &group.UpdatedOn); errScan != nil {
			return nil, r.database.DBErr(errScan)
		}

		groups = append(groups, group)
	}

	if groups == nil {
		groups = []domain.SMGroups{}
	}

	return groups, nil
}

func (r srcdsRepository) Admins(ctx context.Context) ([]domain.SMAdmin, error) {
	groups, errGroups := r.Groups(ctx)
	if errGroups != nil && !errors.Is(errGroups, domain.ErrNoResult) {
		return nil, errGroups
	}

	rows, errRows := r.database.QueryBuilder(ctx, r.database.Builder().
		Select("a.id", "a.steam_id", "a.authtype", "a.identity", "a.password", "a.flags", "a.name", "a.immunity",
			"a.created_on", "a.updated_on", "array_agg(sag.group_id) as group_ids").
		From("sm_admins a").
		LeftJoin("sm_admins_groups sag on a.id = sag.admin_id").
		GroupBy("a.id"))
	if errRows != nil {
		if errors.Is(errRows, domain.ErrNoResult) {
			return []domain.SMAdmin{}, nil
		}

		return nil, r.database.DBErr(errRows)
	}

	var admins []domain.SMAdmin

	for rows.Next() {
		var (
			// array_agg will return a {null} if no group entry exists
			groupIds []*int
			admin    = domain.SMAdmin{Groups: []domain.SMGroups{}}
		)

		if errScan := rows.Scan(&admin.AdminID, &admin.SteamID, &admin.AuthType, &admin.Identity, &admin.Password,
			&admin.Flags, &admin.Name, &admin.Immunity, &admin.CreatedOn, &admin.UpdatedOn, &groupIds); errScan != nil {
			return nil, r.database.DBErr(errScan)
		}

		for _, groupID := range groupIds {
			if groupID == nil {
				continue
			}

			for _, group := range groups {
				if group.GroupID == *groupID {
					admin.Groups = append(admin.Groups, group)
				}
			}
		}

		admins = append(admins, admin)
	}

	return admins, nil
}

func (r srcdsRepository) Groups(ctx context.Context) ([]domain.SMGroups, error) {
	rows, errRows := r.database.QueryBuilder(ctx, r.database.Builder().
		Select("id", "flags", "name", "immunity_level", "created_on", "updated_on").
		From("sm_groups"))
	if errRows != nil {
		if errors.Is(errRows, domain.ErrNoResult) {
			return []domain.SMGroups{}, nil
		}

		return nil, r.database.DBErr(errRows)
	}

	var groups []domain.SMGroups

	for rows.Next() {
		var group domain.SMGroups
		if errScan := rows.Scan(&group.GroupID, &group.Flags, &group.Name, &group.ImmunityLevel,
			&group.CreatedOn, &group.UpdatedOn); errScan != nil {
			return nil, r.database.DBErr(errScan)
		}

		groups = append(groups, group)
	}

	if groups == nil {
		groups = []domain.SMGroups{}
	}

	return groups, nil
}

func (r srcdsRepository) GetAdminByID(ctx context.Context, adminID int) (domain.SMAdmin, error) {
	var (
		admin domain.SMAdmin
		id64  int64
	)

	row, errRow := r.database.QueryRowBuilder(ctx, r.database.Builder().
		Select("id", "steam_id", "authtype", "identity", "password", "flags", "name", "immunity", "created_on", "updated_on").
		From("sm_admins").
		Where(sq.And{sq.Eq{"id": adminID}}))
	if errRow != nil {
		return admin, r.database.DBErr(errRow)
	}

	if errScan := row.Scan(&admin.AdminID, &id64, &admin.AuthType, &admin.Identity,
		&admin.Password, &admin.Flags, &admin.Name, &admin.Immunity, &admin.CreatedOn, &admin.UpdatedOn); errScan != nil {
		return admin, r.database.DBErr(errScan)
	}

	admin.SteamID = steamid.New(id64)

	groups, errGroup := r.GetAdminGroups(ctx, admin)
	if errGroup != nil && !errors.Is(errGroup, domain.ErrNoResult) {
		return domain.SMAdmin{}, errGroup
	}

	admin.Groups = groups

	return admin, nil
}

func (r srcdsRepository) SaveAdmin(ctx context.Context, admin domain.SMAdmin) (domain.SMAdmin, error) {
	admin.UpdatedOn = time.Now()

	var sid64 *int64

	if admin.SteamID.Valid() {
		sidValue := admin.SteamID.Int64()
		sid64 = &sidValue
	}

	if err := r.database.ExecUpdateBuilder(ctx, r.database.Builder().
		Update("sm_admins").
		SetMap(map[string]interface{}{
			"steam_id":   sid64,
			"authtype":   admin.AuthType,
			"identity":   admin.Identity,
			"password":   admin.Password,
			"flags":      admin.Flags,
			"name":       admin.Name,
			"immunity":   admin.Immunity,
			"updated_on": admin.UpdatedOn,
		}).Where(sq.Eq{"id": admin.AdminID})); err != nil {
		return domain.SMAdmin{}, err
	}

	return admin, nil
}

func (r srcdsRepository) GetAdminByIdentity(ctx context.Context, authType domain.AuthType, identity string) (domain.SMAdmin, error) {
	var (
		admin domain.SMAdmin
		id64  int64
	)

	row, errRow := r.database.QueryRowBuilder(ctx, r.database.Builder().
		Select("id", "steam_id", "authtype", "identity", "password", "flags", "name", "immunity", "created_on", "updated_on").
		From("sm_admins").
		Where(sq.And{sq.Eq{"authtype": authType}, sq.Eq{"identity": identity}}))
	if errRow != nil {
		return admin, r.database.DBErr(errRow)
	}

	if errScan := row.Scan(&admin.AdminID, &id64, &admin.AuthType, &admin.Identity,
		&admin.Password, &admin.Flags, &admin.Name, &admin.Immunity, &admin.CreatedOn, &admin.UpdatedOn); errScan != nil {
		return admin, r.database.DBErr(errScan)
	}

	admin.SteamID = steamid.New(id64)

	return admin, nil
}

func (r srcdsRepository) GetGroupByName(ctx context.Context, groupName string) (domain.SMGroups, error) {
	var group domain.SMGroups

	row, errRow := r.database.QueryRowBuilder(ctx, r.database.Builder().
		Select("id", "flags", "name", "immunity_level", "created_on", "updated_on").
		From("sm_admins").
		Where(sq.Eq{"name": groupName}))
	if errRow != nil {
		return group, r.database.DBErr(errRow)
	}

	if errScan := row.Scan(&group.GroupID, &group.Flags, &group.Name, &group.ImmunityLevel,
		&group.CreatedOn, &group.UpdatedOn); errScan != nil {
		return group, r.database.DBErr(errScan)
	}

	return group, nil
}

func (r srcdsRepository) GetGroupByID(ctx context.Context, groupID int) (domain.SMGroups, error) {
	var group domain.SMGroups

	row, errRow := r.database.QueryRowBuilder(ctx, r.database.Builder().
		Select("id", "flags", "name", "immunity_level", "created_on", "updated_on").
		From("sm_groups").
		Where(sq.Eq{"id": groupID}))
	if errRow != nil {
		return group, r.database.DBErr(errRow)
	}

	if errScan := row.Scan(&group.GroupID, &group.Flags, &group.Name, &group.ImmunityLevel,
		&group.CreatedOn, &group.UpdatedOn); errScan != nil {
		return group, r.database.DBErr(errScan)
	}

	return group, nil
}

func (r srcdsRepository) InsertAdminGroup(ctx context.Context, admin domain.SMAdmin, group domain.SMGroups, inheritOrder int) error {
	now := time.Now()

	return r.database.DBErr(r.database.ExecInsertBuilder(ctx, r.database.Builder().
		Insert("sm_admins_groups").
		SetMap(map[string]interface{}{
			"admin_id":      admin.AdminID,
			"group_id":      group.GroupID,
			"inherit_order": inheritOrder,
			"created_on":    now,
			"updated_on":    now,
		})))
}

func (r srcdsRepository) DeleteAdminGroups(ctx context.Context, admin domain.SMAdmin) error {
	if err := r.database.ExecDeleteBuilder(ctx, r.database.Builder().
		Delete("sm_admins_groups").
		Where(sq.Eq{"admin_id": admin.AdminID})); err != nil {
		return r.database.DBErr(err)
	}

	slog.Info("Deleted SM admin groups", slog.String("steam_id", admin.SteamID.String()))

	return nil
}

func (r srcdsRepository) DeleteAdminGroup(ctx context.Context, admin domain.SMAdmin, group domain.SMGroups) error {
	return r.database.ExecDeleteBuilder(ctx, r.database.Builder().
		Delete("sm_admins_groups").
		Where(sq.And{sq.Eq{"admin_id": admin.AdminID}, sq.Eq{"group_id": group.GroupID}}))
}

func (r srcdsRepository) SaveGroup(ctx context.Context, group domain.SMGroups) (domain.SMGroups, error) {
	group.UpdatedOn = time.Now()
	if err := r.database.ExecUpdateBuilder(ctx, r.database.Builder().
		Update("sm_groups").
		SetMap(map[string]interface{}{
			"name":           group.Name,
			"immunity_level": group.ImmunityLevel,
			"flags":          group.Flags,
		}).
		Where(sq.Eq{"id": group.GroupID})); err != nil {
		return group, r.database.DBErr(err)
	}

	return group, nil
}

func (r srcdsRepository) DeleteGroup(ctx context.Context, group domain.SMGroups) error {
	if err := r.database.ExecDeleteBuilder(ctx, r.database.Builder().
		Delete("sm_admins_groups").
		Where(sq.Eq{"group_id": group.GroupID})); err != nil {
		return r.database.DBErr(err)
	}

	if err := r.database.ExecDeleteBuilder(ctx, r.database.Builder().
		Delete("sm_group_overrides").
		Where(sq.Eq{"group_id": group.GroupID})); err != nil {
		return r.database.DBErr(err)
	}

	if err := r.database.ExecDeleteBuilder(ctx, r.database.Builder().
		Delete("sm_group_immunity").
		Where(sq.Or{sq.Eq{"group_id": group.GroupID}, sq.Eq{"other_id": group.GroupID}})); err != nil {
		return r.database.DBErr(err)
	}

	if err := r.database.ExecDeleteBuilder(ctx, r.database.Builder().
		Delete("sm_groups").
		Where(sq.Eq{"id": group.GroupID})); err != nil {
		return r.database.DBErr(err)
	}

	slog.Info("Deleted SM Group", slog.Int("group_id", group.GroupID), slog.String("name", group.Name))

	return nil
}

func (r srcdsRepository) AddGroup(ctx context.Context, group domain.SMGroups) (domain.SMGroups, error) {
	now := time.Now()

	if err := r.database.ExecInsertBuilderWithReturnValue(ctx, r.database.
		Builder().
		Insert("sm_groups").
		SetMap(map[string]interface{}{
			"name":           group.Name,
			"immunity_level": group.ImmunityLevel,
			"flags":          group.Flags,
			"created_on":     now,
			"updated_on":     now,
		}).
		Suffix("RETURNING id"), &group.GroupID); err != nil {
		return group, err
	}

	slog.Info("Created SM Group", slog.Int("group_id", group.GroupID), slog.String("name", group.Name))

	return group, nil
}

func (r srcdsRepository) DelAdmin(ctx context.Context, admin domain.SMAdmin) error {
	if err := r.database.ExecDeleteBuilder(ctx, r.database.Builder().
		Delete("sm_admins_groups").
		Where(sq.Eq{"admin_id": admin.AdminID})); err != nil {
		return r.database.DBErr(err)
	}

	if err := r.database.ExecDeleteBuilder(ctx, r.database.Builder().
		Delete("sm_admins").
		Where(sq.Eq{"id": admin.AdminID})); err != nil {
		return r.database.DBErr(err)
	}

	slog.Info("Deleted SM Admin", slog.Int("id", admin.AdminID),
		slog.String("steam_id", admin.SteamID.String()))

	return nil
}

func (r srcdsRepository) AddAdmin(ctx context.Context, admin domain.SMAdmin) (domain.SMAdmin, error) {
	var nullableSID64 *int64

	if admin.SteamID.Valid() {
		steamID := admin.SteamID.Int64()
		nullableSID64 = &steamID
	}

	now := time.Now()

	admin.CreatedOn = now
	admin.UpdatedOn = now

	if err := r.database.ExecInsertBuilderWithReturnValue(ctx, r.database.Builder().
		Insert("sm_admins").
		SetMap(map[string]interface{}{
			"steam_id":   nullableSID64,
			"authtype":   admin.AuthType,
			"identity":   admin.Identity,
			"password":   admin.Password,
			"flags":      admin.Flags,
			"name":       admin.Name,
			"immunity":   admin.Immunity,
			"created_on": admin.CreatedOn,
			"updated_on": admin.UpdatedOn,
		}).
		Suffix("RETURNING id"), &admin.AdminID); err != nil {
		return admin, r.database.DBErr(err)
	}

	slog.Info("Added SM Admin", slog.Int("id", admin.AdminID), slog.String("identity", admin.Identity))

	return admin, nil
}
