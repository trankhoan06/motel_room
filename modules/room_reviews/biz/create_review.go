package biz

import (
	"context"
	"errors"
	"main.go/common"
	modelRent "main.go/modules/rent/model"
	"main.go/modules/room_reviews/model"
)

func (biz *UserReviewsCommon) NewCreateReview(ctx context.Context, data *model.CreateReviews) error {
	rent, err := biz.rate.FindRent(ctx, map[string]interface{}{"id": data.RentId})
	if err != nil {
		return common.ErrNotFoundRent(errors.New("rent has been deleted or exist"))
	}
	data.IsOwner = (data.UserId == rent.UserId)
	if !data.IsOwner {
		if _, err := biz.review.FindReview(ctx, map[string]interface{}{"user_id": data.UserId, "rent_id": data.RentId}); err == nil {
			return errors.New("you has been rate in the rent")
		}
		if data.Rate > 5 || data.Rate < 0 {
			return errors.New("rate about 1-5")
		}
		if rent.UserId == data.UserId {
			return common.ErrPermission(errors.New("you can't review your rent"))
		}
		if data.ParentId != 0 {
			return common.ErrPermission(errors.New("you can't response this because you don't owner"))
		}
		var updateRate modelRent.RateRent
		updateRate.Id = data.RentId
		if rent.AmountRate == 0 {
			updateRate.AmountRate = 1
			updateRate.Rate = data.Rate
		} else {
			updateRate.Rate = (rent.Rate*float64(rent.AmountRate) + data.Rate) / float64(rent.AmountRate+1)
			updateRate.AmountRate = rent.AmountRate + 1
		}
		if err := biz.rate.UpdateRate(ctx, &updateRate); err != nil {
			return err
		}
	} else {
		if data.ParentId == 0 {
			return common.ErrPermission(errors.New("you need to response a user"))
		}
		data.Rate = 0
		if _, err := biz.review.FindReview(ctx, map[string]interface{}{"id": data.ParentId}); err != nil {
			return common.ErrPermission(errors.New("the review don't exist"))
		}
	}
	if err := biz.review.CreateReview(ctx, data); err != nil {
		return err
	}
	return nil
}
