package common

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

var (
	WorkSpace     string       // config
	appConfigInfo *ConfigModel // all server config information

)

//server config will used in file variables.go
type ConfigModel struct {
	Gin    *GinModel    `yaml:"gin"`
	Server *ServerModel `yaml:"server"`
	MySql  *MysqlModel  `yaml:"mysql"`
	Redis  *RedisModel  `yaml:"redis"`
}

//serverModel get server information from config.yml
type MysqlModel struct {
	Host        string `yaml:"host"`          // 主机
	Port        string `yaml:"port"`          // 端口
	User        string `yaml:"user"`          // 用户名
	Password    string `yaml:"password"`      // 密码
	Dbname      string `yaml:"dbname"`        // 数据库名
	Prefix      string `yaml:"prefix"`        // 前缀
	MaxIdleConn int    `yaml:"max_idle_conn"` // 最大连接数
	MaxOpenConn int    `yaml:"max_open_conn"` // 设置最大打开连接数
	ShowSql     bool   `yaml:"show_sql"`      // 控制台是否打印SQL语句
}

//serverModel get server information from config.yml
type RedisModel struct {
	Host     string `yaml:"host"`      // 主机
	Port     string `yaml:"port"`      // 端口
	Password string `yaml:"password"`  // 密码
	PoolSize int    `yaml:"pool_size"` // 最大连接数
	DbNum    int    `yaml:"db_num"`    // 设置DB数据库
}

//ServerModel get server information from config.yml
type ServerModel struct {
	SessionSecret string `yaml:"session_secret"` // 设置会话密钥（默认随机生成）
	UploadPath    string `yaml:"upload_path"`    // 设置文件上传路径（默认为工作目录下面的 upload 目录）
	SqlitePath    string `yaml:"sqlite_path"`    // 修改默认的 SQLite 数据库文件的位置（默认在工作目录下，名称为 .go-file.db）
}

//GinModel get server information from config.yml
type GinModel struct {
	Mode               string `yaml:"mode"`                 // gin运行模式
	Host               string `yaml:"host"`                 // 运行访问域
	Port               string `yaml:"port"`                 // 端口
	EnableConsoleRoute bool   `yaml:"enable_console_route"` // 控制台是否输出路由
}

//LoadConfigInformation load config information for application
func loadConfigInformation(configPath string, yml string, structObject interface{}) (err error) {
	var (
		filePath string
		wr       string
	)
	if configPath == "" {
		wr, _ = os.Getwd()
		wr = path.Join(wr, "config")
	} else {
		wr = configPath
	}
	WorkSpace = wr
	filePath = path.Join(WorkSpace, yml)
	configData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf(" config file read failed: %s", err)
		os.Exit(-1)
	}
	err = yaml.Unmarshal(configData, structObject)
	if err != nil {
		fmt.Printf(" config parse failed: %s", err)
		os.Exit(-1)
	}
	return nil
}

//
func ReadAppConfig() (*ConfigModel, error) {
	fPath, _ := os.Getwd()
	fPath = path.Join(fPath, "config")
	configPath := flag.String("c", fPath, "config file path")
	flag.Parse()
	err := loadConfigInformation(*configPath, "app.yml", &appConfigInfo)
	return appConfigInfo, err
}
