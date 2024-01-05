package service

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/cuiyuanxin/airuisi-admin/global"
	"github.com/cuiyuanxin/airuisi-admin/pkg/setting"
	"github.com/spf13/cobra"
	"github.com/zztroot/color"
)

var (
	envConf string
	Cmd     = &cobra.Command{
		Use:     "service",
		Short:   "Starting the service",
		Long:    "The default config configuration can be overridden by specifying the config file path, which is optional",
		Example: global.AppName + " service -c config/config.yaml",
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	envConf = os.Getenv("APP_CONF")
	if envConf == "" {
		Cmd.PersistentFlags().StringVarP(&envConf, "config", "c", "config/prod.yaml", "config path, eg: -conf config/prod.yml")
	}
	if envConf == "" {
		envConf = "config/prod.yaml"
	}
	log.Println("Load the conf file:", envConf)
}

func setSttings(settings *setting.Setting) {
	if err := settings.ReadSection("Server", &global.ServerSetting); err != nil {
		log.Fatalf("init.setupLogger err：%s", color.Coat(err.Error(), color.Red))
	}
	//if err := settings.ReadSection("Database", &global.DatabaseSetting);err != nil {
	//	log.Fatalf("init.setupLogger err：%s", color.Coat(err.Error(), color.Red))
	//}
	//if err := settings.ReadSection("Logger", &global.LoggerSetting);err != nil {
	//	log.Fatalf("init.setupLogger err：%s", color.Coat(err.Error(), color.Red))
	//}
	//if err := settings.ReadSection("App", &global.AppSetting);err != nil {
	//	log.Fatalf("init.setupLogger err：%s", color.Coat(err.Error(), color.Red))
	//}
	//if err := settings.ReadSection("JWT", &global.JwtSetting);err != nil {
	//	log.Fatalf("init.setupLogger err：%s", color.Coat(err.Error(), color.Red))
	//}
	//if err := settings.ReadSection("Email", &global.EmailSetting);err != nil {
	//	log.Fatalf("init.setupLogger err：%s", color.Coat(err.Error(), color.Red))
	//}
	//if err := settings.ReadSection("Tracer", &global.TracerSetting);err != nil {
	//	log.Fatalf("init.setupLogger err：%s", color.Coat(err.Error(), color.Red))
	//}
	//if err := settings.ReadSection("Sms", &global.SmsSetting);err != nil {
	//	log.Fatalf("init.setupLogger err：%s", color.Coat(err.Error(), color.Red))
	//}
}

func setup() {
	// 初始化配置
	settings, err := setting.NewSetting(strings.Split(envConf, ".")...)
	if err != nil {
		log.Fatalf("init.setupSetting err：%s", color.Coat(err.Error(), color.Red))
	}
	setSttings(settings)

	fmt.Println(settings)
	log.Println("Configuration initialization successfully...")

	// 初始化日志
	//if err := logger.NewLogger(global.LoggerSetting); err != nil {
	//	log.Fatalf("init.setupLogger err：%s", color.Coat(err.Error(), color.Red))
	//}
	//log.Println("Log initialization successfully...")

	// 初始化链路追踪
	//if err := tracer.NewJaegerTracer(global.AppName+` service`, global.TracerSetting.AgentHostPort); err != nil {
	//	log.Fatalf("init.setupTracet err：%s", color.Coat(err.Error(), color.Red))
	//}
	//log.Println("Link tracking was initialized successfully...")

	// 初始化数据库
	//if err := model.NewDBEngine(global.DatabaseSetting); err != nil {
	//	log.Fatalf("init.setupDBEngine err：%s", color.Coat(err.Error(), color.Red))
	//}
	//log.Println("Database initialization complete...")

}

func run() error {
	//if global.ServerSetting.GinMode != "" {
	//	gin.SetMode(global.ServerSetting.GinMode)
	//}
	//	initRouter()
	//
	//	for _, f := range AppRouters {
	//		f()
	//	}
	//
	//	srv := &http.Server{
	//		Addr:    fmt.Sprintf("%s:%d", config.ApplicationConfig.Host, config.ApplicationConfig.Port),
	//		Handler: sdk.Runtime.GetEngine(),
	//	}
	//
	//	go func() {
	//		jobs.InitJob()
	//		jobs.Setup(sdk.Runtime.GetDb())
	//
	//	}()
	//
	//	if apiCheck {
	//		var routers = sdk.Runtime.GetRouter()
	//		q := sdk.Runtime.GetMemoryQueue("")
	//		mp := make(map[string]interface{})
	//		mp["List"] = routers
	//		message, err := sdk.Runtime.GetStreamMessage("", global.ApiCheck, mp)
	//		if err != nil {
	//			log.Printf("GetStreamMessage error, %s \n", err.Error())
	//			//日志报错错误，不中断请求
	//		} else {
	//			err = q.Append(message)
	//			if err != nil {
	//				log.Printf("Append message error, %s \n", err.Error())
	//			}
	//		}
	//	}
	//
	//	go func() {
	//		// 服务连接
	//		if config.SslConfig.Enable {
	//			if err := srv.ListenAndServeTLS(config.SslConfig.Pem, config.SslConfig.KeyStr); err != nil && !errors.Is(err, http.ErrServerClosed) {
	//				log.Fatal("listen: ", err)
	//			}
	//		} else {
	//			if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
	//				log.Fatal("listen: ", err)
	//			}
	//		}
	//	}()
	//	fmt.Println(pkg.Red(string(global.LogoContent)))
	//	tip()
	//	fmt.Println(pkg.Green("Server run at:"))
	//	fmt.Printf("-  Local:   %s://localhost:%d/ \r\n", "http", config.ApplicationConfig.Port)
	//	fmt.Printf("-  Network: %s://%s:%d/ \r\n", "http", pkg.GetLocaHonst(), config.ApplicationConfig.Port)
	//	fmt.Println(pkg.Green("Swagger run at:"))
	//	fmt.Printf("-  Local:   http://localhost:%d/swagger/admin/index.html \r\n", config.ApplicationConfig.Port)
	//	fmt.Printf("-  Network: %s://%s:%d/swagger/admin/index.html \r\n", "http", pkg.GetLocaHonst(), config.ApplicationConfig.Port)
	//	fmt.Printf("%s Enter Control + C Shutdown Server \r\n", pkg.GetCurrentTimeStr())
	//	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	//	quit := make(chan os.Signal, 1)
	//	signal.Notify(quit, os.Interrupt)
	//	<-quit
	//
	//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//	defer cancel()
	//	fmt.Printf("%s Shutdown Server ... \r\n", pkg.GetCurrentTimeStr())
	//
	//	if err := srv.Shutdown(ctx); err != nil {
	//		log.Fatal("Server Shutdown:", err)
	//	}
	//	log.Println("Server exiting")
	//
	return nil
}

//
////var Router runtime.Router
//
//func tip() {
//	usageStr := `欢迎使用 ` + pkg.Green(`go-admin `+global.Version) + ` 可以使用 ` + pkg.Red(`-h`) + ` 查看命令`
//	fmt.Printf("%s \n\n", usageStr)
//}
//
//func initRouter() {
//	var r *gin.Engine
//	h := sdk.Runtime.GetEngine()
//	if h == nil {
//		h = gin.New()
//		sdk.Runtime.SetEngine(h)
//	}
//	switch h.(type) {
//	case *gin.Engine:
//		r = h.(*gin.Engine)
//	default:
//		log.Fatal("not support other engine")
//		//os.Exit(-1)
//	}
//	if config.SslConfig.Enable {
//		r.Use(handler.TlsHandler())
//	}
//	//r.Use(middleware.Metrics())
//	r.Use(common.Sentinel()).
//		Use(common.RequestId(pkg.TrafficKey)).
//		Use(api.SetRequestLogger)
//
//	common.InitMiddleware(r)
//
//}
