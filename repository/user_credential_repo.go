package repository

import "enigmacamp.com/goacc/model"

type UserCredentialRepo interface {
	GetByUserNameAndPassword(user model.UserCredential) error
}
