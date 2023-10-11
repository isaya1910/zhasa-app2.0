package api

import (
	"github.com/gin-gonic/gin"
	"zhasa2.0/api/entities/sm"
	"zhasa2.0/base"
	"zhasa2.0/statistic"
)

type GetOrderedUsersRequest struct {
	BranchId *int32 `json:"branch_id" form:"branch_id"`
	BrandId  int32  `json:"brand_id" form:"brand_id" binding:"required"`
	Month    int32  `json:"month" form:"month" binding:"required"`
	Year     int32  `json:"year" form:"year" binding:"required"`
	Page     int32  `json:"page" form:"page"`
	Limit    int32  `json:"limit" form:"limit"`
}

func (server *Server) GetOrderedUsers(ctx *gin.Context) {
	var request GetOrderedUsersRequest
	err := ctx.ShouldBindQuery(&request)
	if err != nil {
		ctx.JSON(400, errorResponse(err))
		return
	}
	monthPeriod := statistic.MonthPeriod{
		MonthNumber: request.Month,
		Year:        request.Year,
	}
	users, err := server.getUsersOrderedByRatioForGivenBrandFunc(request.BrandId, monthPeriod, base.Pagination{
		Page:     request.Page,
		PageSize: request.Limit,
	})

	nextUsers, err := server.getUsersOrderedByRatioForGivenBrandFunc(request.BrandId, monthPeriod, base.Pagination{
		Page:     request.Page + 1,
		PageSize: request.Limit,
	})

	if err != nil {
		ctx.JSON(400, errorResponse(err))
		return
	}

	response := sm.SalesManagerRatingItemsResponse{
		Result:  sm.SalesManagerRatingItemsFrom(users),
		Count:   int32(len(users)),
		HasNext: len(nextUsers) > 0,
	}
	ctx.JSON(200, response)

}