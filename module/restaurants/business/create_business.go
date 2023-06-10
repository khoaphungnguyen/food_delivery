package restaurantbusiness

import (
	"context"

	"github.com/khoaphungnguyen/food_delivery/common"
	restaurantmodel "github.com/khoaphungnguyen/food_delivery/module/restaurants/model"
)

type CreateRestaurantStore interface {
	Create(context.Context, *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{
		store: store,
	}
}

func (biz *createRestaurantBiz) CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	if err := biz.store.Create(context, data); err != nil {
		return common.ErrCannotCreateEntity(restaurantmodel.EntityName, err)
	}

	return nil
}
