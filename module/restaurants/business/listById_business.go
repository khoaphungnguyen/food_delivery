package restaurantbusiness

import (
	"context"

	restaurantmodel "github.com/khoaphungnguyen/food_delivery/module/restaurants/model"
)

type ListByIdRestaurantStore interface {
	ListById(
		context context.Context,
		data *restaurantmodel.Restaurant,
		id int,
	) (*restaurantmodel.Restaurant, error)
}

type listByIdRestaurantBiz struct {
	store ListByIdRestaurantStore
}

func NewListByIdRestaurantBiz(store ListByIdRestaurantStore) *listByIdRestaurantBiz {
	return &listByIdRestaurantBiz{
		store: store,
	}
}

func (biz *listByIdRestaurantBiz) ListByIdRestaurant(context context.Context,
	data *restaurantmodel.Restaurant,
	id int,
) (*restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListById(context, data, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}
