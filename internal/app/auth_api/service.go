package authapi

import (
	"context"

	desc "github.com/DimDimi4-and-fixiki/auth-back/internal/gen/go/api/auth_api"
)

type Implementation struct {
	desc.UnimplementedAuthServiceServer

	uc usecases
}

func NewAuthApiService(uc usecases) *Implementation {
	return &Implementation{
		uc: uc,
	}
}

type usecases interface {
	RegisterUser(ctx context.Context) error
}
