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

func NewUseCase(attRepo AttachmentsRepoInterface, contextTimeout time.Duration) AttachmentsUseCaseInterface {
	return &AttachmentsUseCase{
		repo: attRepo,
		ctx:  contextTimeout,
	}
}

func (usecase *AttachmentsUseCase) AttachmentsAdd(ctx context.Context, domain Domain) (Domain, error) {
	if domain.ContentType == "" {
		return Domain{}, err.ErrModuleIdEmpty
	}
	if domain.Title == "" {
		return Domain{}, err.ErrTitleEmpty
	}
	if domain.Url == "" {
		return Domain{}, err.ErrCaptionEmpty
	}

	_, err2 := usecase.repo.CheckContent(ctx, domain.ContentType, domain.ContentId)
	if err2 != nil {
		fmt.Println("Error CheckModule")
		return Domain{}, err.ErrIdModule
	}

	attachments, result := usecase.repo.AttachmentsAdd(ctx, domain)
	if result != nil {
		fmt.Println("Error Repo AttachmentsAdd")
		return Domain{}, result
	}
	return attachments, nil
}

func (usecase *AttachmentsUseCase) AttachmentsUpdate(ctx context.Context, domain Domain, id uint) (Domain, error) {
	domain.Id = (id)
	if domain.ContentType == "" {
		return Domain{}, err.ErrModuleIdEmpty
	}
	if domain.Title == "" {
		return Domain{}, err.ErrTitleEmpty
	}
	if domain.Url == "" {
		return Domain{}, err.ErrCaptionEmpty
	}

	_, err1 := usecase.repo.CheckContent(ctx, domain.ContentType, domain.ContentId)
	if err1 != nil {
		return Domain{}, err.ErrIdModule
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
