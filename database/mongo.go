package database

import (
	"context"
	"fmt"
	"sync"

	_ "compress/zlib"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoDB   *mongo.Database
	mongoOnce sync.Once
)

func GetMongoDB(v *viper.Viper) *mongo.Database {
	if !v.GetBool("mongo.enable") {
		return nil
	}
	mongoOnce.Do(func() {
		uri := fmt.Sprintf("mongodb://%s:%d", v.GetString("mongo.host"), v.GetInt("mongo.port"))
		opts := options.Client().ApplyURI(uri).SetAuth(options.Credential{
			// AuthSource: "admin",
			Username: v.GetString("mongo.username"),
			Password: v.GetString("mongo.password"),
		})
		opts.SetMaxPoolSize(5)
		opts.SetCompressors([]string{"zlib"})
		// Create a new client and connect to the server
		client, err := mongo.Connect(context.TODO(), opts)
		if err != nil {
			panic(err)
		}
		mongoDB = client.Database(v.GetString("mongo.dbname"))
	})
	return mongoDB
}
