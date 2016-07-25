package db

import (
	"fmt"

	log "github.com/cihub/seelog"
	"github.com/focinfi/sakura/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/qor/media_library"
	"github.com/qor/publish"
	"github.com/qor/sorting"
	"github.com/qor/validations"
	"gopkg.in/redis.v3"
)

var (
	// Redis for redis
	Redis *redis.Client

	// DB for gorm
	DB *gorm.DB

	// Publish for content audit
	Publish *publish.Publish
)

func init() {
	var err error

	addr := fmt.Sprintf("%s:%s", config.Config.Redis.Addr, config.Config.Redis.Port)
	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Config.Redis.Password,
		DB:       0, // use default DB
	})

	_, err = Redis.Ping().Result()
	if err != nil {
		log.Errorf("[Redis] failed to connect the redis server, err: %v", err)
	} else {
		log.Infof("[Redis is connectted]")
	}

	dbConfig := config.Config.DataBase
	DB, err = gorm.Open("postgres", fmt.Sprintf("host='%v' port='%v' user='%v' password='%v' dbname='%v' sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Name))

	// SetMaxOpenConns
	DB.DB().SetMaxOpenConns(50)

	if err == nil {
		// switch to log mode if environment is not production
		if !config.IsProduction() {
			DB.LogMode(true)
		}

		sorting.RegisterCallbacks(DB)
		media_library.RegisterCallbacks(DB)
		validations.RegisterCallbacks(DB)
	} else {
		panic(err)
	}
}
