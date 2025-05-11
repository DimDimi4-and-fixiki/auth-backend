package auth

import "github.com/DimDimi4-and-fixiki/auth-back/internal/model"

type numVerifyClient interface {
	ValidatePhoneNumber(phoneNumber string) (*model.PhoneValidationResult, error)
}
