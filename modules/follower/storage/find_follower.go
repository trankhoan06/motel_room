package storage

import (
	"context"
	"main.go/modules/follower/model"
)

func (s *SqlModel) FindFollower(ctx context.Context, cond map[string]interface{}) (*model.Follower, error) {
	var follow model.Follower
	db := s.db.Preload("Owner")
	if err := db.Table("follow").Where(cond).First(&follow).Error; err != nil {
		return nil, err
	}
	return &follow, nil
}
