package domain

import "context"

type AppealRepository interface {
	SaveBanMessage(ctx context.Context, message *BanAppealMessage) error
	GetBanMessages(ctx context.Context, banID int64) ([]BanAppealMessage, error)
	GetBanMessageByID(ctx context.Context, banMessageID int64, message *BanAppealMessage) error
	DropBanMessage(ctx context.Context, message *BanAppealMessage) error
	GetAppealsByActivity(ctx context.Context, opts AppealQueryFilter) ([]AppealOverview, int64, error)
}

type AppealUsecase interface {
	SaveBanMessage(ctx context.Context, curUser UserProfile, reportMessageID int64, newMsg string) (*BanAppealMessage, error)
	GetBanMessages(ctx context.Context, userProfile UserProfile, banID int64) ([]BanAppealMessage, error)
	GetBanMessageByID(ctx context.Context, banMessageID int64, message *BanAppealMessage) error
	DropBanMessage(ctx context.Context, curUser UserProfile, banMessageID int64) error
	GetAppealsByActivity(ctx context.Context, opts AppealQueryFilter) ([]AppealOverview, int64, error)
}

type NewBanMessage struct {
	Message string `json:"message"`
}
