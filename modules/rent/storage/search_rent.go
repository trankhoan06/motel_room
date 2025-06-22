package storage

import (
	"context"
	modelRent "main.go/modules/rent/model"
	"main.go/modules/search/model"
)

func (s *SqlModel) SearchRent(ctx context.Context, data *model.SearchRent) (*[]modelRent.Rent, error) {
	query := s.db.Table("rent").Preload("Owner").Where("province=?", data.Sort.Province).Where("title LIKE ? or description LIKE ? or locate LIKE ?", "%"+data.Content+"%", "%"+data.Content+"%", "%"+data.Content+"%").
		Where("status<>?", modelRent.StatusDeletedRent).Offset(data.OffSet).Limit(data.Limit)
	var res []modelRent.Rent
	if data.Sort.TypeRoom != nil {
		query = query.Where("room_type = ?", data.Sort.TypeRoom)
	}

	if data.Sort.LowPrice > 0 && data.Sort.HighPrice > 0 {
		query = query.Where("price BETWEEN ? AND ?", data.Sort.LowPrice, data.Sort.HighPrice)
	}

	if data.Sort.LowArea > 0 && data.Sort.HighArea > 0 {
		query = query.Where("area BETWEEN ? AND ?", data.Sort.LowArea, data.Sort.HighArea)
	}

	if err := query.Find(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil

}
