package restaurantstorage

import (
	restaurantmodel "github.com/khoaphungnguyen/food_delivery/module/restaurants/model"
	"golang.org/x/net/context"
)

func (s *sqlStore) Create(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
