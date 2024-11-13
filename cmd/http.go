/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/saleh-ghazimoradi/GoGarage/config"
	"github.com/saleh-ghazimoradi/GoGarage/internal/gateway"
	"github.com/saleh-ghazimoradi/GoGarage/logger"
	"github.com/saleh-ghazimoradi/GoGarage/utils"
	"github.com/spf13/cobra"
	"log"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "launching the http rest listen server",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Logger.Info("server has started", "addr", config.Appconfig.ServerAddress, "env", config.Appconfig.Env)

		cfg := utils.PostgresConfig{
			Host:         config.Appconfig.DBHost,
			Port:         config.Appconfig.DBPort,
			User:         config.Appconfig.DBUser,
			Password:     config.Appconfig.DBPassword,
			Database:     config.Appconfig.DBName,
			SSLMode:      config.Appconfig.DBSSLMode,
			MaxIdleTime:  config.Appconfig.MaxIdleTime,
			MaxIdleConns: config.Appconfig.MaxIdleConns,
			MaxOpenConns: config.Appconfig.MaxOpenConns,
			Timeout:      config.Appconfig.Timeout,
		}

		rediscfg := utils.RedisConfig{
			Host:               config.Appconfig.RedisHost,
			Port:               config.Appconfig.RedisPort,
			Password:           config.Appconfig.RedisPassword,
			Db:                 config.Appconfig.RedisDB,
			DialTimeout:        config.Appconfig.RedisDialTimeOut,
			ReadTimeout:        config.Appconfig.RedisReadTimeout,
			WriteTimeout:       config.Appconfig.RedisWriteTimeout,
			IdleCheckFrequency: config.Appconfig.RedisIdleCheckFrequency,
			PoolSize:           config.Appconfig.RedisPoolSize,
			PoolTimeout:        config.Appconfig.RedisPoolTimeout,
			IdleTimeout:        config.Appconfig.RedisIdleTimeout,
		}

		redis, err := utils.RedisConnection(rediscfg)
		if err != nil {
			log.Fatal(err)
		}
		defer redis.Close()

		db, err := utils.PostConnection(cfg)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		routeHandlers := gateway.Handlers{}

		if err := gateway.Server(gateway.Routes(routeHandlers)); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
