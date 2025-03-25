package tokenprovider

import "main.go/modules/user/model"

type TokenProvider interface {
	Generate(payload Payload, expiry int) (Token, error)
	Validate(token string) (Payload, error)
	GetSecret() string
}
type Token interface {
	GetToken() string
}
type Payload interface {
	GetRole() *model.RoleUser
	GetUser() int
}
