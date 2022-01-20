package attachments

import (
	"context"
	"time"
)

type Domain struct {
	Id          uint
	ContentType string
	ContentId   uint
	VideoId     uint
	ReadingId   uint
	// QuizId     uint
	Title    string
	Url      string
	CreateAt time.Time
	UpdateAt time.Time
}

type AttachmentsUseCaseInterface interface {
	AttachmentsAdd(ctx context.Context, domain Domain) (Domain, error)
	AttachmentsUpdate(ctx context.Context, domain Domain, id uint) (Domain, error)
	AttachmentsDelete(ctx context.Context, id uint) error
	AttachmentsGetById(ctx context.Context, moduleId uint) (Domain, error)
}

type AttachmentsRepoInterface interface {
	AttachmentsAdd(ctx context.Context, domain Domain) (Domain, error)
	AttachmentsUpdate(ctx context.Context, domain Domain, id uint) (Domain, error)
	AttachmentsDelete(ctx context.Context, id uint) error
	AttachmentsGetById(ctx context.Context, moduleId uint) (Domain, error)

	CheckContent(ctx context.Context, contentType string, id uint) (interface{}, error)
	CheckVideos(ctx context.Context, videoId uint) error
	CheckReadings(ctx context.Context, readingId uint) error
	// CheckQuizzes(ctx context.Context, quizId uint) error
}
