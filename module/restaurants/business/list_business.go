package restaurantbusiness

import (
	"context"

	"github.com/khoaphungnguyen/food_delivery/common"
	restaurantmodel "github.com/khoaphungnguyen/food_delivery/module/restaurants/model"
)

type ListRestaurantStore interface {
	List(
		context context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
	) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBiz struct {
	store ListRestaurantStore
}

func NewListRestaurantBiz(store ListRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{
		store: store,
	}
}

func (biz *listRestaurantBiz) ListRestaurant(context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.List(context, filter, paging)
	if err != nil {
		return nil, err
	}

	return result, nil
}
