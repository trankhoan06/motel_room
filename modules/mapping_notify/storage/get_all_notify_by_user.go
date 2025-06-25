package storage

import (
	"context"
	"main.go/modules/mapping_notify/model"
)

func (s *SqlModel) GetAllNotifyByUser(ctx context.Context, userId int) (*[]model.MappingNotify, error) {
	var res []model.MappingNotify
	db := s.db.Preload("Notify")
	if err := db.Table("mapping_notify").Where("user_id=?", userId).Find(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}
