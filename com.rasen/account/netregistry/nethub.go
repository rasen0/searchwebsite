package netregistry

import (
	"net/http"
	"time"

	"com.rasen/account/controller"
	"github.com/gin-gonic/gin"
)

func RegistryHub(){
	r := gin.Default()
	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())
	//gin.SetMode(gin.TestMode)
	server := &http.Server{
		Addr: ":5535",
		Handler:r,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 5 * time.Second,
	}
	// 添加http 接口
	registryApi(r)

	server.ListenAndServe()
}

func registryApi(r *gin.Engine){
	r.Static("/","static")
	r.POST("account/v1/search",controller.SearchHandler)
}

