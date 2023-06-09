package restaurantstorage

import (
	restaurantmodel "github.com/khoaphungnguyen/food_delivery/module/restaurants/model"
	"golang.org/x/net/context"
)

func (s *sqlStore) FindDatatWithCondition(
	context context.Context,
	condition map[string]interface{},
	moreKey ...string,
) (*restaurantmodel.Restaurant, error) {
	var data restaurantmodel.Restaurant

	if err := s.db.Where(condition).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
