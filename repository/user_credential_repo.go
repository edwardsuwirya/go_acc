package repository

import (
	"enigmacamp.com/goacc/delivery/appreq"
)

type UserCredentialRepo interface {
	GetByUserNameAndPassword(user appreq.AuthRequest) error
}
