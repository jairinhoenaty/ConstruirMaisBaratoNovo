package app

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"

	pkgmiddleware "construir_mais_barato/infra/web/middleware"

	pkguser "construir_mais_barato/app/domain/user"
	pkguseruc "construir_mais_barato/app/usecase/user"
	pkguserinfra "construir_mais_barato/infra/database/repositories/user"

	pkgregion "construir_mais_barato/app/domain/region"
	pkgregionuc "construir_mais_barato/app/usecase/region"
	pkgregioninfra "construir_mais_barato/infra/database/repositories/region"

	pkgprofession "construir_mais_barato/app/domain/profession"
	pkgprofessionuc "construir_mais_barato/app/usecase/profession"
	pkgprofessioninfra "construir_mais_barato/infra/database/repositories/profession"

	pkgcity "construir_mais_barato/app/domain/city"
	pkgcityuc "construir_mais_barato/app/usecase/city"
	pkgcityinfra "construir_mais_barato/infra/database/repositories/city"

	pkgcontact "construir_mais_barato/app/domain/contact"
	pkgcontactuc "construir_mais_barato/app/usecase/contact"
	pkgcontactinfra "construir_mais_barato/infra/database/repositories/contact"

	pkgchat "construir_mais_barato/app/domain/chat"
	pkgchatuc "construir_mais_barato/app/usecase/chat"
	pkgchatinfra "construir_mais_barato/infra/database/repositories/chat"

	pkgprofessional "construir_mais_barato/app/domain/professional"
	pkgprofessionaluc "construir_mais_barato/app/usecase/professional"
	pkgprofessionalinfra "construir_mais_barato/infra/database/repositories/professional"

	pkgstore "construir_mais_barato/app/domain/store"
	pkgstoreuc "construir_mais_barato/app/usecase/store"
	pkgstoreinfra "construir_mais_barato/infra/database/repositories/store"

	pkgclient "construir_mais_barato/app/domain/client"
	pkgclientuc "construir_mais_barato/app/usecase/client"
	pkgclientinfra "construir_mais_barato/infra/database/repositories/client"

	pkgbudget "construir_mais_barato/app/domain/budget"
	pkgbudgetuc "construir_mais_barato/app/usecase/budget"
	pkgbudgetinfra "construir_mais_barato/infra/database/repositories/budget"

	pkgbanner "construir_mais_barato/app/domain/banner"
	pkgbanneruc "construir_mais_barato/app/usecase/banner"
	pkgbannerinfra "construir_mais_barato/infra/database/repositories/banner"

	pkgproduct "construir_mais_barato/app/domain/product"
	pkgproductuc "construir_mais_barato/app/usecase/product"
	pkgproductinfra "construir_mais_barato/infra/database/repositories/product"

	pkgproductCategory "construir_mais_barato/app/domain/productCategory"
	pkgproductCategoryuc "construir_mais_barato/app/usecase/productCategory"
	pkgproductCategoryinfra "construir_mais_barato/infra/database/repositories/productCategory"

	pkgauthenticateuc "construir_mais_barato/app/usecase/auth"

	pkgcontrollers "construir_mais_barato/infra/web/controllers"
)

type Server struct {
	server *http.Server
}

type dependenceParams struct {
	UserService            pkguser.UserService
	ProfessionService      pkgprofession.ProfessionService
	CityService            pkgcity.CityService
	ContactService         pkgcontact.ContactService
	ChatService         pkgchat.ChatService
	ProfessionalService    pkgprofessional.ProfessionalService
	BudgetService          pkgbudget.BudgetService
	BannerService          pkgbanner.BannerService
	ProductService         pkgproduct.ProductService
	ProductCategoryService pkgproductCategory.ProductCategoryService
	StoreService           pkgstore.StoreService
	ClientService          pkgclient.ClientService
	RegionService      pkgregion.RegionService

}

