package biz

import (
	"context"
	"errors"
	"main.go/common"
	modelRent "main.go/modules/rent/model"
	"main.go/modules/room_reviews/model"
)

func (biz *UserReviewsCommon) NewUpdateReview(ctx context.Context, data *model.UpdateReviews) error {
	//tim rent
	rent, err := biz.rate.FindRent(ctx, map[string]interface{}{"id": data.RentId})
	if err != nil {
		return common.ErrNotFoundRent(errors.New("rent has been deleted or exist"))
	}
	//tim review
	re, err := biz.review.FindReview(ctx, map[string]interface{}{"id": data.Id})
	if err != nil {
		return err
	}
	if re.UserId != data.UserId {
		return common.ErrPermission(errors.New("you can't update this other review"))
	}

	// owner

	// user
	//update rate
	if re.UserId != rent.UserId {
		var updateRate modelRent.RateRent
		updateRate.Id = data.RentId
		if rent.AmountRate == 1 {
			updateRate.Rate = *data.Rate
		} else {
			updateRate.Rate = (rent.Rate*float64(rent.AmountRate) - re.Rate + *data.Rate) / float64(rent.AmountRate)
		}
		if err := biz.rate.UpdateRate(ctx, &updateRate); err != nil {
			return err
		}
	}
	if err := biz.review.UpdateReview(ctx, data); err != nil {
		return err
	}
	return nil
}
