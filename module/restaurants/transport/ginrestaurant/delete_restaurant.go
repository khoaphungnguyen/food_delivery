package ginrestaurant

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/khoaphungnguyen/food_delivery/common"
	"github.com/khoaphungnguyen/food_delivery/component/appcontext"
	restaurantbusiness "github.com/khoaphungnguyen/food_delivery/module/restaurants/business"
	restaurantstorage "github.com/khoaphungnguyen/food_delivery/module/restaurants/storage"
)

func DeleteRestaurants(appCtx appcontext.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		store := restaurantstorage.NewSQlStore(db)
		biz := restaurantbusiness.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
