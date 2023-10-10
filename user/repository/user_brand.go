package repository

import (
	"context"
	"fmt"
	. "zhasa2.0/db/sqlc"
	"zhasa2.0/statistic/entities"
)

// UserBrandGoalFunc UserGoalFunc zero if goal is missing
type UserBrandGoalFunc func(params GetUserBrandGoalParams) int64

func NewUserGoalFunc(ctx context.Context, store UserBrandStore) UserBrandGoalFunc {
	return func(params GetUserBrandGoalParams) int64 {
		goal, err := store.GetUserBrandGoal(ctx, params)
		if err != nil {
			fmt.Println(err)
		}
		return goal
	}
}

type GetUserBrandFunc func(userId int32, brandId int32) (int32, error)

func NewGetUserBrandFunc(ctx context.Context, store UserBrandStore) GetUserBrandFunc {
	return func(userId int32, brandId int32) (int32, error) {
		userBrand, err := store.GetUserBrand(ctx, GetUserBrandParams{
			UserID:  userId,
			BrandID: brandId,
		})
		if err != nil {
			fmt.Println(err)
			return 0, err
		}
		return userBrand, err
	}
}

type UpdateUserBrandRatioFunc func(userId int32, brandId int32, ratio float64, period entities.Period) error

func NewUpdateUserBrandRatioFunc(ctx context.Context, store UserBrandStore) UpdateUserBrandRatioFunc {
	return func(userId int32, brandId int32, ratio float64, period entities.Period) error {
		from, to := period.ConvertToTime()

		err := store.InsertUserBrandRatio(ctx, InsertUserBrandRatioParams{
			UserID:   userId,
			BrandID:  brandId,
			Ratio:    float32(ratio),
			FromDate: from,
			ToDate:   to,
		})
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	}
}