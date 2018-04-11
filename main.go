package main

import (
	"flag"
	"github.com/dingoblog/dingo/app"

	//"github.com/gorilla/mux"
)

func main() {
	//portPtr := flag.String("port", "8000", "The port number for Dingo to listen to.")
	//dbFilePathPtr := flag.String("database", "dingo.db", "The database file path for Djingo to use.")
	privKeyPathPtr := flag.String("priv-key", "dingo.rsa", "The private key file path for JWT.")
	pubKeyPathPtr := flag.String("pub-key", "dingo.rsa.pub", "The public key file path for JWT.")
	flag.Parse()

	Dingo.Init(*privKeyPathPtr, *pubKeyPathPtr)
	Dingo.Run()

	//router := gin.New()
	//router.Delims("{[{", "}]}")
	//gin.SetMode(gin.ReleaseMode)
	//mux.Router{}
	//router.Use(gin.Logger())
	//router.Use(gin.Recovery())
	//fmt.Println(os.Getwd())
	//
	//router.Static("/static", "/view/admin/*")
	//
	//
	//router.LoadHTMLGlob("./view/**/**/**/**/**/*")
	//
	//handler.InitializeGin(router)
	//router.Run(config.Conf.RunPort)
}
