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
	utils.InitLogger(viper.GetString("logDir"))
	logSave.InitCyLog(viper.GetString("cyLogDir"))
}

func main() {
	r := gin.New()
	//gin.DefaultWriter = utils.Logger

	if viper.GetBool("debug") == true {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router.InitRouter(r)
	r.Run(":8057") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
