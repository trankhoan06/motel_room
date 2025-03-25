package storage

import (
	"context"
	"main.go/modules/user/model"
)

func (s *SqlModel) FindCodeVerify(ctx context.Context, cond map[string]interface{}) (*model.VerifyAccount, error) {
	var verify model.VerifyAccount
	if err := s.db.Where(cond).Last(&verify).Error; err != nil {
		return nil, err
	}
	return &verify, nil
}
