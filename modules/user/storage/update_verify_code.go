package storage

import "context"

func (s *SqlModel) UpdateVerifyCode(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table("send_code").Where(cond).Update("verify", 1).Error; err != nil {
		return err
	}
	return nil
}
