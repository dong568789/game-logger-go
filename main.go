package main

import (
	"fmt"
	"game-logger-go/router"
	"game-logger-go/utils"
	"game-logger-go/utils/logSave"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath("./")
	viper.SetConfigName("conf")
	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	//初始化日志
	utils.NewLogMgr()
	//初始化埋点日志
	logSave.InitCyLog(viper.GetString("cyLogDir"))
}

func main() {
	r := gin.New()

	if viper.GetBool("debug") == true {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.DefaultWriter = utils.GinLoggerMgr
		gin.SetMode(gin.ReleaseMode)
	}
	router.InitRouter(r)
	r.Run(":" + viper.GetString("httpPort")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
