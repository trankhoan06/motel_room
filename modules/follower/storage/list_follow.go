package storage

import (
	"context"
	"main.go/modules/follower/model"
)

func (s *SqlModel) ListFollow(ctx context.Context, cond map[string]interface{}) (*[]model.Follower, error) {
	var list []model.Follower
	db := s.db.Preload("Owner")
	if err := db.Table("follow").Where(cond).Find(&list).Error; err != nil {
		return nil, err
	}
	return &list, nil
}
