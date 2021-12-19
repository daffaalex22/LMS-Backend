package student

import (
	"backend/helper/password"
	"context"
	"errors"
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
		return Domain{}, errors.New("name empty")
	}
	if domain.Email == "" {
		return Domain{}, errors.New("email empty")
	}
	if domain.Password == "" {
		return Domain{}, errors.New("password empty")
	}
	hashedPass := password.HashPassword(domain.Password)
	domain.Password = hashedPass
	std, err := usecase.repo.Register(domain, ctx)
	if err != nil {
		return Domain{}, err
	}
	return std, nil
}

func (usecase *StudentUseCase) StudentUpdate(ctx context.Context, domain Domain, id uint) (Domain, error) {
	if domain.Name == "" {
		return Domain{}, errors.New("name empty")
	}
	if domain.Email == "" {
		return Domain{}, errors.New("email empty")
	}
	if domain.Avatar == "" {
		return Domain{}, errors.New("avatar empty")
	}
	if domain.Phone == 0 {
		return Domain{}, errors.New("phone empty")
	}
	if domain.Address == "" {
		return Domain{}, errors.New("address empty")
	}
	if domain.Password == "" {
		return Domain{}, errors.New("password empty")
	}
	if domain.ConfirmPassword == "" {
		return Domain{}, errors.New("confirm password empty")
	}
	if domain.ConfirmPassword != domain.Password {
		return Domain{}, errors.New("password must same with confirm password")
	}
	hashedPass := password.HashPassword(domain.Password)
	domain.Password = hashedPass
	std, err := usecase.repo.StudentUpdate(ctx, domain, id)
	if err != nil {
		return Domain{}, err
	}
	return std, nil
}

func (usecase *StudentUseCase) Login(domain Domain, ctx context.Context) (Domain, error) {
	if domain.Email == "" {
		return Domain{}, errors.New("email empty")
	}
	if domain.Password == "" {
		return Domain{}, errors.New("password empty")
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
	if std.Id == 0 {
		return Domain{}, err
	}
	return std, nil
}
