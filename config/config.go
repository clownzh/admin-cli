package config

type HttpConfig struct {
	Port      int  `json:"port"`
	OpenRedis bool `json:"openRedis"` //是否开启redis
}

type Mysql struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	DataBase string `json:"database"`
	LogoMode bool   `json:"logoMode"`
}

type Redis struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type Jwt struct {
	Key        string `json:"key"`
	ExpireTime int64  `json:"expireTime"`
}

type Log struct {
	Path       string `json:"path"`
	Maxsize    int    `json:"maxsize"`
	MaxBackups int    `json:"maxBackups"`
	MaxAge     int    `json:"maxAge"`
	Compress   bool   `json:"compress"`
}

type Config struct {
	HttpConfig HttpConfig `json:"httpConfig"`
	Redis      Redis      `json:"redis"`
	Mysql      Mysql      `json:"mysql"`
	Jwt        Jwt        `json:"jwt"`
	Log        Log        `json:"log"`
	OpenCache  bool       `json:"openCache"`
	Salt       string     `json:"salt"` //加密盐
}

var config Config

func SetConfig(cfg *Config) {
	config = *cfg
}

func GetConfig() *Config {
	return &config
}