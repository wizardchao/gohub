package main

import (
    "flag"
    "fmt"
    "gohub/bootstrap"
    btsConfig "gohub/config"
    "gohub/pkg/config"
    "gohub/pkg/logger"
    "gohub/pkg/captcha"

    "github.com/gin-gonic/gin"
)

func init() {
    // 加载 config 目录下的配置信息
    btsConfig.Initialize()
}

func main() {

    // 配置初始化，依赖命令行 --env 参数
    var env string
    flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
    flag.Parse()
    config.InitConfig(env)


    // 初始化 Logger
    bootstrap.SetupLogger()


    gin.SetMode(gin.ReleaseMode)
    // new 一个 Gin Engine 实例
    router := gin.New()

    // 初始化 DB
    bootstrap.SetupDB()

      // 初始化 Redis
    bootstrap.SetupRedis()

    // 初始化路由绑定
    bootstrap.SetupRoute(router)
    logger.Dump(captcha.NewCaptcha().VerifyCaptcha("t3yHrqkhbvd2vdZAOKm4", "077896"), "正确的答案")
    logger.Dump(captcha.NewCaptcha().VerifyCaptcha("czEnUuCODhVhZQc6BUcE", "078591"), "错误的答案")

    // 运行服务
    err := router.Run(":" + config.Get("app.port"))
    if err != nil {
        // 错误处理，端口被占用了或者其他错误
        fmt.Println(err.Error())
    }
}