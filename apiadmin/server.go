package apiadmin

import (
	"github.com/gin-gonic/gin"
	branchRepo "zhasa2.0/branch/repository"
	"zhasa2.0/brand"
	"zhasa2.0/user/repository"
	"zhasa2.0/user/service"
)

type Server struct {
	authService  service.AuthorizationService
	tokenService service.TokenService

	getUserByPhoneFunc       repository.GetUserByPhoneFunc
	getUsersWithoutRolesFunc repository.GetUsersWithoutRolesFunc
	createUserFunc           repository.CreateUserFunc
	makeUserAsManagerFunc    repository.MakeUserAsManagerFunc
	getUsersByRoleFunc       repository.GetUsersByRoleFunc
	getUsersWithBranchBrands repository.GetUsersWithBranchBrands
	getUserByIdFunc          repository.GetUserByIdFunc
	getUserBranchFunc        repository.GetUserBranchFunc
	updateUserBrands         repository.UpdateUserBrandsFunc
	updateUserFunc           repository.UpdateUserFunc
	updateUserBranchFunc     repository.UpdateUserBranchFunc
	updateUserRole           repository.UpdateUserRoleFunc
	addUserRole              repository.AddUserRoleFunc
	addUserBranch            repository.AddUserBranchFunc
	addDisabledUserFunc      repository.AddDisabledUserFunc
	removeDisabledUsersFunc  repository.RemoveDisabledUsersFunc

	getAllBranchesFunc         branchRepo.GetAllBranches
	getBranchesFilteredAsc     branchRepo.GetBranchesFilteredAsc
	getBranchesFilteredDesc    branchRepo.GetBranchesFilteredDesc
	getBranchesFilteredCount   branchRepo.GetBranchesFilteredCount
	createBranchWithBrandsFunc branchRepo.CreateBranchWithBrandsFunc
	updateBranchWithBrandsFunc branchRepo.UpdateBranchWithBrandsFunc

	getAllBrandsFunc                 brand.GetAllBrandsFunc
	getUserBrandsFunc                brand.GetUserBrandsFunc
	getFilteredUsersWithBranchBrands repository.GetFilteredUsersWithBranchBrands
	createBrandFunc                  brand.CreateBrandFunc
	updateBrandFunc                  brand.UpdateBrandFunc
}

func NewServer(
	authService service.AuthorizationService,
	tokenService service.TokenService,
	getUserByPhoneFunc repository.GetUserByPhoneFunc,
	createUserFunc repository.CreateUserFunc,
	makeManagerAsUserFunc repository.MakeUserAsManagerFunc,
	getUsersWithoutRolesFunc repository.GetUsersWithoutRolesFunc,
	getUsersByRoleFunc repository.GetUsersByRoleFunc,
	getUsersWithBranchBrands repository.GetUsersWithBranchBrands,
	getUserByIdFunc repository.GetUserByIdFunc,
	getUserBranchFunc repository.GetUserBranchFunc,
	updateUserBrands repository.UpdateUserBrandsFunc,
	updateUserFunc repository.UpdateUserFunc,
	updateUserBranchFunc repository.UpdateUserBranchFunc,
	branchesFunc branchRepo.GetAllBranches,
	createBranchWithBrandsFunc branchRepo.CreateBranchWithBrandsFunc,
	updateBranchWithBrandsFunc branchRepo.UpdateBranchWithBrandsFunc,
	getBranchesFilteredAsc branchRepo.GetBranchesFilteredAsc,
	getBranchesFilteredDesc branchRepo.GetBranchesFilteredDesc,
	getBranchesFilteredCount branchRepo.GetBranchesFilteredCount,
	brandsFunc brand.GetAllBrandsFunc,
	addDisabledUserFunc repository.AddDisabledUserFunc,
	getUserBrandsFunc brand.GetUserBrandsFunc,
	getFilteredUsersWithBranchBrands repository.GetFilteredUsersWithBranchBrands,
	addUserRole repository.AddUserRoleFunc,
	addUserBranch repository.AddUserBranchFunc,
	updateUserRole repository.UpdateUserRoleFunc,
	createBrandFunc brand.CreateBrandFunc,
	updateBrandFunc brand.UpdateBrandFunc,
	removeDisabledUsersFunc repository.RemoveDisabledUsersFunc,
) *Server {
	return &Server{
		authService:                      authService,
		tokenService:                     tokenService,
		getUserByPhoneFunc:               getUserByPhoneFunc,
		createUserFunc:                   createUserFunc,
		makeUserAsManagerFunc:            makeManagerAsUserFunc,
		getUsersWithoutRolesFunc:         getUsersWithoutRolesFunc,
		getUsersByRoleFunc:               getUsersByRoleFunc,
		getUsersWithBranchBrands:         getUsersWithBranchBrands,
		getUserByIdFunc:                  getUserByIdFunc,
		getUserBranchFunc:                getUserBranchFunc,
		updateUserBrands:                 updateUserBrands,
		updateUserFunc:                   updateUserFunc,
		updateUserBranchFunc:             updateUserBranchFunc,
		addDisabledUserFunc:              addDisabledUserFunc,
		getAllBranchesFunc:               branchesFunc,
		getAllBrandsFunc:                 brandsFunc,
		getUserBrandsFunc:                getUserBrandsFunc,
		getFilteredUsersWithBranchBrands: getFilteredUsersWithBranchBrands,
		addUserRole:                      addUserRole,
		addUserBranch:                    addUserBranch,
		updateUserRole:                   updateUserRole,
		createBranchWithBrandsFunc:       createBranchWithBrandsFunc,
		updateBranchWithBrandsFunc:       updateBranchWithBrandsFunc,
		createBrandFunc:                  createBrandFunc,
		updateBrandFunc:                  updateBrandFunc,
		getBranchesFilteredAsc:           getBranchesFilteredAsc,
		getBranchesFilteredDesc:          getBranchesFilteredDesc,
		getBranchesFilteredCount:         getBranchesFilteredCount,
		removeDisabledUsersFunc:          removeDisabledUsersFunc,
	}
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
