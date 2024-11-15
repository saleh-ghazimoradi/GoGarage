package config

import (
	"github.com/spf13/viper"
	"time"
)

var Appconfig *Config

type Config struct {
	DBDriver                string        `mapstructure:"DB_DRIVER"`
	DBSource                string        `mapstructure:"DB_SOURCE"`
	ServerAddress           string        `mapstructure:"SERVER_ADDRESS"`
	LogLevel                int8          `mapstructure:"LOG_LEVEL"`
	Env                     string        `mapstructure:"ENV"`
	DBHost                  string        `mapstructure:"DB_HOST"`
	DBPort                  string        `mapstructure:"DB_PORT"`
	DBUser                  string        `mapstructure:"POSTGRES_USER"`
	DBPassword              string        `mapstructure:"POSTGRES_PASSWORD"`
	DBName                  string        `mapstructure:"POSTGRES_DB"`
	DBSSLMode               string        `mapstructure:"DB_SSLMODE"`
	MaxOpenConns            int           `mapstructure:"MAX_OPEN_CONNS"`
	MaxIdleConns            int           `mapstructure:"MAX_IDLE_CONNS"`
	MaxIdleTime             time.Duration `mapstructure:"MAX_IDLE_TIME"`
	Timeout                 time.Duration `mapstructure:"TIMEOUT"`
	RedisPassword           string        `mapstructure:"REDIS_PASSWORD"`
	RedisHost               string        `mapstructure:"REDIS_HOST"`
	RedisPort               string        `mapstructure:"REDIS_PORT"`
	RedisDB                 int           `mapstructure:"REDIS_DB"`
	RedisDialTimeOut        time.Duration `mapstructure:"REDIS_DIALTIMEOUT"`
	RedisReadTimeout        time.Duration `mapstructure:"REDIS_READTIMEOUT"`
	RedisWriteTimeout       time.Duration `mapstructure:"REDIS_WRITETIMEOUT"`
	RedisIdleCheckFrequency time.Duration `mapstructure:"REDIS_IDLECHECKFREQUENCY"`
	RedisPoolSize           int           `mapstructure:"REDIS_POOLSIZE"`
	RedisPoolTimeout        time.Duration `mapstructure:"REDIS_POOLTIMEOUT"`
	RedisIdleTimeout        time.Duration `mapstructure:"REDIS_IDLETIMEOUT"`
}

func LoadingConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return err
	}

	Appconfig = &config
	return nil
}
