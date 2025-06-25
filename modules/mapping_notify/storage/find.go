package storage

import (
	"context"
	"main.go/modules/mapping_notify/model"
)

func (s *SqlModel) FindMappingNotify(ctx context.Context, cond map[string]interface{}) (*model.MappingNotify, error) {
	var res model.MappingNotify
	db := s.db.Preload("Notify")
	if err := db.Where(cond).First(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}
