package storage

import (
	"context"
	"main.go/modules/email/model"
)

func (s *SqlModel) UpdateSendCodeEmail(ctx context.Context, data *model.CreateVerifyAccount) error {
	if err := s.db.Where("email=?", data.Email).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