func buildDependenciesParams(db *gorm.DB) dependenceParams {
	params := dependenceParams{}

	params.UserService = pkguser.NewUserService(pkguserinfra.NewUserRepositoryImpl(db))
	params.ProfessionService = pkgprofession.NewProfessionService(pkgprofessioninfra.NewProfessionRepositoryImpl(db))
	params.CityService = pkgcity.NewCityService(pkgcityinfra.NewCityRepositoryImpl(db))
	params.ContactService = pkgcontact.NewContactService(pkgcontactinfra.NewContactRepositoryImpl(db))
	params.ChatService = pkgchat.NewChatService(pkgchatinfra.NewChatRepositoryImpl(db))
	params.ProfessionalService = pkgprofessional.NewProfessionalService(pkgprofessionalinfra.NewProfessionalRepositoryImpl(db))
	params.BudgetService = pkgbudget.NewBudgetService(pkgbudgetinfra.NewBudgetRepositoryImpl(db))
	params.BannerService = pkgbanner.NewBannerService(pkgbannerinfra.NewBannerRepositoryImpl(db))
	params.ProductService = pkgproduct.NewProductService(pkgproductinfra.NewProductRepositoryImpl(db))
	params.ProductCategoryService = pkgproductCategory.NewProductCategoryService(pkgproductCategoryinfra.NewProductCategoryRepositoryImpl(db))
	params.StoreService = pkgstore.NewStoreService(pkgstoreinfra.NewStoreRepositoryImpl(db))
	params.ClientService = pkgclient.NewClientService(pkgclientinfra.NewClientRepositoryImpl(db))
	params.RegionService = pkgregion.NewRegionService(pkgregioninfra.NewRegionRepositoryImpl(db))

	return params
}

func buildUserEndPoint(dependency *dependenceParams, g *echo.Group) {

	// parametros do caso de uso FindAll
	findAllParams := pkguseruc.FindAllUserUCParams{
		Service: dependency.UserService,
	}

	findByIdParams := pkguseruc.FindByIdUCParamns{
		Service: dependency.UserService,
	}

	findByEmailParams := pkguseruc.FindByEmailUCParams{
		Service: dependency.UserService,
	}

	saveParams := pkguseruc.SaveUserUCParams{
		Service: dependency.UserService,
	}

	deleteParams := pkguseruc.DeleteUserUCParams{
		Service: dependency.UserService,
	}

	// parametros do userController
	userControllerParams := pkgcontrollers.UserControllerParams{
		FindByIdUCParams:    findByIdParams,
		SaveUserUCParams:    saveParams,
		DeleteUserUCParams:  deleteParams,
		FindAllUserUCParams: findAllParams,
		FindByEmailUCParams: findByEmailParams,
	}

	pkgcontrollers.NewUserController(&userControllerParams, g)
}

func buildProfessionEndPoint(dependency *dependenceParams, g *echo.Group) {
	// parametros do caso de uso FindAll
	findAllParams := pkgprofessionuc.FindAllProfessionUCParams{
		Service: dependency.ProfessionService,
	}
	findByIdParams := pkgprofessionuc.FindByIdUCParams{
		Service: dependency.ProfessionService,
	}

	saveParams := pkgprofessionuc.SaveProfessionUCParams{
		Service: dependency.ProfessionService,
	}

	deleteParams := pkgprofessionuc.DeleteProfessionUCParams{
		Service: dependency.ProfessionService,
	}

	// parametros do userController
	professionControllerParams := pkgcontrollers.ProfessionControllerParams{
		FindAllProfessionUCParams: findAllParams,
		FindByIdUCParams:          findByIdParams,
		SaveProfessionUCParams:    saveParams,
		DeleteProfessionUCParams:  deleteParams,
	}

	pkgcontrollers.NewProfessionController(&professionControllerParams, g)
}


