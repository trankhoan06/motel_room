package storage

import (
	"golang.org/x/net/context"
	"main.go/modules/rent/model"
)

func (s *SqlModel) FindRent(ctx context.Context, cond map[string]interface{}) (*model.Rent, error) {
	var data model.Rent
	if err := s.db.Table("rent").Where(cond).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
