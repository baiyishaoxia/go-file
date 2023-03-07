package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"go-file/common"
	"go-file/model"
	"go-file/router"
	"html/template"
	"log"
	"strconv"
)

var AppConfig = new(common.ConfigModel)

func loadTemplate() *template.Template {
	var funcMap = template.FuncMap{
		"unescape": common.UnescapeHTML,
	}
	t := template.Must(template.New("").Funcs(funcMap).ParseFS(common.FS, "public/*.html"))
	return t
}

func main() {
	_config, _err := common.ReadAppConfig()
	if _err != nil {
		panic("系统配置文件缺失")
	}
	AppConfig = _config
	common.Init()
	//
	if AppConfig.Gin.Mode != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	// Initialize SQL Database
	_, err := model.InitDB(AppConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Redis
	err = common.InitRedisClient()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Initialize options
	model.InitOptionMap()

	// Initialize HTTP server
	server := gin.Default()
	server.SetHTMLTemplate(loadTemplate())

	// Initialize session store
	if common.RedisEnabled {
		opt := common.ParseRedisOption()
		store, _ := redis.NewStore(opt.MinIdleConns, opt.Network, opt.Addr, opt.Password, []byte(common.SessionSecret))
		server.Use(sessions.Sessions("session", store))
	} else {
		store := cookie.NewStore([]byte(common.SessionSecret))
		server.Use(sessions.Sessions("session", store))
	}

	router.SetRouter(server)
	var realPort = AppConfig.Gin.Port
	if realPort == "" {
		realPort = strconv.Itoa(*common.Port)
	}
	if *common.Host == "localhost" {
		ip := common.GetIp()
		if ip != "" {
			*common.Host = ip
		}
	}
	serverUrl := "http://" + *common.Host + ":" + realPort + "/"
	if !*common.NoBrowser {
		common.OpenBrowser(serverUrl)
	}
	if *common.EnableP2P {
		go common.StartP2PServer()
	}
	err = server.Run(":" + realPort)
	if err != nil {
		log.Println(err)
	}
}
