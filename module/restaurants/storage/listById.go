package restaurantstorage

import (
	"github.com/khoaphungnguyen/food_delivery/common"
	restaurantmodel "github.com/khoaphungnguyen/food_delivery/module/restaurants/model"
	"golang.org/x/net/context"
)

func (s *sqlStore) ListById(
	context context.Context,
	data *restaurantmodel.Restaurant,
	id int,
) (*restaurantmodel.Restaurant, error) {
	if err := s.db.Where("id = ?", id).First(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return data, nil
}
