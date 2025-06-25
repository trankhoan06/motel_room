package storage

import (
	"context"
	"main.go/modules/notify/model"
)

func (s *SqlModel) FindNotify(ctx context.Context, cond map[string]interface{}) (*model.Notify, error) {
	var res model.Notify
	if err := s.db.Where(cond).First(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}
