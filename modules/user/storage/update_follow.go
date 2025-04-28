package storage

import (
	"context"
	"gorm.io/gorm"
)

func (s *SqlModel) UpdateFollow(ctx context.Context, cond map[string]interface{}, Expr, calculation string) error {
	if err := s.db.Table("user").Where(cond).Update(Expr, gorm.Expr(Expr+calculation+"?", 1)).Error; err != nil {
		return err
	}
	return nil
}
