package restaurantstorage

import (
	restaurantmodel "github.com/khoaphungnguyen/food_delivery/module/restaurants/model"
	"golang.org/x/net/context"
)

func (s *sqlStore) Delete(
	context context.Context,
	id int,
) error {

	if err := s.db.Table(restaurantmodel.Restaurant{}.
		TableName()).
		Where("id =?", id).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return err
	}
	return nil
}
