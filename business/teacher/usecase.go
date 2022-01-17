package teacher

import (
	"backend/helper/err"
	"backend/helper/password"
	"context"
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
	std, err := usecase.repo.TeacherRegister(domain, ctx)
	if err != nil {
		return Domain{}, err
	}
	return std, nil
}

func (usecase *TeacherUseCase) TeacherUpdate(ctx context.Context, domain Domain, id uint) (Domain, error) {
	//put in top cause high priority (check pass before anything)
	if domain.Password != "" {
		data, err1 := usecase.repo.TeacherGetProfile(ctx, id)
		if err1 != nil {
			return Domain{}, err1
		}
		if !password.CheckSamePassword(domain.Password, data.Password) {
			return Domain{}, err.ErrWrongPassword
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
	if domain.BackGround == "" {
		return Domain{}, err.ErrBackGroundEmpty
	}
	if domain.NewPassword != "" {
		hashedPass := password.HashPassword(domain.Password)
		domain.Password = hashedPass
	}

	std, err := usecase.repo.TeacherUpdate(ctx, domain, id)
	if err != nil {
		return Domain{}, err
	}
	return std, nil
}

func (usecase *TeacherUseCase) TeacherLogin(domain Domain, ctx context.Context) (Domain, error) {
	if domain.Email == "" {
		return Domain{}, err.ErrEmailEmpty
	}
	if domain.Password == "" {
		return Domain{}, err.ErrPasswordEmpty
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
	return std, nil
}
