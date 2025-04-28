package biz

import "context"

func (biz *UserCommonBiz) NewDeletedAccount(ctx context.Context, cond map[string]interface{}) error {
	if err := biz.store.DeletedUser(ctx, cond); err != nil {
		return err
	}
	return nil
}
