package app

import (
	_ "github.com/go-redis/redis"
	_ "github.com/hetianyi/base64Captcha"
	"github.com/hetianyi/easygo/base"
	"github.com/hetianyi/easygo/file"
	_ "gopkg.in/yaml.v2"
	_ "gorm.io/driver/mysql"
	_ "gorm.io/gorm"
	"io"
)

// Server 为http服务器的配置
type Server struct {
	Host         string     `yaml:"host"`         // HTTP服务监听地址
	Port         int        `yaml:"port"`         // HTTP服务监听端口
	GinMode      string     `yaml:"mode"`         // 服务启动模式：debug|release
	UseGinLogger bool       `yaml:"useGinLogger"` // 是否开启gin的默认Logger
	MiddleWares  []string   `yaml:"middleWares"`  // 中间件名称列表
	Validators   []string   `yaml:"validators"`   // 校验器名称列表
	ApiGroup     []ApiGroup `yaml:"groups"`       // 路由配置
}

type ApiGroup struct {
	Prefix      string     `yaml:"prefix"`      // API组的路由前缀
	MiddleWares []string   `yaml:"middleWares"` // API组的中间件名称列表
	ApiGroup    []ApiGroup `yaml:"groups"`      // 路由配置
	Apis        []Api      `yaml:"apis"`        // API组的API列表
}

type Api struct {
	// 路由，格式为: <Method> <Route>
	//
	// 例如：
	//  GET /api/user
	//  post /api/user/new
	Route    string   `yaml:"route"`
	Handlers []string `yaml:"handlers"` // API处理器名称列表
}

// RedisConfig 为redis的配置
//
// requires github.com/go-redis/redis
type RedisConfig struct {
	Host      string `yaml:"host"` // 服务器地址
	Port      int    `yaml:"port"` // 服务端口
	Pass      string `yaml:"pass"` // 密码
	DefaultDB int    `yaml:"db"`   // 默认数据库
}

// MysqlConfig 为mysql的配置
//
// requires gorm.io/driver/mysql and gorm.io/gorm
type MysqlConfig struct {
	// mysql连接，格式:
	//
	//  <username>:<password>@tcp(<host>:<port>)/<db>?charset=utf8&parseTime=True&loc=Local
	Conn string `yaml:"conn"`
	// 最大空闲连接
	MaxIdleConns int `yaml:"maxIdleConns"`
	// 最大连接数
	MaxOpenConns int `yaml:"maxOpenConns"`
	// 每个连接最长声明周期
	ConnMaxLifetime int `yaml:"connMaxLifetime"`
}

// CaptchaConfig requires github.com/hetianyi/base64Captcha
type CaptchaConfig struct {
	Width           int    `yaml:"width"`      // 图片宽度
	Height          int    `yaml:"height"`     // 图片高度
	Source          string `yaml:"seed"`       // 种子字符
	Length          int    `yaml:"length"`     // 验证码字符数
	NoiseCount      int    `yaml:"noiseCount"` // 干扰字符数
	FontName        string `yaml:"fontName"`   // 字体名称
	BackgroundColor string `yaml:"background"` // 背景色
	Store           string `yaml:"store"`      // 验证码保存：memory|redis
}

// Config 为全部配置
// requires gopkg.in/yaml.v2
type Config struct {
	Server        Server        `yaml:"server"`
	RedisConfig   RedisConfig   `yaml:"redis"`
	MysqlConfig   MysqlConfig   `yaml:"mysql"`
	CaptchaConfig CaptchaConfig `yaml:"captcha"`
}

type App struct {
	Config *Config
}

// NewApp 初始化一个App
func NewApp() *App {
	return &App{
		Config: &Config{},
	}
}

// LoadFromYamlFile 从指定路径的yaml文件加载配置
func (a *App) LoadFromYamlFile(path string) error {
	configFile, err := file.GetFile(path)
	if err != nil {
		//logger.Fatal("无法加载配置文件\"", path, "\": ", err)
		return err
	}
	content, err := io.ReadAll(configFile)
	if err != nil {
		//logger.Fatal("无法读取配置文件\"", path, "\": ", err)
		return err
	}
	a.Config = &Config{}
	err = base.ParseYamlFromString(string(content), a.Config)
	if err != nil {
		//logger.Fatal("配置文件错误: ", err)
		return err
	}
	return nil
}

// LoadFromYamlContent 从给定的yaml内容加载配置
func (a *App) LoadFromYamlContent(yaml string) error {
	a.Config = &Config{}
	err := base.ParseYamlFromString(yaml, a.Config)
	if err != nil {
		return err
	}
	return nil
}
