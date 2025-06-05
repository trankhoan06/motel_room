package biz

import (
	"context"
	"errors"
	"main.go/common"
	modelRent "main.go/modules/rent/model"
	"main.go/modules/room_reviews/model"
)

func (biz *UserReviewsCommon) NewDeleteReview(ctx context.Context, data *model.DeletedReviews) error {
	rent, err := biz.rate.FindRent(ctx, map[string]interface{}{"id": data.RentId})
	if err != nil {
		return common.ErrNotFoundRent(errors.New("rent has been deleted or exist"))
	}
	review, err := biz.review.FindReview(ctx, map[string]interface{}{"id": data.Id, "rent_id": data.RentId})
	if err != nil {
		return common.ErrReview(errors.New(err.Error()))
	}
	if review.UserId != data.UserId {
		return common.ErrPermission(errors.New("you can't deleted the other people's review"))
	}
	if !review.IsOwner {
		var rate modelRent.RateRent
		rate.Id = review.RentId
		rate.AmountRate = rent.AmountRate - 1
		rate.Rate = (rent.Rate*float64(rent.AmountRate) - review.Rate) / float64(rate.AmountRate)
		if err := biz.rate.UpdateRate(ctx, &rate); err != nil {
			return err
		}
	}
	if err := biz.review.DeletedReview(ctx, map[string]interface{}{"id": data.Id, "rent_id": data.RentId}); err != nil {
		return err
	}
	return nil
}
