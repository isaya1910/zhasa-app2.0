package api

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"zhasa2.0/apiadmin"
	. "zhasa2.0/branch/repository"
	. "zhasa2.0/branch_director/repo"
	. "zhasa2.0/brand"
	. "zhasa2.0/brand/repository"
	. "zhasa2.0/db/hand-made"
	generated "zhasa2.0/db/sqlc"
	. "zhasa2.0/news/repository"
	. "zhasa2.0/owner/repository"
	"zhasa2.0/rating"
	. "zhasa2.0/sale/repository"
	. "zhasa2.0/statistic/repository"
	"zhasa2.0/user/entities"
	. "zhasa2.0/user/repository"
	"zhasa2.0/user/service"
)

type Server struct {
	router      *gin.Engine
	environment string
	apiadmin.Server
	tokenService                                  service.TokenService
	authService                                   service.AuthorizationService
	rankRepo                                      RankingsRepository
	postRepository                                PostRepository
	ownerRepository                               OwnerRepository
	saleTypeRepo                                  SaleTypeRepository
	directorRepo                                  BranchDirectorRepository
	saleRepo                                      SaleRepository
	userBrandGoal                                 UserBrandGoalFunc
	getUserBrandFunc                              GetUserBrandFunc
	updateUserBrandRatio                          UpdateUserBrandRatioFunc
	getUserRatingFunc                             rating.GetUserRatingFunc
	getUserBranchFunc                             GetUserBranchFunc
	calculateUserBrandRatio                       CalculateUserBrandRatio
	getBranchBrands                               GetBranchBrandsFunc
	getAllBrands                                  GetAllBrandsFunc
	getUserBrands                                 GetUserBrandsFunc
	getBranchBrandFunc                            GetBranchBrandFunc
	getBranchBrandSaleSumFunc                     GetBranchBrandSaleSumFunc
	getBranchBrandGoalFunc                        GetBranchBrandGoalFunc
	getUsersOrderedByRatioForGivenBrandFunc       GetUsersOrderedByRatioForGivenBrandFunc
	getBranchUsersOrderedByRatioForGivenBrandFunc GetBranchUsersOrderedByRatioForGivenBrandFunc
	getBranchByIdFunc                             GetBranchByIdFunc
	getBranchesByBrandFunc                        GetBranchesByBrandFunc
	setBranchBrandSaleTypeGoal                    SetBranchBrandSaleTypeGoal
	setUserBrandGoalRequest                       SetUserBrandSaleTypeGoalFunc
	getUserByBranchBrandRoleFunc                  GetUserByBranchBrandRoleFunc
	getBranchBrandMonthlyYearStatisticFunc        GetBranchBrandMonthlyYearStatisticFunc
	getUsersByBranchBrandRoleFunc                 GetUsersByBranchBrandRoleFunc
	getSaleSumByUserBrandTypePeriodFunc           GetSaleSumByUserBrandTypePeriodFunc
	salesByBrandUserFunc                          SalesByBrandUserFunc
	saleAddFunc                                   SaleAddFunc
	saleEditFunc                                  SaleEditFunc
	ratedBranchesFunc                             RatedBranchesFunc
	setBrandSaleTypeGoal                          SetBrandSaleTypeGoalFunc
	getBrandSaleSumFunc                           GetBrandSaleSumFunc
	getBrandOverallGoalFunc                       GetBrandOverallGoalFunc
	getBrandMonthlyYearStatisticFunc              GetBrandMonthlyYearStatisticFunc

	// user functions
	createUserFunc                 CreateUserFunc
	getUserByPhoneFunc             GetUserByPhoneFunc
	getUserByPhoneWithPasswordFunc GetUserByPhoneWithPasswordFunc
	getUserByIdFunc                GetUserByIdFunc
	uploadAvatarFunc               UploadAvatarFunc
	deleteAvatarFunc               DeleteAvatarFunc
	checkDisabledUserFunc          CheckDisabledUserFunc
	searchUsersFunc                SearchUsersFunc
	updateUserProfileAbout         UpdateUserProfileAbout
}

func (server *Server) InitSuperUser() error {
	request := entities.CreateUserRequest{
		FirstName: "admin",
		LastName:  "admin",
		Phone:     "+77081070480",
	}
	_, err := server.getUserByPhoneFunc(request.Phone)

	if err == nil {
		fmt.Println("super user already exist")
		return nil
	}

	_, err = server.createUserFunc(request.FirstName, request.LastName, request.Phone)
	if err != nil {
		return err
	}
	fmt.Println("super user created")
	return nil
}

