package restaurantstorage

import (
	restaurantmodel "github.com/khoaphungnguyen/food_delivery/module/restaurants/model"
	"golang.org/x/net/context"
)

func (s *sqlStore) Update(
	context context.Context,
	data *restaurantmodel.RestaurantUpdate,
	id int,
) error {

	if err := s.db.
		Where("id =?", id).
		Updates(&data).Error; err != nil {
		return err
	}
	return nil
}
