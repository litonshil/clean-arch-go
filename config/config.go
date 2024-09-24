package config

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"

	_ "github.com/spf13/viper/remote"
)

const (
	LOCAL = "LOCAL"
)

type Config struct {
	App   *AppConfig
	Db    *DbClient
	Cache *CacheClient
	Queue *QueueClient
}

// AppConfig application specific config
type AppConfig struct {
	Name         string
	Port         int
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
	AppKeyHeader string
	AppKey       string
	ENV          string
	LogLevel     string
}

type ServiceConfig struct {
	ServiceURL   string
	AppKeyHeader string
	AppKey       string
	Timeout      time.Duration
}

type DbClient struct {
	Master   *DBConfig
	Replica1 *DBConfig
}

type CacheClient struct {
	Redis *RedisConfig
}

type QueueClient struct {
	Asynq *AsynqConfig
}

// DBConfig database specific config
type DBConfig struct {
	Host        string
	Port        int
	Name        string
	Username    string
	Password    string
	MaxLifeTime time.Duration
	MaxIdleConn int
	MaxOpenConn int
	Debug       bool
	PrepareStmt bool
}

type RedisConfig struct {
	Host            string
	Port            int
	Name            string
	Username        string
	Password        string
	ValueExpiredIn  int // seconds
	MaxIdleConn     int
	MaxOpenConn     int
	Database        int
	MandatoryPrefix string
}

type AsynqConfig struct {
	RedisAddr        string
	DB               int
	Password         string
	Concurrency      int
	Queue            string
	SyncTaskQueue    string
	Retention        time.Duration // in hours
	UniquenessTTL    time.Duration
	RetryCount       int
	TaskExecTimeUnit string
}

// c is the configuration instance
var c Config

// Get returns all configurations
func Get() Config {
	return c
}

func App() *AppConfig {
	return c.App
}

func DB() *DbClient {
	return c.Db
}

func Cache() *CacheClient {
	return c.Cache
}

func Queue() *QueueClient {
	return c.Queue
}

// Load the config
func Load() error {
	setDefaultConfig()

	_ = viper.BindEnv("env")
	_ = viper.BindEnv("CONSUL_URL")
	_ = viper.BindEnv("CONSUL_PATH")

	consulURL := viper.GetString("CONSUL_URL")
	consulPath := viper.GetString("CONSUL_PATH")

	if consulURL != "" && consulPath != "" {
		viper.SetConfigType("json")
		// Use the correct provider and path
		err := viper.AddRemoteProvider("consul", consulURL, consulPath)
		if err != nil {
			log.Fatalf("Failed to add remote provider: %v", err)
		}

		err = viper.ReadRemoteConfig()

		log.Printf("consul_url = %v \nconsul_path = %v", consulURL, consulPath)

		if err != nil {
			log.Printf("Failed to read remote config: %v", err)
			//log.Printf("%s named \"%s\"", err.Error(), consulPath)
		}

		c = Config{}

		if err := viper.Unmarshal(&c); err != nil {
			panic(err)
		}

		if r, err := json.MarshalIndent(&c, "", "  "); err == nil {
			fmt.Println(string(r))
		}
	} else {
		fmt.Println("CONSUL_URL or CONSUL_PATH is missing from ENV")
	}

	return nil
}

func ReadDotENV() string {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	env := viper.Get("ENV")
	if env != nil {
		return env.(string)
	}
	return ""
}

func setDefaultConfig() {
	env := ReadDotENV()

	c.App = &AppConfig{
		Name:         "SMART-MESS",
		Port:         8080,
		ReadTimeout:  30,
		WriteTimeout: 30,
		IdleTimeout:  30,
		AppKeyHeader: "app-key",
		AppKey:       "appkey",
		ENV:          env,
	}

	c.Db = &DbClient{
		Master: &DBConfig{
			Host:        "localhost",
			Name:        "smart-mess",
			Port:        3306,
			Username:    "root",
			Password:    "",
			MaxLifeTime: 30,
			MaxIdleConn: 1,
			MaxOpenConn: 2,
			Debug:       true,
			PrepareStmt: true,
		},
		Replica1: &DBConfig{
			Host:     "localhost",
			Name:     "menu",
			Port:     3306,
			Username: "root",
			Password: "",
			Debug:    true,
		},
	}

	if c.App.ENV == LOCAL {
		c.Db.Master.Port = 3306
		c.Db.Master.Username = "root"
		c.Db.Master.Password = ""

		c.Db.Replica1.Port = 3306
		c.Db.Replica1.Username = "root"
		c.Db.Replica1.Password = ""
	}

	c.Cache = &CacheClient{
		Redis: &RedisConfig{
			Host:            "127.0.0.1",
			Port:            6379,
			Username:        "",
			Password:        "",
			Database:        4,
			ValueExpiredIn:  0,
			MandatoryPrefix: "menu_",
		},
	}

	c.Queue = &QueueClient{
		Asynq: &AsynqConfig{
			RedisAddr:        "127.0.0.1:6379",
			DB:               15,
			Concurrency:      10,
			Queue:            "menu",
			SyncTaskQueue:    "menu_sync_tasks",
			Retention:        168,
			RetryCount:       25,
			UniquenessTTL:    1,
			TaskExecTimeUnit: "SECOND",
		},
	}

}
