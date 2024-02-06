package discord

import (
	"context"
	"errors"
	"github.com/bwmarrin/discordgo"
	"github.com/leighmacdonald/gbans/internal/domain"
	"regexp"
	"time"
)

type discordUsecase struct {
	dr  domain.DiscordRepository
	wfu domain.WordFilterUsecase
}

func NewDiscordUsecase(repository domain.DiscordRepository, wfu domain.WordFilterUsecase) domain.DiscordUsecase {
	return &discordUsecase{dr: repository, wfu: wfu}
}

func (d discordUsecase) Shutdown(guildID string) {
	d.dr.Shutdown(guildID)
}

func (d discordUsecase) RegisterHandler(cmd domain.Cmd, handler domain.SlashCommandHandler) error {
	return d.dr.RegisterHandler(cmd, handler)
}

func (d discordUsecase) Start() error {
	return d.dr.Start()
}

func (d discordUsecase) SendPayload(channelID domain.DiscordChannel, embed *discordgo.MessageEmbed) {
	d.dr.SendPayload(channelID, embed)
}

func (d discordUsecase) FilterAdd(ctx context.Context, user domain.PersonInfo, pattern string, isRegex bool) (*discordgo.MessageEmbed, error) {
	if isRegex {
		_, rxErr := regexp.Compile(pattern)
		if rxErr != nil {
			return nil, errors.Join(rxErr, domain.ErrInvalidFilterRegex)
		}
	}

	filter := domain.Filter{
		AuthorID:  user.GetSteamID(),
		Pattern:   pattern,
		IsRegex:   isRegex,
		IsEnabled: true,
		CreatedOn: time.Now(),
		UpdatedOn: time.Now(),
	}
	if errFilterAdd := d.wfu.SaveFilter(ctx, user, &filter); errFilterAdd != nil {
		return nil, domain.ErrCommandFailed
	}

	FilterAddMessage(filter)
}
