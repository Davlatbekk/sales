package api

import (
	_ "app/api/docs"

	"app/api/handler"
	"app/config"
	"app/pkg/logger"
	"app/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func NewApi(r *gin.Engine, cfg *config.Config, store storage.StorageI, logger logger.LoggerI) {
	handler := handler.NewHandler(cfg, store, logger)
	// category api
	r.POST("/category", handler.CreateCategory)
	r.GET("/category/:id", handler.GetByIdCategory)
	r.GET("/category", handler.GetListCategory)
	r.PUT("/category/:id", handler.UpdateCategory)
	r.DELETE("/category/:id", handler.DeleteCategory)

	// product api
	r.POST("/product", handler.CreateProduct)
	r.GET("/product/:id", handler.GetByIdProduct)
	r.GET("/product", handler.GetListProduct)
	r.PUT("/product/:id", handler.UpdateProduct)
	r.DELETE("/product/:id", handler.DeleteProduct)



	// brands api
	r.POST("/brand", handler.Create)
	r.GET("/brand/:id", handler.GetById)
	r.GET("/brand", handler.GetList)
	r.PUT("/brand/:id", handler.Update)
	r.DELETE("/brand/:id", handler.Delete)


	// stocks api
	r.POST("/stock", handler.CreateStock)
	r.GET("/stock/:id", handler.GetByIDStock)
	r.GET("/stock", handler.GetListStock)
	r.PUT("/stock/:id", handler.UpdateStock)
	r.DELETE("/stock/:id", handler.DeleteStock)

	// order api
	// r.POST("/order", handler.CreateOrder)
	// r.GET("/order/:id", handler.GetByIdOrder)
	// r.GET("/order", handler.GetListOrder)
	// r.PUT("/order/:id", handler.UpdateOrder)
	// r.PATCH("/order/:id", handler.UpdatePatchOrder)
	// r.DELETE("/order/:id", handler.DeleteOrder)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
