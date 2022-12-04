package main

import (
	"dbo_backend_test/app/adapter/client"
	"dbo_backend_test/app/controller"
	"dbo_backend_test/app/infrastructure/connection"
	"dbo_backend_test/app/infrastructure/middleware"
	"dbo_backend_test/app/infrastructure/router"
	"dbo_backend_test/app/usecase/authentication"
	"dbo_backend_test/app/usecase/customer_uc"
	"dbo_backend_test/app/usecase/order_uc"
	"dbo_backend_test/app/usecase/product_uc"
	_ "dbo_backend_test/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title DBO BackEnd Test API
// @version 1.0
// @description This is Swagger Documentation API For BackEnd DBO Test.

// @contact.name Afdhalul Ihsan
// @contact.email afdhalulhsn74@gmail.com

// @license.name Apache 2.0

// @host localhost:8830
// @BasePath /
// @schemes http
func main() {

	route := gin.Default()
	url := ginSwagger.URL("http://localhost:8830/swagger/doc.json") // The url pointing to API definition
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	defautl := route.Group("/")
	//newAuthServie
	userCl := client.NewUserClient(connection.DbConnection)
	autCl := client.NewAuthClient(connection.DbConnection)
	authUc := authentication.NewAuthenticationUC(userCl, autCl)
	authSvc := controller.NewAuthenticationService(authUc)

	customerClient := client.NewCustomerClient(connection.DbConnection)
	customeruc := customer_uc.NewCustomerUC(customerClient)
	custSvc := controller.NewCustomerService(customeruc)

	orderClinet := client.NewOrderClient(connection.DbConnection)
	orderUc := order_uc.NewOrderUc(orderClinet)
	oderderSvc := controller.NewOrderController(orderUc)

	pdCli := client.NewProductClient(connection.DbConnection)
	pdUc := product_uc.NewProductUC(pdCli)
	pdSvc := controller.NewProductUc(pdUc)

	//URL Customer
	api := defautl.Group("/api/v1")
	api.Use(middleware.JwtAuthMiddleware())
	customerRoute := api.Group("/customer")
	router.Customer(customerRoute, custSvc)

	//urlOrder
	order := api.Group("/order")
	router.OrderRouter(order, oderderSvc)

	//RegisterRouter
	router.AuthenticationTouter(defautl, authSvc)

	//AdminRouter
	admin := route.Group("/admin")
	admin.Use(middleware.JwtAuthMiddlewareAdmin())
	router.AdminRouter(admin, authSvc)
	customerAdmin := admin.Group("/customer-admin")
	producAdmin := admin.Group("/produk")
	router.ProductRouter(producAdmin, pdSvc)
	router.CustomerAdmin(customerAdmin, custSvc)

	route.Run("localhost:8830")
}
