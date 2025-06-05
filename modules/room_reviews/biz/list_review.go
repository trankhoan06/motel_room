package biz

import (
	"errors"
	"golang.org/x/net/context"
	"main.go/common"
	"main.go/modules/room_reviews/model"
)

func (biz *UserReviewsCommon) ArrChild(ctx context.Context, treeReviews []*common.TreeReview) {
	for _, review := range treeReviews {
		data := review.Val.(model.Reviews)
		arrChild, err := biz.review.ListReview(ctx, map[string]interface{}{"parent_id": data.Id, "rent_id": data.RentId})
		if err == nil {
			for _, arr := range *arrChild {
				*review.Child = append(*review.Child, *common.NewTreeReview(arr))

			}
		}
	}
}
func (biz *UserReviewsCommon) NewListReview(ctx context.Context, rentId int) ([]*common.TreeReview, error) {
	_, err := biz.rate.FindRent(ctx, map[string]interface{}{"id": rentId})
	if err != nil {
		return nil, common.ErrNotFoundRent(errors.New("rent has been deleted or exist"))
	}
	data, err := biz.review.ListReview(ctx, map[string]interface{}{"rent_id": rentId, "parent_id": 0})
	if err != nil {
		return nil, err
	}
	var treeReviews []*common.TreeReview
	for _, child := range *data {
		treeReviews = append(treeReviews, common.NewTreeReview(child))
	}
	biz.ArrChild(ctx, treeReviews)
	return treeReviews, nil
}
