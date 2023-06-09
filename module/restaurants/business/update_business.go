package restaurantbusiness

import (
	"context"

	restaurantmodel "github.com/khoaphungnguyen/food_delivery/module/restaurants/model"
)

type UpdateRestaurantStore interface {
	Update(
		context context.Context,
		data *restaurantmodel.RestaurantUpdate,
		id int,
	) error
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{
		store: store,
	}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(context context.Context,
	data *restaurantmodel.RestaurantUpdate,
	id int) error {

	if err := biz.store.Update(context, data, id); err != nil {
		return err
	}

	return nil
}
