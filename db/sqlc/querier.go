// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package generated

import (
	"context"
)

type Querier interface {
	AddLike(ctx context.Context, arg AddLikeParams) (Like, error)
	// add sale into sales by given sale_type_id, amount, date, user_id and on conflict replace
	AddSaleOrReplace(ctx context.Context, arg AddSaleOrReplaceParams) (Sale, error)
	AddSaleToBrand(ctx context.Context, arg AddSaleToBrandParams) (SalesBrand, error)
	CreateBranch(ctx context.Context, arg CreateBranchParams) error
	CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error)
	CreatePost(ctx context.Context, arg CreatePostParams) (Post, error)
	CreatePostImages(ctx context.Context, arg CreatePostImagesParams) error
	CreateSaleType(ctx context.Context, arg CreateSaleTypeParams) (int32, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (int32, error)
	CreateUserCode(ctx context.Context, arg CreateUserCodeParams) (int32, error)
	DeleteComment(ctx context.Context, id int32) error
	DeleteLike(ctx context.Context, arg DeleteLikeParams) error
	DeletePost(ctx context.Context, id int32) error
	DeleteSale(ctx context.Context, id int32) error
	DeleteUserAvatar(ctx context.Context, userID int32) error
	GetAuthCodeById(ctx context.Context, id int32) (UsersCode, error)
	GetBranchBrandGoalByGivenDateRange(ctx context.Context, arg GetBranchBrandGoalByGivenDateRangeParams) (int64, error)
	GetBranchBrands(ctx context.Context, branchID int32) ([]GetBranchBrandsRow, error)
	GetBranchById(ctx context.Context, id int32) (Branch, error)
	GetBranches(ctx context.Context) ([]Branch, error)
	GetBrands(ctx context.Context, arg GetBrandsParams) ([]GetBrandsRow, error)
	GetCommentById(ctx context.Context, id int32) (Comment, error)
	GetCommentsAndAuthorsByPostId(ctx context.Context, arg GetCommentsAndAuthorsByPostIdParams) ([]GetCommentsAndAuthorsByPostIdRow, error)
	GetPostById(ctx context.Context, id int32) (Post, error)
	GetPostLikedUsers(ctx context.Context, arg GetPostLikedUsersParams) ([]GetPostLikedUsersRow, error)
	GetPostLikesCount(ctx context.Context, postID int32) (int64, error)
	GetPostsAndPostAuthors(ctx context.Context, arg GetPostsAndPostAuthorsParams) ([]GetPostsAndPostAuthorsRow, error)
	GetSaleBrandBySaleId(ctx context.Context, saleID int32) (GetSaleBrandBySaleIdRow, error)
	// Assuming you also have a sales table as previously discussed.
	// Assuming you also have a sales table as previously discussed.
	// Join with relevant tables
	GetSaleSumByBranchByTypeByBrand(ctx context.Context, arg GetSaleSumByBranchByTypeByBrandParams) ([]GetSaleSumByBranchByTypeByBrandRow, error)
	// get the sales sums for a specific sales manager and each sale type within the given period.
	GetSaleSumByManagerByTypeByBrand(ctx context.Context, arg GetSaleSumByManagerByTypeByBrandParams) (GetSaleSumByManagerByTypeByBrandRow, error)
	GetSaleSumByUserIdBrandIdPeriodSaleTypeId(ctx context.Context, arg GetSaleSumByUserIdBrandIdPeriodSaleTypeIdParams) (int64, error)
	GetSaleTypeById(ctx context.Context, id int32) (SaleType, error)
	GetSalesByBrandId(ctx context.Context, brandID int32) ([]GetSalesByBrandIdRow, error)
	GetSalesByBrandIdAndUserId(ctx context.Context, arg GetSalesByBrandIdAndUserIdParams) ([]GetSalesByBrandIdAndUserIdRow, error)
	GetSalesTypes(ctx context.Context) ([]SaleType, error)
	GetUserBranch(ctx context.Context, id int32) (GetUserBranchRow, error)
	GetUserBrand(ctx context.Context, arg GetUserBrandParams) (int32, error)
	GetUserBrandGoal(ctx context.Context, arg GetUserBrandGoalParams) (int64, error)
	GetUserBrands(ctx context.Context, userID int32) ([]GetUserBrandsRow, error)
	GetUserById(ctx context.Context, id int32) (GetUserByIdRow, error)
	GetUserByPhone(ctx context.Context, phone string) (UserAvatarView, error)
	GetUserPostLike(ctx context.Context, arg GetUserPostLikeParams) (int32, error)
	GetUserRank(ctx context.Context, arg GetUserRankParams) (int64, error)
	// SELECT users for given brand ordered by ratio
	GetUsersOrderedByRatioForGivenBrand(ctx context.Context, arg GetUsersOrderedByRatioForGivenBrandParams) ([]GetUsersOrderedByRatioForGivenBrandRow, error)
	InsertUserBrandRatio(ctx context.Context, arg InsertUserBrandRatioParams) error
	ListPosts(ctx context.Context) ([]Post, error)
	UploadUserAvatar(ctx context.Context, arg UploadUserAvatarParams) error
}

var _ Querier = (*Queries)(nil)
