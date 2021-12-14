package student

import (
	_middleware "backend/app/middleware"
	"backend/helper/password"
	"context"
	"errors"
	"fmt"
	"time"
)

type StudentUseCase struct {
	//repo
	repo StudentRepoInterface
	ctx  time.Duration
	jwt  *_middleware.ConfigJWT
}

func NewUseCase(stdRepo StudentRepoInterface, contextTimeout time.Duration, configJWT *_middleware.ConfigJWT) StudentUseCaseInterface {
	return &StudentUseCase{
		repo: stdRepo,
		ctx:  contextTimeout,
		jwt:  configJWT,
	}
}

func (usecase *StudentUseCase) Register(domain *Domain, ctx context.Context) (Domain, error) {
	if domain.Name == "" {
		return Domain{}, errors.New("namd empty")
	}
	if domain.Email == "" {
		return Domain{}, errors.New("email empty")
	}
	if domain.Password == "" {
		return Domain{}, errors.New("password empty")
	}
	if domain.Avatar == "" {
		return Domain{}, errors.New("avatar empty")
	}
	if domain.Phone == 0 {
		return Domain{}, errors.New("email empty")
	}
	if domain.Address == "" {
		return Domain{}, errors.New("address empty")

	}
	hashedPass := password.HashPassword(domain.Password)
	domain.Password = hashedPass
	std, err := usecase.repo.Register(domain, ctx)
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
	pass := domain.Password
	std, err := usecase.repo.Login(domain, ctx)
	fmt.Println(std)
	if err != nil {
		return Domain{}, err
	}
	if password.CheckSamePassword(pass, domain.Password) {
		fmt.Println(std)
		return Domain{}, errors.New("wrong password")
	}
	std.Token = usecase.jwt.GenerateToken(std.Id)
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
