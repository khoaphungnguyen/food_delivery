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

func ListRestaurants(appCtx appcontext.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.Fulfill()

		var filter restaurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter.Status = []int{1}

		store := restaurantstorage.NewSQlStore(db)
		biz := restaurantbusiness.NewListRestaurantBiz(store)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &pagingData)
		if err != nil {
			panic(err)
		}
		for i := range result {
			result[i].Mask(false)
		}
		c.JSON(http.StatusOK, common.NewSuccessReponse(result, pagingData, filter))
	}
}
