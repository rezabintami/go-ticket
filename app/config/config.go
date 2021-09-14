package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Debug bool `mapstructure:"DEBUG"`

	//! Server
	SERVER_PORT    string `mapstructure:"SERVER_PORT"`
	SERVER_TIMEOUT int    `mapstructure:"SERVER_TIMEOUT"`

	//! MYSQL
	MYSQL_DB_HOST string `mapstructure:"MYSQL_DB_HOST"`
	MYSQL_DB_PORT string `mapstructure:"MYSQL_DB_PORT"`
	MYSQL_DB_USER string `mapstructure:"MYSQL_DB_USER"`
	MYSQL_DB_PASS string `mapstructure:"MYSQL_DB_PASS"`
	MYSQL_DB_NAME string `mapstructure:"MYSQL_DB_NAME"`

	//! MONGO DB
	MONGO_DB_HOST string `mapstructure:"MONGO_DB_HOST"`
	MONGO_DB_PORT string `mapstructure:"MONGO_DB_PORT"`
	MONGO_DB_USER string `mapstructure:"MONGO_DB_USER"`
	MONGO_DB_PASS string `mapstructure:"MONGO_DB_PASS"`
	MONGO_DB_NAME string `mapstructure:"MONGO_DB_NAME"`

	//! OUATH2 GOOGLE
	GOOGLE_AUTH_CLIENT string `mapstructure:"GOOGLE_AUTH_CLIENT"`
	GOOGLE_AUTH_SECRET string `mapstructure:"GOOGLE_AUTH_SECRET"`

	//! OAUTH2 FACEBOOK
	FACEBOOK_AUTH_CLIENT string `mapstructure:"FACEBOOK_AUTH_CLIENT"`
	FACEBOOK_AUTH_SECRET string `mapstructure:"FACEBOOK_AUTH_SECRET"`
}

func GetConfig() Config {
	var conf Config

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		viper.BindEnv(conf.MYSQL_DB_HOST)
		viper.BindEnv(conf.MYSQL_DB_PORT)
		viper.BindEnv(conf.MYSQL_DB_USER)
		viper.BindEnv(conf.MYSQL_DB_PASS)
		viper.BindEnv(conf.MYSQL_DB_NAME)

		viper.BindEnv(GetConfig().FACEBOOK_AUTH_CLIENT)
		viper.BindEnv(GetConfig().FACEBOOK_AUTH_SECRET)

		viper.BindEnv(GetConfig().GOOGLE_AUTH_CLIENT)
		viper.BindEnv(GetConfig().GOOGLE_AUTH_SECRET)
	}

	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}
	return conf
}
