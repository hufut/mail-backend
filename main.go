package main

import (
	"fmt"
	"net/http"

	"github.com/hukaixuan/mall-backend/pkg/setting"
	"github.com/hukaixuan/mall-backend/routers"
)

// @title api for mall app
// @version 1.0
// @description This is a server for mall app.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email hukx.michael@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host fujingwen.com
// @BasePath /api/v1
func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
