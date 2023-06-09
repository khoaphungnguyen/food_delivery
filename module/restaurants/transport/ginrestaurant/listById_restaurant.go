package ginrestaurant

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/khoaphungnguyen/food_delivery/common"
	"github.com/khoaphungnguyen/food_delivery/component/appcontext"
	restaurantbusiness "github.com/khoaphungnguyen/food_delivery/module/restaurants/business"
	restaurantmodel "github.com/khoaphungnguyen/food_delivery/module/restaurants/model"
	restaurantstorage "github.com/khoaphungnguyen/food_delivery/module/restaurants/storage"
)

func ListByIdRestaurants(appCtx appcontext.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data restaurantmodel.Restaurant
		db := appCtx.GetMainDBConnection()
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		store := restaurantstorage.NewSQlStore(db)
		biz := restaurantbusiness.NewListByIdRestaurantBiz(store)

		result, err := biz.ListByIdRestaurant(c.Request.Context(), &data, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
