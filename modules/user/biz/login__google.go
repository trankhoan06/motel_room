package biz

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/api/idtoken"
	"main.go/common"
	"main.go/component/tokenprovider"
	"main.go/modules/upload"
	"main.go/modules/user/model"
)

func (biz *LoginBiz) NewLoginGoogle(ctx context.Context, data *model.LoginMedia, expiry int) (*tokenprovider.Token, error) {
	// Kiểm tra nếu ID token không có trong yêu cầu
	if data.YourAccessToken == "" {
		return nil, errors.New("ID token is required")
	}

	// Kiểm tra tính hợp lệ của ID token với Google
	payload, err := idtoken.Validate(ctx, data.YourAccessToken, biz.cfg.Google.ClientID) // "YOUR_CLIENT_ID" là Client ID của ứng dụng
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to validate ID token: %s", err))
	}
	user, _ := biz.store.FindUser(ctx, map[string]interface{}{"email": payload.Claims["email"]})
	if user == nil {
		account := model.Register{
			Email:    payload.Claims["email"].(string),
			Salt:     common.GetSalt(50),
			FistName: payload.Claims["given_name"].(string),
			LastName: payload.Claims["name"].(string),
			Image:    payload.Claims["picture"].(*upload.Image),
			Phone:    payload.Claims["phone_number"].(string),
			Address:  payload.Claims["address"].(string),
			IsEMail:  payload.Claims["email_verified"].(bool),
		}
		if err := biz.store.CreateUser(ctx, &account); err != nil {
			return nil, err
		}
	}
	var payloadToken = &common.Payload{
		UId:  user.Id,
		Role: user.Role,
	}
	token, errToken := biz.provider.Generate(payloadToken, expiry)
	if errToken != nil {
		return nil, errToken
	}
	return &token, nil
}
