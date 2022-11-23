package agollo

import (
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/agcache"
	"github.com/apolloconfig/agollo/v4/env/config"
	"os"
	"singo/util"
)

var (
	Client           agollo.Client
	defaultNamespace agcache.CacheInterface
)

func Load() {
	c := &config.AppConfig{
		AppID:         os.Getenv("AppID"),
		Cluster:       os.Getenv("Cluster"),
		IP:            os.Getenv("IP"),
		NamespaceName: os.Getenv("NamespaceName"),
		Secret:        os.Getenv("Secret"),
	}
	Client, _ = agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	defaultNamespace = Client.GetConfigCache(c.NamespaceName)
	//启动监听，说是监听，实际上是被回调
	Client.AddChangeListener(&CustomChangeListener{})
}

func GetString(key string) string {
	value, err := defaultNamespace.Get(key)
	if err != nil {
		util.Log().Error("获取apollo配置失败", err)
	}
	return value.(string)
}

func GetInt(key string) int {
	value, err := defaultNamespace.Get(key)
	if err != nil {
		util.Log().Error("获取apollo配置失败", err)
	}
	return value.(int)
}

func GetFloat32(key string) float32 {
	value, err := defaultNamespace.Get(key)
	if err != nil {
		util.Log().Error("获取apollo配置失败", err)
	}
	return value.(float32)
}
