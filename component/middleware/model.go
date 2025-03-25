package middleware

import (
	"main.go/component/tokenprovider"
	"main.go/modules/user/storage"
)

type ModelMiddleware struct {
	authen *storage.SqlModel
	token  tokenprovider.TokenProvider
}

func NewModelMiddleware(au *storage.SqlModel, token tokenprovider.TokenProvider) *ModelMiddleware {
	return &ModelMiddleware{authen: au, token: token}
}