func buildRegionEndPoint(dependency *dependenceParams, g *echo.Group) {
	// parametros do caso de uso FindAll
	findAllParams := pkgregionuc.FindAllRegionUCParams{
		Service: dependency.RegionService,
	}
	findByIdParams := pkgregionuc.FindByIdUCParams{
		Service: dependency.RegionService,
	}

	saveParams := pkgregionuc.SaveRegionUCParams{
		Service: dependency.RegionService,
	}

	deleteParams := pkgregionuc.DeleteRegionUCParams{
		Service: dependency.RegionService,
	}

	// parametros do userController
	regionControllerParams := pkgcontrollers.RegionControllerParams{
		FindAllRegionUCParams: findAllParams,
		FindByIdUCParams:          findByIdParams,
		SaveRegionUCParams:    saveParams,
		DeleteRegionUCParams:  deleteParams,
	}

	pkgcontrollers.NewRegionController(&regionControllerParams, g)
}

func buildCityEndPoint(dependency *dependenceParams, g *echo.Group) {
	// parametros do caso de uso FindAll
	findAllParams := pkgcityuc.FindAllCityUCParams{
		Service: dependency.CityService,
	}
	findByIdParams := pkgcityuc.FindByIdUCParams{
		Service: dependency.CityService,
	}

	findByUFParams := pkgcityuc.FindByUFUCParams{
		Service: dependency.CityService,
	}

	saveParams := pkgcityuc.SaveCityUCParams{
		Service: dependency.CityService,
	}

	deleteParams := pkgcityuc.DeleteCityUCParams{
		Service: dependency.CityService,
	}

	// parametros do userController
	cityControllerParams := pkgcontrollers.CityControllerParams{
		FindAllCityUCParams: findAllParams,
		FindByIdUCParams:    findByIdParams,
		FindByUFUCParams:    findByUFParams,
		SaveCityUCParams:    saveParams,
		DeleteCityUCParams:  deleteParams,
	}

	pkgcontrollers.NewCityController(&cityControllerParams, g)
}

func buildContactEndPoint(dependency *dependenceParams, g *echo.Group) {
	// parametros do caso de uso FindAll
	findAllParams := pkgcontactuc.FindAllContactUCParams{
		Service: dependency.ContactService,
	}

	findByUserParams := pkgcontactuc.FindByUserContactUCParams{
		Service: dependency.ContactService,
	}

	findByIdParams := pkgcontactuc.FindByIdUCParams{
		Service: dependency.ContactService,
	}

	saveParams := pkgcontactuc.SaveContactUCParams{
		Service: dependency.ContactService,
	}

	deleteParams := pkgcontactuc.DeleteContactUCParams{
		Service: dependency.ContactService,
	}

	// parametros do userController
	contactControllerParams := pkgcontrollers.ContactControllerParams{
		FindAllContactUCParams:    findAllParams,
		FindByUserContactUCParams: findByUserParams,
		FindByIdUCParams:          findByIdParams,
		SaveContactUCParams:       saveParams,
		DeleteContactUCParams:     deleteParams,
	}

	pkgcontrollers.NewContactController(&contactControllerParams, g)
}


func buildChatEndPoint(dependency *dependenceParams, g *echo.Group) {
	// parametros do caso de uso FindAll
	findAllParams := pkgchatuc.FindAllChatUCParams{
		Service: dependency.ChatService,
	}

	findByUserParams := pkgchatuc.FindByUserChatUCParams{
		Service: dependency.ChatService,
	}

	/*findByIdParams := pkgchatuc.FindByIdUCParams{
		Service: dependency.ChatService,
	}*/

	saveParams := pkgchatuc.SaveChatUCParams{
		Service: dependency.ChatService,
	}

	deleteParams := pkgchatuc.DeleteChatUCParams{
		Service: dependency.ChatService,
	}

	// parametros do userController
	chatControllerParams := pkgcontrollers.ChatControllerParams{
		FindAllChatUCParams:    findAllParams,
		FindByUserChatUCParams: findByUserParams,
		//FindByIdUCParams:          findByIdParams,
		SaveChatUCParams:       saveParams,
		DeleteChatUCParams:     deleteParams,
	}

	pkgcontrollers.NewChatController(&chatControllerParams, g)
}


