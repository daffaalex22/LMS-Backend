package attachments

import (
	"backend/helper/err"
	"context"
	"fmt"
	"time"
)

type AttachmentsUseCase struct {
	//repo
	repo AttachmentsRepoInterface
	ctx  time.Duration
}

func NewAttachmentsUseCase(attRepo AttachmentsRepoInterface, contextTimeout time.Duration) AttachmentsUseCaseInterface {
	return &AttachmentsUseCase{
		repo: attRepo,
		ctx:  contextTimeout,
	}
}

func (usecase *AttachmentsUseCase) AttachmentsAdd(ctx context.Context, domain Domain) (Domain, error) {
	if domain.ContentId == 0 {
		return Domain{}, err.ErrContentIdEmpty
	}
	if domain.ContentType == "" {
		return Domain{}, err.ErrContentTypeEmpty
	}
	if domain.Title == "" {
		return Domain{}, err.ErrTitleEmpty
	}
	if domain.Url == "" {
		return Domain{}, err.ErrUrlEmpty
	}

	var err1 error
	if domain.ContentType == "Videos" {
		domain.VideoId = domain.ContentId
		err1 = usecase.repo.CheckVideos(ctx, domain.ContentId)
	} else if domain.ContentType == "Readings" {
		domain.ReadingId = domain.ContentId
		err1 = usecase.repo.CheckReadings(ctx, domain.ContentId)
	} else {
		return Domain{}, err.ErrContentType
	}
	// else if domain.ContentType == "Quizzes" {
	// 	domain.VideoId = domain.ContentId
	// }

	if err1 != nil {
		fmt.Println("Error CheckContent")
		return Domain{}, err.ErrContentNotFound
	}

	attachments, result := usecase.repo.AttachmentsAdd(ctx, domain)
	if result != nil {
		return Domain{}, result
	}
	return attachments, nil
}

func (usecase *AttachmentsUseCase) AttachmentsUpdate(ctx context.Context, domain Domain, id uint) (Domain, error) {
	domain.Id = (id)
	if domain.ContentId == 0 {
		return Domain{}, err.ErrContentIdEmpty
	}
	if domain.ContentType == "" {
		return Domain{}, err.ErrContentTypeEmpty
	}
	if domain.Title == "" {
		return Domain{}, err.ErrTitleEmpty
	}
	if domain.Url == "" {
		return Domain{}, err.ErrUrlEmpty
	}

	_, err1 := usecase.repo.CheckContent(ctx, domain.ContentType, domain.ContentId)
	if err1 != nil {
		return Domain{}, err.ErrContentNotFound
	}

	attachments, result := usecase.repo.AttachmentsUpdate(ctx, domain, id)
	if result != nil {
		return Domain{}, result
	}
	return attachments, nil
}

func (usecase *AttachmentsUseCase) AttachmentsGetById(ctx context.Context, moduleId uint) (Domain, error) {
	attachments, err := usecase.repo.AttachmentsGetById(ctx, moduleId)
	if err != nil {
		return Domain{}, err
	}
	return attachments, nil
}

func (usecase *AttachmentsUseCase) AttachmentsDelete(ctx context.Context, id uint) error {
	result := usecase.repo.AttachmentsDelete(ctx, id)
	if result != nil {
		return result
	}
	return nil
}
