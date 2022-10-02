package main

import (
	"github.com/gin-gonic/gin"
	"shoexsmass/helper"
	"shoexsmass/model/category"
	"shoexsmass/model/notification"
	"shoexsmass/model/product"
	"shoexsmass/model/topup"
	"shoexsmass/model/transaction"
	"shoexsmass/model/transactiondetail"
	"shoexsmass/model/user"
)

func main() {
	db, err := helper.DbConnect()
	helper.ErrorHelper(err)
	userRepository := user.NewRepository(db)

	categoryRepository := category.NewRepository(db)
	categoryService := category.NewService(categoryRepository)
	categoryController := category.NewCategoryController(categoryService)

	notificationRepository := notification.NewRepository(db)
	notificationService := notification.NewService(notificationRepository)
	notificationController := notification.NewNotificationController(notificationService)

	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productController := product.NewProductController(productService)

	topUpRepository := topup.NewRepository(db)
	topUpService := topup.NewService(topUpRepository, userRepository)
	topUpController := topup.NewTopUpController(topUpService)

	transactionDetailRepository := transactiondetail.NewRepository(db)
	transactionRepository := transaction.NewRepository(db, transactionDetailRepository, userRepository)
	transactionService := transaction.NewService(transactionRepository)
	transactionController := transaction.NewTransactionController(transactionService)

	userService := user.NewService(userRepository)
	userController := user.NewUserController(userService)

	router := gin.Default()
	router.Use(CORSMiddleware())

	api := router.Group("/api/")

	router.Static("/images", "./images")

	api.GET("/category", categoryController.GetALl)

	api.POST("/notif", notificationController.CreateNotif)
	api.GET("/notif/:id", notificationController.GetByID)
	api.DELETE("/notif/:id", notificationController.Delete)
	api.GET("/notif/user/:id", notificationController.GetByUserID)

	api.GET("/products", productController.GetAll)
	api.GET("/products/cat/:id", productController.GetByCategoryID)
	api.GET("/products/:id", productController.GetByID)

	api.POST("/topup", topUpController.TopUp)
	api.GET("/topup/:id", topUpController.GetByID)
	api.GET("/topup/user/:id", topUpController.GetByUserID)

	api.POST("/transaction", transactionController.CreateTransaction)
	api.GET("/transaction/:id", transactionController.GetByID)
	api.GET("/transaction/user/:id", transactionController.GetByUserID)

	api.POST("/user", userController.RegisterUser)
	api.POST("/user/login", userController.Login)
	api.POST("/user/check", userController.CheckEmailAvailability)
	api.POST("/user/avatar/:id", userController.UploadAvatar)
	api.GET("/user/:id", userController.FetchUser)
	api.POST("/user/update", userController.UpdateUser)

	router.Run(":8000")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}