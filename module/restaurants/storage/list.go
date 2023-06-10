package restaurantstorage

import (
	"github.com/khoaphungnguyen/food_delivery/common"
	restaurantmodel "github.com/khoaphungnguyen/food_delivery/module/restaurants/model"
	"golang.org/x/net/context"
)

func (s *sqlStore) List(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	var data []restaurantmodel.Restaurant

	db := s.db.Table(restaurantmodel.Restaurant{}.TableName())

	if f := filter; f != nil {
		if f.OwnerId > 0 {
			db = db.Where("owner_id =?", f.OwnerId)
		}
		if len(f.Status) > 0 {
			db = db.Where("status in (?)", f.Status)
		}
	}
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := s.db.Offset((paging.Page - 1) * paging.Limit).
		Order("id desc").Limit(paging.Limit).
		Find(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return data, nil
}