func corsHandler(h gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.Status(http.StatusOK)
			c.Header("Access-Control-Allow-Origin", "http://doschamp.doscar.kz")
			c.Header("Vary", "Origin")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			return
		}
		h(c)
	}
}

func NewServer(ctx context.Context, environment string) *Server {
	server := &Server{
		environment: environment,
	}
	initDependencies(server, ctx)

	router := gin.Default()
	//corsConfig := cors.DefaultConfig()
	//corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	//corsConfig.AllowAllOrigins = true
	config := cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"https://localhost:3000",
			"http://localhost:3001",
			"https://localhost:3001",
			"https://localhost:8080",
			"https://localhost:8081",

			"http://doschamp.doscar.kz",
			"https://doschamp.doscar.kz",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"X-Content-Type, Content-Length", "Content-Type", "Authorization", "Accept"},
		AllowCredentials: true,
	}

	router.Use(cors.New(config))
	router.LoadHTMLGlob("templates/*")

	router.POST("/image/avatar/upload", verifyToken(server.tokenService), server.HandleAvatarUpload)
	router.POST("/image/news/upload", verifyToken(server.tokenService), server.HandleNewsUpload)

	router.POST("/user/avatar", verifyToken(server.tokenService), server.UploadUserAvatar)
	router.DELETE("/user/avatar", verifyToken(server.tokenService), server.DeleteAvatar)
	router.DELETE("/user", verifyToken(server.tokenService), server.deleteAccount)
	router.POST("auth/signup", server.signup)
	router.POST("/csv/managers", verifyToken(server.tokenService), server.HandleManagersUpload)
	router.POST("/csv/directors", verifyToken(server.tokenService), server.HandleDirectorsUpload)

	router.GET("/create-user", server.GetUserForm)
	router.POST("/create-user", server.CreateUserFromForm)
	router.GET("/create-manager", server.CreateManagerForm)
	router.POST("/create-manager", server.CreateManagerFromForm)
	router.GET("/users/all", server.GetAllUsersForm)
	router.GET("/users/edit/:id", server.EditUserForm)
	router.POST("/users/edit/:id", server.PerformEditUserFromForm)
	router.GET("/users/disable/:id", server.DisableUserForm)

	router.GET("/policy", server.PrivacyPolicy)

	authRoute := router.Group("auth/")
	{
		authRoute.POST("/request-code", server.requestAuthCode)
		authRoute.POST("/try", server.tryAuth)
	}

	router.GET("user/get-user-profile", server.getUserProfile)

	router.POST("admin/login", server.AdminLogin)

	adminRoute := router.Group("admin/") //.Use(verifyToken(server.tokenService))
	{
		adminRoute.GET("/branches", corsHandler(server.GetAllBranches))
		adminRoute.GET("/brands", corsHandler(server.GetAllBrands))
		adminRoute.POST("/brand", corsHandler(server.CreateBrand))
		adminRoute.PUT("/brand", corsHandler(server.UpdateBrand))

		adminRoute.GET("/users", corsHandler(server.GetAllUsers))
		adminRoute.OPTIONS("/users", corsHandler(server.GetAllUsers))
		adminRoute.POST("/user", corsHandler(server.CreateUser))
		adminRoute.DELETE("/users", corsHandler(server.DeleteUsers))
		adminRoute.PUT("/users/activate", corsHandler(server.ActivateUsers))
		adminRoute.PUT("/update-user", corsHandler(server.UpdateUser))

		adminRoute.PUT("/change-users-role", corsHandler(server.ChangeUsersRole))
		adminRoute.PUT("/change-users-brands", corsHandler(server.ChangeUsersBrands))
		adminRoute.PUT("/change-users-branch", corsHandler(server.ChangeUsersBranch))

		adminRoute.POST("/branch", corsHandler(server.CreateBranchWithBrands))
		adminRoute.PUT("/branch", corsHandler(server.UpdateBranchWithBrands))
		adminRoute.GET("/sale-type/list", corsHandler(server.getSaleTypes))
	}

	smRoute := router.Group("sales-manager/")
	smRoute.GET("/year-statistic", server.GetUserBrandYearStatistic).Use(verifyToken(server.tokenService))
	smRoute.GET("/year-statistic-v2", server.GetUserSalesStatistics).Use(verifyToken(server.tokenService))
	smRoute.GET("/sale/list", server.GetSales).Use(verifyToken(server.tokenService))

	router.GET("branches", server.GetBranchesByBrand).Use(verifyToken(server.tokenService))
	branchRoute := router.Group("branch/").Use(verifyToken(server.tokenService))
	{
		branchRoute.GET("/year-statistic", server.GetBranchBrandYearStatistic)
		branchRoute.GET("/year-statistic-v2", server.GetBranchSalesStatistics)
		branchRoute.GET("/sales-managers", server.GetBranchSalesManagerList)
	}

	router.DELETE("sales/delete", server.DeleteSale).Use(verifyToken(server.tokenService))
	router.POST("sales-manager/sale/new", server.AddSale).Use(verifyToken(server.tokenService))
	router.POST("sales/edit", server.EditSale).Use(verifyToken(server.tokenService))

	directorRouter := router.Group("director/")
	{
		directorRouter.POST("sales-manager/goal", server.SetUserBrandGoal).Use(verifyToken(server.tokenService))
		directorRouter.GET("sales-manager/goal", server.GetSmGoal).Use(verifyToken(server.tokenService))
		directorRouter.POST("branch/goal", server.SetBranchGoal).Use(verifyToken(server.tokenService))
		directorRouter.GET("branch/goal", server.GetBranchGoal).Use(verifyToken(server.tokenService))
	}

	router.GET("sales-manager/dashboard", server.SMDashboard).Use(verifyToken(server.tokenService))

	router.GET("branch/dashboard", server.BranchDashboard).Use(verifyToken(server.tokenService))
	router.GET("rating/branches", server.GetRatedBranches).Use(verifyToken(server.tokenService))
	router.GET("rating/managers", server.GetOrderedUsers).Use(verifyToken(server.tokenService))

	router.GET("news", verifyToken(server.tokenService), server.GetPosts)
	router.POST("news/new", verifyToken(server.tokenService), server.CreatePost)
	router.DELETE("news/delete", verifyToken(server.tokenService), server.DeletePost)
	router.POST("news/like-toggle", verifyToken(server.tokenService), server.ToggleLike)

	router.GET("news/comments", verifyToken(server.tokenService), server.GetComments)
	router.POST("news/comments/new", verifyToken(server.tokenService), server.CreateComment)
	router.DELETE("news/comments/delete", verifyToken(server.tokenService), server.DeleteComment)

	router.GET("user/brands", verifyToken(server.tokenService), server.GetUserBrands)
	router.GET("branch/brands", verifyToken(server.tokenService), server.GetBranchBrands)
	router.GET("brands", verifyToken(server.tokenService), server.GetAllBrands)

	router.POST("owner/brand-goal", verifyToken(server.tokenService), server.SetOwnerDashboardGoal)
	router.GET("owner/brand-goal", verifyToken(server.tokenService), server.GetOwnerDashboardBySaleTypes)
	router.GET("owner/brand-goal-branches", verifyToken(server.tokenService), server.GetOwnerDashboardByBranches)
	router.GET("owner/goal", verifyToken(server.tokenService), server.GetOwnerGoal)
	router.GET("owner/year-statistic", verifyToken(server.tokenService), server.GetOwnerYearStatistic)

	router.GET("/users/search", server.SearchUsers)
	router.GET("/users", server.GetUser)
	router.PUT("/users/about", server.UpdateUserAbout)

	server.router = router
	return server
}

