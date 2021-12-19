package teacher

import (
	"backend/helper/password"
	"context"
	"errors"
	"time"
)

type TeacherUseCase struct {
	//repo
	repo TeacherRepoInterface
	ctx  time.Duration
}

func NewUseCase(tchRepo TeacherRepoInterface, contextTimeout time.Duration) TeacherUseCaseInterface {
	return &TeacherUseCase{
		repo: tchRepo,
		ctx:  contextTimeout,
	}
}

func (usecase *TeacherUseCase) TeacherRegister(domain *Domain, ctx context.Context) (Domain, error) {
	if domain.Name == "" {
		return Domain{}, errors.New("namd empty")
	}
	if domain.Email == "" {
		return Domain{}, errors.New("email empty")
	}
	if domain.Password == "" {
		return Domain{}, errors.New("password empty")
	}
	hashedPass := password.HashPassword(domain.Password)
	domain.Password = hashedPass
	std, err := usecase.repo.TeacherRegister(domain, ctx)
	if err != nil {
		return Domain{}, err
	}
	return std, nil
}

func (usecase *TeacherUseCase) TeacherUpdate(ctx context.Context, domain Domain, id uint) (Domain, error) {
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
	std, err := usecase.repo.TeacherUpdate(ctx, domain, id)
	if err != nil {
		return Domain{}, err
	}
	return std, nil
}

func (usecase *TeacherUseCase) TeacherLogin(domain Domain, ctx context.Context) (Domain, error) {
	if domain.Email == "" {
		return Domain{}, errors.New("email empty")
	}
	if domain.Password == "" {
		return Domain{}, errors.New("password empty")
	}
	std, err := usecase.repo.TeacherLogin(domain, ctx)
	if err != nil {
		return Domain{}, err
	}
	return std, nil

}

func (usecase *TeacherUseCase) TeacherGetProfile(ctx context.Context, id uint) (Domain, error) {
	std, err := usecase.repo.TeacherGetProfile(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if std.Id == 0 {
		return Domain{}, err
	}
	return std, nil
}