func buildProfessionalEndPoint(dependency *dependenceParams, g *echo.Group) {
	// parametros do caso de uso FindAll
	findAllParams := pkgprofessionaluc.FindAllProfessionalUCParams{
		Service: dependency.ProfessionalService,
	}
	findByIdParams := pkgprofessionaluc.FindByIdUCParamns{
		Service:     dependency.ProfessionalService,
		ServiceUser: dependency.UserService,
	}

	saveParams := pkgprofessionaluc.SaveProfessionalUCParams{
		Service:     dependency.ProfessionalService,
		ServiceUser: dependency.UserService,
	}

	deleteParams := pkgprofessionaluc.DeleteProfessionalUCParams{
		Service:     dependency.ProfessionalService,
		ServiceUser: dependency.UserService,
	}

	findLastProfessionalsUCParams := pkgprofessionaluc.FindLastProfessionalsUCParams{
		Service: dependency.ProfessionalService,
	}

	findByNamedUCParams := pkgprofessionaluc.FindByNamedUCParams{
		Service: dependency.ProfessionalService,
	}

	findByProfessionAndLocationUCParams := pkgprofessionaluc.FindByProfessionAndLocationUCParams{
		Service: dependency.ProfessionalService,
	}

	countProfessionalsByProfessionUCParams := pkgprofessionaluc.CountProfessionalsByProfessionUCParams{
		Service: dependency.ProfessionalService,
	}

	countCityProfessionalsByStateUCParams := pkgprofessionaluc.CountCityProfessionalsByStateUCParams{
		Service: dependency.ProfessionalService,
	}

	countProfessionalsByStateUCParams := pkgprofessionaluc.CountProfessionalsByStateUCParams{
		Service: dependency.ProfessionalService,
	}

	countProfessionalsByProfessionInCityUCParams := pkgprofessionaluc.CountProfessionalsByProfessionInCityUCParams{
		Service: dependency.ProfessionalService,
	}

	exportXLSXProfessionalUCParams := pkgprofessionaluc.ExportXLSXProfessionalUCParams{
		Service: dependency.ProfessionalService,
	}

	// parametros do userController
	professionalControllerParams := pkgcontrollers.ProfessionalControllerParams{
		FindAllProfessionalUCParams:                  findAllParams,
		FindByIdUCParams:                             findByIdParams,
		FindByNamedUCParams:                          findByNamedUCParams,
		FindByProfessionAndLocationUCParams:          findByProfessionAndLocationUCParams,
		SaveProfessionalUCParams:                     saveParams,
		DeleteProfessionalUCParams:                   deleteParams,
		ExportXLSXProfessionalUCParams:               exportXLSXProfessionalUCParams,
		FindLastProfessionalsUCParams:                findLastProfessionalsUCParams,
		CountProfessionalsByProfessionUCParams:       countProfessionalsByProfessionUCParams,
		CountProfessionalsByStateUCParams:            countProfessionalsByStateUCParams,
		CountProfessionalsByProfessionInCityUCParams: countProfessionalsByProfessionInCityUCParams,
		CountCityProfessionalsByStateUCParams:        countCityProfessionalsByStateUCParams,
	}

	pkgcontrollers.NewProfessionalController(&professionalControllerParams, g)
}

