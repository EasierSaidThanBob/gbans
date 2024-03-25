package domain

import (
	"context"
	"errors"
	"io"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	"github.com/leighmacdonald/steamid/v4/steamid"
)

type DemoUsecase interface {
	Start(ctx context.Context)
	ExpiredDemos(ctx context.Context, limit uint64) ([]DemoInfo, error)
	GetDemoByID(ctx context.Context, demoID int64, demoFile *DemoFile) error
	GetDemoByName(ctx context.Context, demoName string, demoFile *DemoFile) error
	GetDemos(ctx context.Context, opts DemoFilter) ([]DemoFile, int64, error)
	Create(ctx context.Context, name string, content io.Reader, demoName string, serverID int) (*DemoFile, error)
	DropDemo(ctx context.Context, demoFile *DemoFile) error
}

type DemoRepository interface {
	ExpiredDemos(ctx context.Context, limit uint64) ([]DemoInfo, error)
	GetDemoByID(ctx context.Context, demoID int64, demoFile *DemoFile) error
	GetDemoByName(ctx context.Context, demoName string, demoFile *DemoFile) error
	GetDemos(ctx context.Context, opts DemoFilter) ([]DemoFile, int64, error)
	SaveDemo(ctx context.Context, demoFile *DemoFile) error
	DropDemo(ctx context.Context, demoFile *DemoFile) error
}

type DemoPlayerStats struct {
	Score      int `json:"score"`
	ScoreTotal int `json:"score_total"`
	Deaths     int `json:"deaths"`
}

type DemoMetaData struct {
	MapName string                     `json:"map_name"`
	Scores  map[string]DemoPlayerStats `json:"scores"`
}

type DemoFile struct {
	DemoID          int64                     `json:"demo_id"`
	ServerID        int                       `json:"server_id"`
	ServerNameShort string                    `json:"server_name_short"`
	ServerNameLong  string                    `json:"server_name_long"`
	Title           string                    `json:"title"`
	CreatedOn       time.Time                 `json:"created_on"`
	Downloads       int64                     `json:"downloads"`
	Size            int64                     `json:"size"`
	MapName         string                    `json:"map_name"`
	Archive         bool                      `json:"archive"` // When true, will not get auto deleted when flushing old demos
	Stats           map[steamid.SteamID]gin.H `json:"stats"`
	AssetID         uuid.UUID                 `json:"asset_id"`
}

type DemoInfo struct {
	DemoID  int64
	Title   string
	AssetID uuid.UUID
}

type Asset struct {
	AssetID  uuid.UUID `json:"asset_id"`
	Bucket   string    `json:"bucket"`
	Path     string    `json:"path"`
	Name     string    `json:"name"`
	MimeType string    `json:"mime_type"`
	Size     int64     `json:"size"`
	OldID    int64     `json:"old_id"` // Pre S3 id
}

func NewAsset(content []byte, bucket string, name string) (Asset, error) {
	detectedMime := mimetype.Detect(content)

	newID, errID := uuid.NewV4()
	if errID != nil {
		return Asset{}, errors.Join(errID, ErrUUIDCreate)
	}

	if name == "" {
		name = newID.String()
	}

	asset := Asset{
		AssetID:  newID,
		Bucket:   bucket,
		Path:     "/",
		Name:     name,
		MimeType: detectedMime.String(),
		Size:     int64(len(content)),
	}

	return asset, nil
}
