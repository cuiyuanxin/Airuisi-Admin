package server

import (
	"github.com/cuiyuanxin/airuisi-admin/global"
	"github.com/cuiyuanxin/airuisi-admin/pkg/logger"
	"github.com/cuiyuanxin/airuisi-admin/pkg/setting"
)

func setupSetting(configs string) error {
	settings, err := setting.NewSetting(configs)
	if err != nil {
		return err
	}
	if err = settings.ReadSection("Server", &global.ServerSetting); err != nil {
		return err
	}
	if err = settings.ReadSection("Logger", &global.LoggerSetting); err != nil {
		return err
	}
	//err = settings.ReadSection("Database", &global.DatabaseSetting)
	//if err != nil {
	//	return err
	//}

	//err = settings.ReadSection("App", &global.AppSetting)
	//if err != nil {
	//	return err
	//}
	////err = settings.ReadSection("JWT", &global.JwtSetting)
	////if err != nil {
	////	return err
	////}
	//err = settings.ReadSection("Email", &global.EmailSetting)
	//if err != nil {
	//	return err
	//}
	//err = settings.ReadSection("Tracer", &global.TracerSetting)
	//if err != nil {
	//	return err
	//}
	//err = settings.ReadSection("Sms", &global.SmsSetting)
	//if err != nil {
	//	return err
	//}
	return nil
}

func setupLogger() error {
	var err error
	global.Logger, err = logger.NewLogger(global.LoggerSetting)
	if err != nil {
		return err
	}
	return nil
}

//func setupDBEngine() error {
//	var err error
//	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
//	if err != nil {
//		return err
//	}
//
//	if err = global.DBEngine.Ping(); err != nil {
//		return err
//	}
//
//	// 延迟关闭数据库
//	//defer global.DBEngine.Close()
//
//	return nil
//}
//

//
//func setupTracer() error {
//	jaegerTracer, _, err := tracer.NewJaegerTracer("beifang-server", global.TracerSetting.AgentHostPort)
//	if err != nil {
//		return err
//	}
//
//	global.Tracet = jaegerTracer
//	return nil
//}
//
//func setupTranslations() error {
//	err := validators.NewTranslator()
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func setupCasbin() error {
//	var err error
//	global.Casbin, err = casbin.NewCasbin(global.DatabaseSetting)
//	if err != nil {
//		return err
//	}
//	return nil
//}
