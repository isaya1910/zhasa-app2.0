package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zhasa2.0/api/entities"
)

type UserBrandYearStatisticRequestBody struct {
	UserId  int32 `form:"user_id" json:"user_id"`
	Year    int32 `form:"year" json:"year"`
	BrandId int32 `form:"brand_id" json:"brand_id"`
}

type BranchBrandYearStatisticRequestBody struct {
	BranchId int32 `form:"branch_id" json:"branch_id"`
	Year     int32 `form:"year" json:"year"`
	BrandId  int32 `form:"brand_id" json:"brand_id"`
}

func (server *Server) GetUserBrandYearStatistic(ctx *gin.Context) {
	// retrieve year statistic for user with given request body
	var requestBody UserBrandYearStatisticRequestBody
	if err := ctx.ShouldBindQuery(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// use server.saleRepo to retrieve statistic from db
	stats, err := server.saleRepo.GetUserBrandMonthlyYearStatistic(requestBody.Year, requestBody.BrandId, requestBody.BrandId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// convert stats to YearStatisticResultResponse and return
	response := entities.YearStatisticResultResponse{
		Result: make([]entities.YearStatisticResponse, 0),
	}
	for _, stat := range stats {
		response.Result = append(response.Result, entities.YearStatisticResponse{
			SaleType: entities.SaleTypeResponse{
				Id:        stat.SaleType.Id,
				Title:     stat.SaleType.Title,
				Color:     stat.SaleType.Color,
				ValueType: stat.SaleType.ValueType,
			},
			Month:  stat.Month,
			Amount: stat.Amount,
			Goal:   stat.Goal,
		})
	}
	ctx.JSON(http.StatusOK, response)
}

func (server *Server) getBranchMonthStatistic(ctx *gin.Context) {
	// retrieve month statistic for branch with given request body
	var requestBody BranchBrandYearStatisticRequestBody
	if err := ctx.ShouldBindQuery(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

}