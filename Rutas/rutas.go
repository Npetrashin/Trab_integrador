package rutas

import (
	"github.com/Npetrashin/Trab_integrador/Controladores"
	
		"github.com/gin-gonic/gin"
	)

	func UserRoutes(incomingRoutes *gin.Engine) {
		incomingRoutes.POST("/users/signup", Controladores.SignUp())
		incomingRoutes.POST("/users/login", Controladores.Login())
		incomingRoutes.POST("/admin/addproduct", Controladores.ProductViewerAdmin())
		incomingRoutes.GET("/users/productview", Controladores.SearchProduct())
		incomingRoutes.GET("/users/search", Controladores.SearchProductByQuery())
	}