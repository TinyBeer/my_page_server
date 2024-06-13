package database

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/spf13/viper"
)

func GetES(v *viper.Viper) *elasticsearch.TypedClient {
	if !v.GetBool("elasticsearch.enable") {
		return nil
	}
	// ES 配置
	cfg := elasticsearch.Config{
		Addresses: v.GetStringSlice("elasticsearch.addresses"),
	}

	// 创建客户端连接
	typedClient, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		panic(err)
	}
	return typedClient
}