func buildStoreEndPoint(dependency *dependenceParams, g *echo.Group) {
	// parametros do caso de uso FindAll
	findAllParams := pkgstoreuc.FindAllStoreUCParams{
		Service: dependency.StoreService,
	}

	findByIdParams := pkgstoreuc.FindByIdUCParamns{
		Service:     dependency.StoreService,
		ServiceUser: dependency.UserService,
	}

	saveParams := pkgstoreuc.SaveStoreUCParams{
		Service:     dependency.StoreService,
		ServiceUser: dependency.UserService,
	}

	deleteParams := pkgstoreuc.DeleteStoreUCParams{
		Service:     dependency.StoreService,
		ServiceUser: dependency.UserService,
	}

	findLastStoresUCParams := pkgstoreuc.FindLastStoresUCParams{
		Service: dependency.StoreService,
	}

	findByNamedUCParams := pkgstoreuc.FindByNamedUCParams{
		Service: dependency.StoreService,
	}
	/*
		countProfessionalsByProfessionUCParams := pkgprofessionaluc.CountProfessionalsByProfessionUCParams{
			Service: dependency.ProfessionalService,
		}

		countProfessionalsByStateUCParams := pkgprofessionaluc.CountProfessionalsByStateUCParams{
			Service: dependency.ProfessionalService,
		}
		countProfessionalsByProfessionInCityUCParams := pkgprofessionaluc.CountProfessionalsByProfessionInCityUCParams{
			Service: dependency.ProfessionalService,
		}*/

	exportXLSXStoreUCParams := pkgstoreuc.ExportXLSXStoreUCParams{
		Service: dependency.StoreService,
	}

	// parametros do userController
	storeControllerParams := pkgcontrollers.StoreControllerParams{
		FindAllStoreUCParams:    findAllParams,
		FindByIdUCParams:        findByIdParams,
		FindByNamedUCParams:     findByNamedUCParams,
		SaveStoreUCParams:       saveParams,
		DeleteStoreUCParams:     deleteParams,
		ExportXLSXStoreUCParams: exportXLSXStoreUCParams,
		FindLastStoresUCParams:  findLastStoresUCParams,
		/*
			CountProfessionalsByProfessionUCParams:       countProfessionalsByProfessionUCParams,
			CountProfessionalsByStateUCParams:            countProfessionalsByStateUCParams,
			CountProfessionalsByProfessionInCityUCParams: countProfessionalsByProfessionInCityUCParams,*/
	}

	pkgcontrollers.NewStoreController(&storeControllerParams, g)
}

func buildClientEndPoint(dependency *dependenceParams, g *echo.Group) {
	// parametros do caso de uso FindAll
	findAllParams := pkgclientuc.FindAllClientUCParams{
		Service: dependency.ClientService,
	}

	findByIdParams := pkgclientuc.FindByIdUCParamns{
		Service:     dependency.ClientService,
		ServiceUser: dependency.UserService,
	}

	saveParams := pkgclientuc.SaveClientUCParams{
		Service:     dependency.ClientService,
		ServiceUser: dependency.UserService,
	}

	deleteParams := pkgclientuc.DeleteClientUCParams{
		Service:     dependency.ClientService,
		ServiceUser: dependency.UserService,
	}

	findLastClientsUCParams := pkgclientuc.FindLastClientsUCParams{
		Service: dependency.ClientService,
	}

	findByNamedUCParams := pkgclientuc.FindByNamedUCParams{
		Service: dependency.ClientService,
	}
	/*
		countProfessionalsByProfessionUCParams := pkgprofessionaluc.CountProfessionalsByProfessionUCParams{
			Service: dependency.ProfessionalService,
		}

		countProfessionalsByStateUCParams := pkgprofessionaluc.CountProfessionalsByStateUCParams{
			Service: dependency.ProfessionalService,
		}
		countProfessionalsByProfessionInCityUCParams := pkgprofessionaluc.CountProfessionalsByProfessionInCityUCParams{
			Service: dependency.ProfessionalService,
		}
	*/
	exportXLSXClientUCParams := pkgclientuc.ExportXLSXClientUCParams{
		Service: dependency.ClientService,
	}

	// parametros do userController
	clientControllerParams := pkgcontrollers.ClientControllerParams{
		FindAllClientUCParams:    findAllParams,
		FindByIdUCParams:         findByIdParams,
		FindByNamedUCParams:      findByNamedUCParams,
		SaveClientUCParams:       saveParams,
		DeleteClientUCParams:     deleteParams,
		ExportXLSXClientUCParams: exportXLSXClientUCParams,
		FindLastClientsUCParams:  findLastClientsUCParams,
		/*
			CountProfessionalsByProfessionUCParams:       countProfessionalsByProfessionUCParams,
			CountProfessionalsByStateUCParams:            countProfessionalsByStateUCParams,
			CountProfessionalsByProfessionInCityUCParams: countProfessionalsByProfessionInCityUCParams,*/
	}

	pkgcontrollers.NewClientController(&clientControllerParams, g)
}

