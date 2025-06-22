package storage

import (
	"context"
	"gorm.io/gorm"
)

func (s *SqlModel) UpdateSearchTime(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table("search").Where(cond).Update("search_time", gorm.Expr("search_time + ?", 1)).Error; err != nil {
		return err
	}
	return nil

}
