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

func CreateRestaurants(appCtx appcontext.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data restaurantmodel.RestaurantCreate
		if err := c.ShouldBindJSON(&data); err != nil {
			panic(err)
		}
		store := restaurantstorage.NewSQlStore(db)
		biz := restaurantbusiness.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(false)
		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(data.FakedId.String()))
	}
}