func initDependencies(server *Server, ctx context.Context) {
	dbDriver := os.Getenv("DB_DRIVER")
	dbSource := os.Getenv("DATA_BASE_URL")
	conn, err := sql.Open(dbDriver, dbSource)

	log.Println(dbDriver, dbSource)
	if err != nil {
		log.Panic("Cannot connect to db", err)
	}

	store := generated.NewStore(conn)
	customQuerier := NewCustomQuerier(conn)
	saleTypeRepo := NewSaleTypeRepository(ctx, store)
	branchRepo := NewBranchRepository(ctx, store, customQuerier, saleTypeRepo)
	directorRepo := NewBranchDirectorRepository(ctx, store)
	rankingsRepo := NewRankingsRepository(ctx, customQuerier, branchRepo)
	postRepo := NewPostRepository(ctx, store, customQuerier)
	ownerRepo := NewOwnerRepo(ctx, store)

	getUserByPhoneFunc := NewGetUserByPhoneFunc(ctx, store)
	getUserByPhoneWithPasswordFunc := NewGetUserByPhoneWithPasswordFunc(ctx, store)
	getUserByIdFunc := NewGetUserByIdFunc(ctx, store)
	addUserCodeFunc := NewAddUserCodeFunc(ctx, store)
	getAuthCodeByIdFunc := NewGetAuthCodeByIdFunc(ctx, store)
	checkDisabledUserFunc := NewCheckDisabledUserFunc(ctx, store)

	authService := service.NewAuthorizationService(
		ctx,
		getUserByPhoneFunc,
		getUserByPhoneWithPasswordFunc,
		addUserCodeFunc,
		getUserByIdFunc,
		getAuthCodeByIdFunc,
		checkDisabledUserFunc,
	)

	brandGoal := NewUserGoalFunc(ctx, store)
	userSaleSum := NewGetSaleSumByUserBrandTypePeriodFunc(ctx, store)
	saleRepo := NewSaleRepo(ctx, store, saleTypeRepo, brandGoal, userSaleSum)
	allBrands := NewGetAllBrandsFunc(ctx, store)
	encKey := []byte("YELLOW SUBMARINE, BLACK WIZARDRY")

	tokenService := service.NewTokenService(&encKey)
	server.tokenService = tokenService
	server.authService = authService
	server.postRepository = postRepo
	server.ownerRepository = ownerRepo
	server.directorRepo = directorRepo
	server.rankRepo = rankingsRepo
	server.saleRepo = saleRepo
	server.saleTypeRepo = saleTypeRepo
	server.userBrandGoal = brandGoal
	server.getUserBrandFunc = NewGetUserBrandFunc(ctx, store)
	server.updateUserBrandRatio = NewUpdateUserBrandRatioFunc(ctx, store)
	server.getUserRatingFunc = rating.NewGetUserRatingFunc(ctx, store)
	server.getUserBranchFunc = NewGetUserBranchFunc(ctx, store)
	server.calculateUserBrandRatio = NewCalculateUserBrandRatio(saleTypeRepo, userSaleSum, server.userBrandGoal)
	server.getBranchBrands = NewGetBranchBrandsFunc(ctx, store)
	server.getAllBrands = allBrands
	server.getUserBrands = NewGetUserBrandsFunc(ctx, store)
	server.getBranchBrandFunc = NewGetBranchBrand(ctx, store)
	server.getBranchBrandSaleSumFunc = NewGetBranchBrandSaleSumFunc(ctx, store)
	server.getBranchBrandGoalFunc = NewGetBranchBrandGoalFunc(ctx, store)
	server.getUsersOrderedByRatioForGivenBrandFunc = NewGetUsersOrderedByRatioForGivenBrandFunc(ctx, store)
	server.getBranchUsersOrderedByRatioForGivenBrandFunc = NewGetUsersOrderedBYRatioForGivenBrandAndBranchFunc(ctx, store)
	server.getBranchByIdFunc = NewGetBranchByIdFunc(ctx, store)
	server.getBranchesByBrandFunc = NewGetBranchesByBrandFunc(ctx, store)
	server.setBranchBrandSaleTypeGoal = NewSetBranchGoalFunc(ctx, store)
	server.setUserBrandGoalRequest = NewSetUserBrandSaleTypeGoalFunc(ctx, store)
	server.getUserByBranchBrandRoleFunc = NewGetUserByBranchBrandRoleFunc(ctx, store)
	server.getBranchBrandMonthlyYearStatisticFunc = NewGetBranchBrandMonthlyYearStatisticFunc(saleTypeRepo, server.getBranchBrandGoalFunc, server.getBranchBrandFunc, server.getBranchBrandSaleSumFunc)
	server.getUsersByBranchBrandRoleFunc = NewGetUsersByBranchBrandRoleFunc(ctx, store)
	server.getSaleSumByUserBrandTypePeriodFunc = userSaleSum
	server.salesByBrandUserFunc = NewSalesByBrandUserFunc(ctx, store)
	server.saleAddFunc = NewSaleAddFunc(ctx, store)
	server.saleEditFunc = NewSaleEditFunc(ctx, store)

	server.ratedBranchesFunc = NewRatedBranchesFunc(ctx, store, server.getBranchBrandSaleSumFunc, server.getBranchBrandGoalFunc)
	server.setBrandSaleTypeGoal = NewSetBrandSaleTypeGoalFunc(ctx, store)
	server.getBrandSaleSumFunc = NewGetBrandSaleSumFunc(ctx, store)
	server.getBrandOverallGoalFunc = NewGetBrandOverallGoalFunc(ctx, store)
	server.getBrandMonthlyYearStatisticFunc = NewGetBrandMonthlyYearStatisticFunc(server.saleTypeRepo, server.getBrandOverallGoalFunc, server.getBrandSaleSumFunc)

	// user functions
	server.createUserFunc = NewCreateUserFunc(ctx, store)
	server.getUserByPhoneFunc = NewGetUserByPhoneFunc(ctx, store)
	server.getUserByPhoneWithPasswordFunc = NewGetUserByPhoneWithPasswordFunc(ctx, store)
	server.getUserByIdFunc = NewGetUserByIdFunc(ctx, store)
	server.uploadAvatarFunc = NewUploadAvatarFunc(ctx, store)
	server.deleteAvatarFunc = NewDeleteAvatarFunc(ctx, store)
	server.searchUsersFunc = NewSearchUsersFunc(ctx, store)
	server.updateUserProfileAbout = NewUpdateUserProfileAbout(ctx, store)

	getUserByPhoneFunc = NewGetUserByPhoneFunc(ctx, store)
	createUserFunc := NewCreateUserFunc(ctx, store)
	makeUserAsManagerFunc := NewMakeUserAsManagerFunc(ctx, store)
	getUsersWithoutRolesFunc := NewGetUsersWithoutRolesFunc(ctx, store)
	getUsersByRoleFunc := NewGetUsersByRoleFunc(ctx, store)
	getUsersWithBranchBrands := NewGetUsersWithBranchBrands(ctx, store)
	getFilteredUsersWithBranchBrands := NewGetFilteredUsersWithBranchBrands(ctx, store)
	getUserByIdFunc = NewGetUserByIdFunc(ctx, store)
	updateUserBrandsFunc := NewUpdateUserBrandsFunc(ctx, store)
	updateUserFunc := NewUpdateUserFunc(ctx, store)
	updateUserBranchFunc := NewUpdateUserBranchFunc(ctx, store)
	getUserBranchFunc := NewGetUserBranchFunc(ctx, store)
	getAllBranches := NewGetAllBranchesFunc(ctx, store)
	getUserBrandsFunc := NewGetUserBrandsFunc(ctx, store)
	addDisabledUserFunc := NewAddDisabledUserFunc(ctx, store)
	addUserRoleFunc := NewAddUserRoleFunc(ctx, store)
	addUserBranchFunc := NewAddUserBranchFunc(ctx, store)
	updateUserRole := NewUpdateUserRoleFunc(ctx, store)
	createBranchWithBrands := NewCreateBranchWithBrandsFunc(ctx, store)
	updateBranchWithBrands := NewUpdateBranchWithBrandsFunc(ctx, store)
	createBrandFunc := NewCreateBrandFunc(ctx, store)
	updateBrandFunc := NewUpdateBrandFunc(ctx, store)
	getBranchesFilteredAsc := NewGetBranchesFilteredAsc(ctx, store)
	getBranchesFilteredDesc := NewGetBranchesFilteredDesc(ctx, store)
	getBranchesFilteredCount := NewGetBranchesFilteredCount(ctx, store)
	removeDisabledUsersFunc := NewRemoveDisabledUsersFunc(ctx, store)

	server.Server = *apiadmin.NewServer(
		authService,
		tokenService,
		getUserByPhoneFunc,
		createUserFunc,
		makeUserAsManagerFunc,
		getUsersWithoutRolesFunc,
		getUsersByRoleFunc,
		getUsersWithBranchBrands,
		getUserByIdFunc,
		getUserBranchFunc,
		updateUserBrandsFunc,
		updateUserFunc,
		updateUserBranchFunc,
		getAllBranches,
		createBranchWithBrands,
		updateBranchWithBrands,
		getBranchesFilteredAsc,
		getBranchesFilteredDesc,
		getBranchesFilteredCount,
		allBrands,
		addDisabledUserFunc,
		getUserBrandsFunc,
		getFilteredUsersWithBranchBrands,
		addUserRoleFunc,
		addUserBranchFunc,
		updateUserRole,
		createBrandFunc,
		updateBrandFunc,
		removeDisabledUsersFunc,
	)
}

// Start runs the HTTP server a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