func buildBudgetEndPoint(dependency *dependenceParams, g *echo.Group) {
	// parametros do caso de uso FindAll
	findAllParams := pkgbudgetuc.FindAllBudgetUCParams{
		Service: dependency.BudgetService,
	}
	findByIdParams := pkgbudgetuc.FindByIdUCParams{
		Service: dependency.BudgetService,
	}

	saveParams := pkgbudgetuc.SaveBudgetUCParams{
		Service: dependency.BudgetService,
	}

	deleteParams := pkgbudgetuc.DeleteBudgetUCParams{
		Service: dependency.BudgetService,
	}

	findByMonthAndProfessionalIDUCParams := pkgbudgetuc.FindByMonthAndProfessionalIDUCParams{
		Service:             dependency.BudgetService,
		ServiceUser:         dependency.UserService,
		ServiceProfessional: dependency.ProfessionalService,
	}

	FindByEmailUCParams := pkgbudgetuc.FindByEmailUCParams{
		Service: dependency.BudgetService,
	}
	// parametros do userController
	budgetControllerParams := pkgcontrollers.BudgetControllerParams{
		FindAllBudgetUCParams:                findAllParams,
		FindByIdUCParams:                     findByIdParams,
		SaveBudgetUCParams:                   saveParams,
		DeleteBudgetUCParams:                 deleteParams,
		FindByMonthAndProfessionalIDUCParams: findByMonthAndProfessionalIDUCParams,
		FindByEmailUCParams:                  FindByEmailUCParams,
	}

	pkgcontrollers.NewBudgetController(&budgetControllerParams, g)
}

func buildLoginEndPoint(dependency *dependenceParams, g *echo.Group) {

	authenticateParams := pkgauthenticateuc.AuthenticateUCParams{
		UserService:         dependency.UserService,
		ProfessionalService: dependency.ProfessionalService,
	}

	authenticateControllerParams := pkgcontrollers.AuthControllerParams{
		AuthenticateUCParams: authenticateParams,
	}

	pkgcontrollers.NewAuthController(authenticateControllerParams, g)
}

