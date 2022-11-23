package conf

import (
	"github.com/gin-gonic/gin"
	"singo/agollo"
	"singo/model"
	"singo/util"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	_ = godotenv.Load()

	// 从Apollo读取环境变量
	agollo.Load()

	gin.SetMode(agollo.GetString("GIN_MODE"))

	// 设置日志级别
	util.BuildLogger(agollo.GetString("LOG_LEVEL"))

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}

	// 连接数据库
	model.Database(agollo.GetString("MYSQL_DSN"))
}
