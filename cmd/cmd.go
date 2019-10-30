package cmd

import (
	"BusinessSys/config"
	"BusinessSys/ginengine"
	"BusinessSys/middleware"
	routes "BusinessSys/routers"
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = cobra.Command{
	Use: "BusinessSys",
	Run: func(cmd *cobra.Command, args []string) {
		// command := exec.Command("BusinessSys")

		// command.Start()

		// fmt.Printf("BusinessSys start, [pid] %d running...\n", command.Process.Pid)

		// ioutil.WriteFile("BusinessSys.lock", []byte(fmt.Sprintf("%d", command.Process.Pid)), 0666)

		// os.Exit(0)

		//初始化配置
		config.InitConfig()

		//初始化GinEngine
		ginengine.InitEngine()

		//初始化中间件
		middleware.SetUp()

		//初始化routes
		routes.Setup()

		//初始化数据库
		// db.InitDB()

		//gin run
		err := ginengine.Engine.Run(fmt.Sprintf(":%d", config.AppConfig.Port))

		if err != nil {
			panic("start gin failed!")
		}
	},
}
