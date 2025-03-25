package storage

import "context"

func (s *SqlModel) ChangePassword(ctx context.Context, id int, password string) error {
	if err := s.db.Table("user").Where("id=?", id).Update("password", password).Error; err != nil {
		return err
	}
	return nil
}
