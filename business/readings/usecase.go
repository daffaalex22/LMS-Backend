package readings

import (
	"backend/helper/err"
	"context"
	"fmt"
	"time"
)

type ReadingsUseCase struct {
	//repo
	repo ReadingsRepoInterface
	ctx  time.Duration
}

func NewUseCase(mdsRepo ReadingsRepoInterface, contextTimeout time.Duration) ReadingsUseCaseInterface {
	return &ReadingsUseCase{
		repo: mdsRepo,
		ctx:  contextTimeout,
	}
}

func (usecase *ReadingsUseCase) ReadingsGetById(ctx context.Context, id uint) (Domain, error) {
	result, err := usecase.repo.ReadingsGetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return result, nil
}

func (usecase *ReadingsUseCase) ReadingsAdd(ctx context.Context, domain Domain) (Domain, error) {
	if domain.ModuleId == 0 {
		return Domain{}, err.ErrModuleIdEmpty
	}
	if domain.Title == "" {
		return Domain{}, err.ErrTitleEmpty
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

	readings, result := usecase.repo.ReadingsAdd(ctx, domain)
	if result != nil {
		fmt.Println("Error Repo ReadingsAdd")
		return Domain{}, result
	}
	return readings, nil
}

func (usecase *ReadingsUseCase) ReadingsUpdate(ctx context.Context, domain Domain, id uint) (Domain, error) {
	domain.Id = (id)
	if domain.ModuleId == 0 {
		fmt.Println("Error CheckModuleID")
		return Domain{}, err.ErrModuleIdEmpty
	}
	if domain.Title == "" {
		return Domain{}, err.ErrTitleEmpty
	}
	if domain.Order == 0 {
		return Domain{}, err.ErrOrderEmpty
	}
	dataModule, err1 := usecase.repo.CheckModule(ctx, domain.ModuleId)
	if err1 != nil {
		return Domain{}, err.ErrIdModule
	}
	domain.Module = dataModule

	readings, result := usecase.repo.ReadingsUpdate(ctx, domain, id)
	if result != nil {
		return Domain{}, result
	}
	return readings, nil
}

func (usecase *ReadingsUseCase) ReadingsGetByModuleId(ctx context.Context, moduleId uint) ([]Domain, error) {
	readings, err := usecase.repo.ReadingsGetByModuleId(ctx, moduleId)
	if err != nil {
		return []Domain{}, err
	}
	return readings, nil
}

func (usecase *ReadingsUseCase) ReadingsDelete(ctx context.Context, id uint) error {
	result := usecase.repo.ReadingsDelete(ctx, id)
	if result != nil {
		return result
	}
	return nil
}
