package student

import (
	"backend/helper/err"
	"backend/helper/password"
	"context"
	"time"
)

type StudentUseCase struct {
	//repo
	repo StudentRepoInterface
	ctx  time.Duration
}

func NewUseCase(stdRepo StudentRepoInterface, contextTimeout time.Duration) StudentUseCaseInterface {
	return &StudentUseCase{
		repo: stdRepo,
		ctx:  contextTimeout,
	}
}

func (usecase *StudentUseCase) Register(domain *Domain, ctx context.Context) (Domain, error) {
	if domain.Name == "" {
		return Domain{}, err.ErrNameEmpty
	}
	if domain.Email == "" {
		return Domain{}, err.ErrEmailEmpty
	}
	if domain.Password == "" {
		return Domain{}, err.ErrPasswordEmpty
	}
	hashedPass := password.HashPassword(domain.Password)
	domain.Password = hashedPass
	std, err1 := usecase.repo.Register(domain, ctx)
	if err1 != nil {
		return Domain{}, err1
	}
	return std, nil
}

func (usecase *StudentUseCase) StudentUpdate(ctx context.Context, domain Domain, id uint) (Domain, error) {
	//checking pass input
	if domain.Password != "" {
		data, err1 := usecase.repo.GetProfile(ctx, id)
		if err1 != nil {
			return Domain{}, err1
		}
		if !password.CheckSamePassword(domain.Password, data.Password) {
			return Domain{}, err.ErrWrongPassword
		} else {
			domain.Password = data.Password
		}
	} else {
		return Domain{}, err.ErrPasswordEmpty
	}

	if domain.Name == "" {
		return Domain{}, err.ErrNameEmpty
	}
	if domain.Email == "" {
		return Domain{}, err.ErrEmailEmpty
	}
	if domain.Avatar == "" {
		return Domain{}, err.ErrAvatarEmpty
	}
	if domain.Phone == "" {
		return Domain{}, err.ErrPhoneEmpty
	}
	if domain.Address == "" {
		return Domain{}, err.ErrAddressEmpty
	}
	if domain.NewPassword != "" {
		hashedPass := password.HashPassword(domain.NewPassword)
		domain.Password = hashedPass
	}

	hashedPass := password.HashPassword(domain.Password)
	domain.Password = hashedPass
	std, err1 := usecase.repo.StudentUpdate(ctx, domain, id)
	if err1 != nil {
		return Domain{}, err1
	}
	return std, nil
}

func (usecase *StudentUseCase) Login(domain Domain, ctx context.Context) (Domain, error) {
	if domain.Email == "" {
		return Domain{}, err.ErrEmailEmpty
	}
	if domain.Password == "" {
		return Domain{}, err.ErrPasswordEmpty
	}
	std, err := usecase.repo.Login(domain, ctx)
	if err != nil {
		return Domain{}, err
	}
	return std, nil

}

func (usecase *StudentUseCase) GetProfile(ctx context.Context, id uint) (Domain, error) {
	std, err := usecase.repo.GetProfile(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return std, nil
}
