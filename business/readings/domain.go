package readings

import (
	"backend/business/modules"
	"context"
	"time"
)

type Domain struct {
	Id       uint
	ModuleId uint
	Module   modules.Domain
	Title    string
	Content  string
	Order    int
	CreateAt time.Time
	UpdateAt time.Time
}

type ReadingsUseCaseInterface interface {
	ReadingsAdd(ctx context.Context, domain Domain) (Domain, error)
	ReadingsUpdate(ctx context.Context, domain Domain, id uint) (Domain, error)
	ReadingsDelete(ctx context.Context, id uint) error
	ReadingsGetByModuleId(ctx context.Context, moduleId uint) ([]Domain, error)
}

type ReadingsRepoInterface interface {
	ReadingsAdd(ctx context.Context, domain Domain) (Domain, error)
	CheckModule(ctx context.Context, id uint) (modules.Domain, error)
	ReadingsUpdate(ctx context.Context, domain Domain, id uint) (Domain, error)
	ReadingsDelete(ctx context.Context, id uint) error
	ReadingsGetByModuleId(ctx context.Context, moduleId uint) ([]Domain, error)
}
