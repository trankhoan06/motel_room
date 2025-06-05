package biz

import (
	"context"
	modelRent "main.go/modules/rent/model"
	"main.go/modules/room_reviews/model"
)

type ReviewBiz interface {
	CreateReview(ctx context.Context, data *model.CreateReviews) error
	UpdateReview(ctx context.Context, data *model.UpdateReviews) error
	DeletedReview(ctx context.Context, cond map[string]interface{}) error
	FindReview(ctx context.Context, cond map[string]interface{}) (*model.Reviews, error)
	ListReview(ctx context.Context, cond map[string]interface{}) (*[]model.Reviews, error)
}
type RateRentBiz interface {
	UpdateRate(ctx context.Context, data *modelRent.RateRent) error
	FindRent(ctx context.Context, cond map[string]interface{}) (*modelRent.Rent, error)
}
type UserReviewsCommon struct {
	review ReviewBiz
	rate   RateRentBiz
}
type OwnerReviewsCommon struct {
	review ReviewBiz
}

func NewUserReviewsCommon(review ReviewBiz, rate RateRentBiz) *UserReviewsCommon {
	return &UserReviewsCommon{
		review: review,
		rate:   rate,
	}
}
func NewOwnerReviewsCommon(review ReviewBiz) *OwnerReviewsCommon {
	return &OwnerReviewsCommon{
		review: review,
	}
}
