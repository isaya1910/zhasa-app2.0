// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package generated

import (
	"context"
	"time"
)

type Querier interface {
	AddLike(ctx context.Context, arg AddLikeParams) (Like, error)
	// add sale into sales by given sale_type_id, amount, date, sales_manager_id and on conflict replace
	AddSaleOrReplace(ctx context.Context, arg AddSaleOrReplaceParams) (Sale, error)
	CreateBranch(ctx context.Context, arg CreateBranchParams) error
	CreateBranchDirector(ctx context.Context, arg CreateBranchDirectorParams) (int32, error)
	CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error)
	CreatePost(ctx context.Context, arg CreatePostParams) (Post, error)
	CreatePostImages(ctx context.Context, arg CreatePostImagesParams) error
	CreateSaleType(ctx context.Context, arg CreateSaleTypeParams) (int32, error)
	CreateSalesManager(ctx context.Context, arg CreateSalesManagerParams) error
	CreateSalesManagerGoalByType(ctx context.Context, arg CreateSalesManagerGoalByTypeParams) error
	CreateUser(ctx context.Context, arg CreateUserParams) (int32, error)
	CreateUserCode(ctx context.Context, arg CreateUserCodeParams) (int32, error)
	DeleteComment(ctx context.Context, id int32) error
	DeleteLike(ctx context.Context, arg DeleteLikeParams) error
	DeletePost(ctx context.Context, id int32) error
	DeleteSaleById(ctx context.Context, id int32) (Sale, error)
	DeleteUserAvatar(ctx context.Context, userID int32) error
	EditSaleById(ctx context.Context, arg EditSaleByIdParams) (Sale, error)
	GetAuthCodeById(ctx context.Context, id int32) (UsersCode, error)
	GetBranchById(ctx context.Context, id int32) (Branch, error)
	GetBranchDirectorByBranchId(ctx context.Context, branchID int32) (BranchDirectorsView, error)
	GetBranchDirectorByUserId(ctx context.Context, userID int32) (BranchDirectorsView, error)
	GetBranchGoalByGivenDateRange(ctx context.Context, arg GetBranchGoalByGivenDateRangeParams) (int64, error)
	GetBranches(ctx context.Context) ([]Branch, error)
	GetCommentById(ctx context.Context, id int32) (Comment, error)
	GetCommentsAndAuthorsByPostId(ctx context.Context, arg GetCommentsAndAuthorsByPostIdParams) ([]GetCommentsAndAuthorsByPostIdRow, error)
	GetManagerSales(ctx context.Context, arg GetManagerSalesParams) ([]GetManagerSalesRow, error)
	GetManagerSalesByPeriod(ctx context.Context, arg GetManagerSalesByPeriodParams) ([]GetManagerSalesByPeriodRow, error)
	GetOrderedBranchesByGivenPeriod(ctx context.Context, arg GetOrderedBranchesByGivenPeriodParams) ([]GetOrderedBranchesByGivenPeriodRow, error)
	GetOrderedSalesManagers(ctx context.Context, arg GetOrderedSalesManagersParams) ([]GetOrderedSalesManagersRow, error)
	GetOrderedSalesManagersOfBranch(ctx context.Context, arg GetOrderedSalesManagersOfBranchParams) ([]GetOrderedSalesManagersOfBranchRow, error)
	GetPostById(ctx context.Context, id int32) (Post, error)
	GetPostLikedUsers(ctx context.Context, arg GetPostLikedUsersParams) ([]GetPostLikedUsersRow, error)
	GetPostLikesCount(ctx context.Context, postID int32) (int64, error)
	GetPostsAndPostAuthors(ctx context.Context, arg GetPostsAndPostAuthorsParams) ([]GetPostsAndPostAuthorsRow, error)
	GetRating(ctx context.Context, arg GetRatingParams) (GetRatingRow, error)
	GetSMGoal(ctx context.Context, arg GetSMGoalParams) (int64, error)
	GetSMRatio(ctx context.Context, arg GetSMRatioParams) (float64, error)
	GetSaleTypeById(ctx context.Context, id int32) (SaleType, error)
	GetSalesByDate(ctx context.Context, saleDate time.Time) ([]Sale, error)
	GetSalesCount(ctx context.Context, salesManagerID int32) (int64, error)
	GetSalesManagerByUserId(ctx context.Context, userID int32) (SalesManagersView, error)
	GetSalesManagerGoalByGivenDateRangeAndSaleType(ctx context.Context, arg GetSalesManagerGoalByGivenDateRangeAndSaleTypeParams) (int64, error)
	// get the sales sums for a specific sales manager and each sale type within the given period.
	GetSalesManagerSumsByType(ctx context.Context, arg GetSalesManagerSumsByTypeParams) (GetSalesManagerSumsByTypeRow, error)
	GetSalesTypes(ctx context.Context) ([]SaleType, error)
	GetUserById(ctx context.Context, id int32) (UserAvatarView, error)
	GetUserByPhone(ctx context.Context, phone string) (UserAvatarView, error)
	GetUserPostLike(ctx context.Context, arg GetUserPostLikeParams) (int32, error)
	ListPosts(ctx context.Context) ([]Post, error)
	SetSMRatio(ctx context.Context, arg SetSMRatioParams) error
	SetSmGoalBySaleType(ctx context.Context, arg SetSmGoalBySaleTypeParams) error
	UploadUserAvatar(ctx context.Context, arg UploadUserAvatarParams) error
}

var _ Querier = (*Queries)(nil)
