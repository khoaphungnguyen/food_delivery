package ginrestaurant

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khoaphungnguyen/food_delivery/common"
	"github.com/khoaphungnguyen/food_delivery/component/appcontext"
	restaurantbusiness "github.com/khoaphungnguyen/food_delivery/module/restaurants/business"
	restaurantmodel "github.com/khoaphungnguyen/food_delivery/module/restaurants/model"
	restaurantstorage "github.com/khoaphungnguyen/food_delivery/module/restaurants/storage"
)

func UpdateRestaurants(appCtx appcontext.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data restaurantmodel.RestaurantUpdate
		if err := c.ShouldBindJSON(&data); err != nil {
			panic(err)
		}
		store := restaurantstorage.NewSQlStore(db)
		biz := restaurantbusiness.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurant(c.Request.Context(), &data, int(uid.GetLocalID())); err != nil {
			panic(err)
		}
		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(&data))
	}
}