func buildPublicEndPoint(dependency *dependenceParams, g *echo.Group) {

	findRegionByCityUCParams := pkgregionuc.FindByCityUCParams{
		Service: dependency.RegionService,
	}

	FindByEmailUCParams := pkguseruc.FindByEmailUCParams{
		Service: dependency.UserService,
	}

	SaveClientUCParams := pkgclientuc.SaveClientUCParams{
		Service:     dependency.ClientService,
		ServiceUser: dependency.UserService,
	}

	SaveStoreUCParams := pkgstoreuc.SaveStoreUCParams{
		Service:     dependency.StoreService,
		ServiceUser: dependency.UserService,
	}

	FindByPageUCParams := pkgbanneruc.FindByPageUCParams{
		Service: dependency.BannerService,
	}

	SaveContactUCParams := pkgcontactuc.SaveContactUCParams{
		Service: dependency.ContactService,
	}

	findProfessionsParams := pkgprofessionuc.FindProfessionUCParams{
		Service: dependency.ProfessionService,
	}

	findProfessionsWithCountIdUCParams := pkgprofessionuc.FindProfessionsWithCountIdUCParams{
		Service: dependency.ProfessionService,
	}

	findByUFUCParams := pkgcityuc.FindByUFUCParams{
		Service: dependency.CityService,
	}

	findByProfessionalByCityAndProfessionUCParamns := pkgprofessionaluc.FindByProfessionalByCityAndProfessionUCParamns{
		Service: dependency.ProfessionalService,
	}

	findAllProfessionUCParams := pkgprofessionuc.FindAllProfessionUCParams{
		Service: dependency.ProfessionService,
	}

	saveProfessionalUCParams := pkgprofessionaluc.SaveProfessionalUCParams{
		Service:     dependency.ProfessionalService,
		ServiceUser: dependency.UserService,
	}

	saveBudgetUCParams := pkgbudgetuc.SaveBudgetUCParams{
		Service: dependency.BudgetService,
	}

	userSendEmailUCParams := pkguseruc.UserSendEmailUCParams{
		Service: dependency.UserService,
	}

	resetPasswordUCParams := pkguseruc.ResetPasswordUCParams{
		Service: dependency.UserService,
	}

	findByNameProfessionalAndCityAndProfessionUCParamns := pkgprofessionaluc.FindByNameProfessionalAndCityAndProfessionUCParamns{
		Service: dependency.ProfessionalService,
	}

	findByCityAndProfessionUCParams := pkgbanneruc.FindByCityAndProfessionUCParams{
		Service: dependency.BannerService,
	}

	findAllWithoutPaginationProfessionParams := pkgprofessionuc.FindAllWithoutPaginationProfessionParams{
		Service: dependency.ProfessionService,
	}

	findDayOfferProductsUCParams := pkgproductuc.FindDayofferProductUCParams{
		Service: dependency.ProductService,
	}

	findAllProductsUCParams := pkgproductuc.FindAllProductUCParams{
		Service: dependency.ProductService,
	}

	findByCityProductsUCParams := pkgproductuc.FindByCityUCParams{
		Service: dependency.ProductService,
	}

	// parametros do userController
	publicControllerParams := pkgcontrollers.PublicControllerParams{
		FindRegionByCityIdUCParams:  						 findRegionByCityUCParams,
		FindByEmailUCParams:                                 FindByEmailUCParams,
		SaveClientUCParams:                                  SaveClientUCParams,
		SaveStoreUCParams:                                   SaveStoreUCParams,
		FindByPageUCParams:                                  FindByPageUCParams,
		SaveContactUCParams:                                 SaveContactUCParams,
		FindByCityProductUCParams:                           findByCityProductsUCParams,
		FindAllProductUCParams:                              findAllProductsUCParams,
		FindDayofferProductUCParams:                         findDayOfferProductsUCParams,
		FindByUFUCParams:                                    findByUFUCParams,
		SaveBudgetUCParams:                                  saveBudgetUCParams,
		UserSendEmailUCParams:                               userSendEmailUCParams,
		ResetPasswordUCParams:                               resetPasswordUCParams,
		FindProfessionUCParams:                              findProfessionsParams,
		SaveProfessionalUCParams:                            saveProfessionalUCParams,
		FindAllProfessionUCParams:                           findAllProfessionUCParams,
		FindByCityAndProfessionUCParams:                     findByCityAndProfessionUCParams,
		FindProfessionsWithCountIdUCParams:                  findProfessionsWithCountIdUCParams,
		FindByProfessionalByCityAndProfessionUCParamns:      findByProfessionalByCityAndProfessionUCParamns,
		FindByNameProfessionalAndCityAndProfessionUCParamns: findByNameProfessionalAndCityAndProfessionUCParamns,
		FindAllWithoutPaginationProfessionParams:            findAllWithoutPaginationProfessionParams,
	}

	pkgcontrollers.NewPublicController(&publicControllerParams, g)
}

func buildBannerEndPoint(dependenceParams *dependenceParams, g *echo.Group) {

	deleteBannerUCParams := pkgbanneruc.DeleteBannerUCParams{
		Service: dependenceParams.BannerService,
	}

	saveBannerUCParams := pkgbanneruc.SaveBannerUCParams{
		Service: dependenceParams.BannerService,
	}
	findByPageUCParams := pkgbanneruc.FindByPageUCParams{
		Service: dependenceParams.BannerService,
	}

	bannerControllerParams := pkgcontrollers.BannerControllerParams{
		DeleteBannerUCParams: deleteBannerUCParams,
		SaveBannerUCParams:   saveBannerUCParams,
		FindByPageUCParams:   findByPageUCParams,
	}
	pkgcontrollers.NewBannerController(&bannerControllerParams, g)
}

