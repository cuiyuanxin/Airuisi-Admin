package setting

// 服务地址和端口
type Server struct {
	Mode string `mapstructure:"Mode"`
	//HttpAddr     string        `mapstructure:"HttpAddr"`
	//HttpPort     string        `mapstructure:"HttpPort"`
	//ReadTimeout  time.Duration `mapstructure:"ReadTimeout"`
	//WriteTimeout time.Duration `mapstructure:"WriteTimeout"`
}

// 日志配置
type Logger struct {
	//Open       bool   `mapstructure:"Open"`
	//FilePath   string `mapstructure:"FilePath"`
	//MaxSize    int    `mapstructure:"MaxSize"`
	//MaxBackups int    `mapstructure:"MaxBackups"`
	//MaxAge     int    `mapstructure:"MaxAge"`
	//Compress   bool   `mapstructure:"Compress"`
}

// Tracer
type Tracer struct {
	//Open          bool   `mapstructure:"Open"`
	//AgentHostPort string `mapstructure:"AgentHostPort"`
}

// 数据库
type Database struct {
	//DBType                 string `mapstructure:"DBType"`
	//UserName               string `mapstructure:"UserName"`
	//Password               string `mapstructure:"Password"`
	//Host                   string `mapstructure:"Host"`
	//Port                   string `mapstructure:"Port"`
	//DBName                 string `mapstructure:"DBName"`
	//Prefix                 string `mapstructure:"Prefix"`
	//Charset                string `mapstructure:"Charset"`
	//MaxIdleConns           int    `mapstructure:"MaxIdleConns"`
	//MaxOpenConns           int    `mapstructure:"MaxOpenConns"`
	//DbConnMaxLifetimeHours int64  `mapstructure:"DbConnMaxLifetimeHours"`
}

// redis
//type Rdb struct {
//	Addr     string `mapstructure:"Addr"`
//	Port     int    `mapstructure:"Port"`
//	Password string `mapstructure:"Password"`
//	Db       int    `mapstructure:"Db"`
//}

// App通用配置
//type App struct {
//	Limit                 int           `mapstructure:"Limit"`
//	IsDel                 bool          `mapstructure:"IsDel"`
//	UploadSavePath        string        `mapstructure:"UploadSavePath"`
//	UploadServerUrl       string        `mapstructure:"UploadServerUrl"`
//	UploadImageMaxSize    int           `mapstructure:"UploadImageMaxSize"`
//	UploadImageAllowExts  []string      `mapstructure:"UploadImageAllowExts"`
//	UploadFileMaxSize     int           `mapstructure:"UploadFileMaxSize"`
//	UploadFileAllowExts   []string      `mapstructure:"UploadFileAllowExts"`
//	DefaultContextTimeout time.Duration `mapstructure:"DefaultContextTimeout"`
//	CryptKey              string        `mapstructure:"CryptKey"`
//	InfoValidTime         int           `mapstructure:"InfoValidTime"`
//	QrCryptLength         int           `mapstructure:"QrCryptLength"`
//}

// jwt
//type JWT struct {
//	AppKey    string        `mapstructure:"AppKey"`
//	Secret    string        `mapstructure:"Secret"`
//	Issuer    string        `mapstructure:"Issuer"`
//	ExpiresAt time.Duration `mapstructure:"ExpiresAt"`
//}

// Email
//type Email struct {
//	Host     string   `mapstructure:"Host"`
//	Port     int      `mapstructure:"Port"`
//	UserName string   `mapstructure:"UserName"`
//	Password string   `mapstructure:"Password"`
//	IsSSL    bool     `mapstructure:"IsSSL"`
//	From     string   `mapstructure:"From"`
//	To       []string `mapstructure:"To"`
//}

// Permission
//type Permission struct {
//	SecretKey string `mapstructure:"SecretKey"`
//	UserId    int64  `mapstructure:"UserId"`
//}
