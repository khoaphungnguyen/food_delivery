package restaurantbusiness

import (
	"context"
	"errors"

	restaurantmodel "github.com/khoaphungnguyen/food_delivery/module/restaurants/model"
)

type DeleteRestaurantStore interface {
	FindDatatWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (*restaurantmodel.Restaurant, error)
	Delete(context context.Context, id int) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{
		store: store,
	}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(context context.Context, id int) error {
	oldData, err := biz.store.FindDatatWithCondition(context, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if oldData.Status == 0 {
		return errors.New("restaurant has been deleted")
	}
	if err := biz.store.Delete(context, id); err != nil {
		return err
	}
	return nil
}