func buildProductEndPoint(dependenceParams *dependenceParams, g *echo.Group) {

	saveProductUCParams := pkgproductuc.SaveProductUCParams{
		Service: dependenceParams.ProductService,
	}

	findByIdUCParams := pkgproductuc.FindByIdUCParams{
		Service: dependenceParams.ProductService,
	}

	findAllProductsUCParams := pkgproductuc.FindAllProductUCParams{
		Service: dependenceParams.ProductService,
	}

	deleteProductUCParams := pkgproductuc.DeleteProductUCParams{
		Service: dependenceParams.ProductService,
	}

	productControllerParams := pkgcontrollers.ProductControllerParams{
		SaveProductUCParams:    saveProductUCParams,
		FindByIdUCParams:       findByIdUCParams,
		FindAllProductUCParams: findAllProductsUCParams,
		DeleteProductUCParams:  deleteProductUCParams,
	}
	pkgcontrollers.NewProductController(&productControllerParams, g)
}

func buildProductCategoryEndPoint(dependenceParams *dependenceParams, g *echo.Group) {

	findByProfessionUCParams := pkgproductCategoryuc.FindByProfessionUCParams{
		Service: dependenceParams.ProductCategoryService,
	}

	saveProductCategoryUCParams := pkgproductCategoryuc.SaveProductCategoryUCParams{
		Service: dependenceParams.ProductCategoryService,
	}

	productCategoryControllerParams := pkgcontrollers.ProductCategoryControllerParams{
		FindByProfessionUCParams:    findByProfessionUCParams,
		SaveProductCategoryUCParams: saveProductCategoryUCParams,
	}
	pkgcontrollers.NewProductCategoryController(&productCategoryControllerParams, g)
}

func Start(db *gorm.DB) {

	dependency := buildDependenciesParams(db)

	router := echo.New()

	// ****************************************  Middleware CORS
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAuthorization},
	}))

	// **************************************** Rotas públicas
	publicRouter := router.Group("/publica")
	buildLoginEndPoint(&dependency, publicRouter)
	buildPublicEndPoint(&dependency, publicRouter)

	// **************************************** Rotas privadas
	routerGroup := router.Group("/api/v1")

	// middleware para todas as rotas do grupo
	routerGroup.Use(pkgmiddleware.VerifyAndValidateToken)

	buildUserEndPoint(&dependency, routerGroup)
	buildProfessionEndPoint(&dependency, routerGroup)
	buildCityEndPoint(&dependency, routerGroup)
	buildChatEndPoint(&dependency, routerGroup)
	buildContactEndPoint(&dependency, routerGroup)
	buildProfessionalEndPoint(&dependency, routerGroup)
	buildBudgetEndPoint(&dependency, routerGroup)
	buildBannerEndPoint(&dependency, routerGroup)
	buildProductEndPoint(&dependency, routerGroup)
	buildProductCategoryEndPoint(&dependency, routerGroup)
	buildStoreEndPoint(&dependency, routerGroup)
	buildClientEndPoint(&dependency, routerGroup)
	buildRegionEndPoint(&dependency, routerGroup)

	// Adicione a rotina de desativação de orçamentos
	go deactivateExpiredBudgetsRoutine(dependency.BudgetService)

	// server := newServer(os.Getenv("APP_PORT"), router)
	server := newServer("5000", router)
	server.ListenAndServe()

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	<-stopChan

}

func newServer(port string, handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:         ":" + port,
			Handler:      handler,
			ReadTimeout:  5 * 60 * time.Second,
			WriteTimeout: 5 * 60 * time.Second,
		},
	}
}

func (s *Server) ListenAndServe() {
	go func() {
		fmt.Println("Server runing in port 5000")
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("erro: %s", err)
		}
	}()

}

func deactivateExpiredBudgetsRoutine(budgetService pkgbudget.BudgetService) {
	for {
		// Espera 24 horas
		time.Sleep(24 * time.Hour)

		//chamar o caso de uso para pesquisar por orçamentos vencidos
		params := pkgbudgetuc.CheckExpiredBudgetUCParams{
			Service: budgetService,
		}
		uc := pkgbudgetuc.NewCheckExpiredBudgetUC(params)
		err := uc.Execute()
		if err != nil {
			fmt.Println("Erro ao desativar orçamentos com mais de 15 dias => ", err)
		}
	}
}
