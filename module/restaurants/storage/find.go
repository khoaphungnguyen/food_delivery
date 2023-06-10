package restaurantstorage

import (
	"github.com/khoaphungnguyen/food_delivery/common"
	restaurantmodel "github.com/khoaphungnguyen/food_delivery/module/restaurants/model"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

func (s *sqlStore) FindDatatWithCondition(
	context context.Context,
	condition map[string]interface{},
	moreKey ...string,
) (*restaurantmodel.Restaurant, error) {
	var data restaurantmodel.Restaurant

	if err := s.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
