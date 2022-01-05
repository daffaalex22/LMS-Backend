package videos

import (
	"backend/helper/err"
	"context"
	"fmt"
	"time"
)

type VideosUseCase struct {
	//repo
	repo VideosRepoInterface
	ctx  time.Duration
}

func NewUseCase(mdsRepo VideosRepoInterface, contextTimeout time.Duration) VideosUseCaseInterface {
	return &VideosUseCase{
		repo: mdsRepo,
		ctx:  contextTimeout,
	}
}

func (usecase *VideosUseCase) VideosAdd(ctx context.Context, domain Domain) (Domain, error) {
	if domain.ModuleId == 0 {
		return Domain{}, err.ErrModuleIdEmpty
	}
	if domain.Title == "" {
		return Domain{}, err.ErrTitleEmpty
	}
	if domain.Caption == "" {
		return Domain{}, err.ErrCaptionEmpty
	}
	if domain.Url == "" {
		return Domain{}, err.ErrUrlEmpty
	}
	if domain.Order == 0 {
		return Domain{}, err.ErrOrderEmpty
	}

	dataModule, err2 := usecase.repo.CheckModule(ctx, domain.ModuleId)
	if err2 != nil {
		fmt.Println("Error CheckModule")
		return Domain{}, err.ErrIdModule
	}
	domain.Module = dataModule

	Videos, result := usecase.repo.VideosAdd(ctx, domain)
	if result != nil {
		fmt.Println("Error Repo VideosAdd")
		return Domain{}, result
	}
	return Videos, nil
}

func (usecase *VideosUseCase) VideosUpdate(ctx context.Context, domain Domain, id uint) (Domain, error) {
	domain.Id = (id)
	if domain.ModuleId == 0 {
		return Domain{}, err.ErrModuleIdEmpty
	}
	if domain.Title == "" {
		return Domain{}, err.ErrTitleEmpty
	}
	if domain.Caption == "" {
		return Domain{}, err.ErrCaptionEmpty
	}
	if domain.Url == "" {
		return Domain{}, err.ErrUrlEmpty
	}
	if domain.Order == 0 {
		return Domain{}, err.ErrOrderEmpty
	}

	dataModule, err1 := usecase.repo.CheckModule(ctx, domain.ModuleId)
	if err1 != nil {
		return Domain{}, err.ErrIdModule
	}
	domain.Module = dataModule

	Videos, result := usecase.repo.VideosUpdate(ctx, domain, id)
	if result != nil {
		return Domain{}, result
	}
	return Videos, nil
}

func (usecase *VideosUseCase) VideosGetByModuleId(ctx context.Context, moduleId uint) ([]Domain, error) {
	Videos, err := usecase.repo.VideosGetByModuleId(ctx, moduleId)
	if err != nil {
		return []Domain{}, err
	}
	return Videos, nil
}

func (usecase *VideosUseCase) VideosDelete(ctx context.Context, id uint) error {
	result := usecase.repo.VideosDelete(ctx, id)
	if result != nil {
		return result
	}
	return nil
}
