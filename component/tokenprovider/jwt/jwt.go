package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"main.go/common"
	"main.go/component/tokenprovider"
	"time"
)

type JwtProvider struct {
	Secret string
	Prefix string
}
type token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}
type MyClaim struct {
	Payload common.Payload `json:"payload"`
	jwt.StandardClaims
}

func NewJwtProvider(secret string, prefix string) *JwtProvider {
	return &JwtProvider{Secret: secret, Prefix: prefix}
}
func (t *token) GetToken() string {
	return t.GetToken()
}
func (j *JwtProvider) GetSecret() string {
	return j.Secret
}
func (j *JwtProvider) Generate(payload tokenprovider.Payload, expiry int) (tokenprovider.Token, error) {
	now := time.Now()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &MyClaim{
		Payload: common.Payload{
			UId:  payload.GetUser(),
			Role: payload.GetRole(),
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(time.Duration(expiry) * time.Hour).Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    fmt.Sprint(now.UnixNano()),
		},
	})
	MyToken, err := t.SignedString([]byte(j.Secret))
	if err != nil {
		return nil, err
	}
	return &token{
		Token:   MyToken,
		Expiry:  expiry,
		Created: now,
	}, nil
}
func (j *JwtProvider) Validate(token string) (tokenprovider.Payload, error) {
	MyToken, err := jwt.ParseWithClaims(token, &MyClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil

	})
	if err != nil {
		return nil, err
	}
	if MyToken == nil {
		return nil, errors.New("token is nil")
	}
	if !MyToken.Valid {
		return nil, errors.New("token is invalid")
	}
	claim, ok := MyToken.Claims.(*MyClaim)
	if !ok {
		return nil, errors.New("token is invalid")
	}
	return &claim.Payload, nil
}
