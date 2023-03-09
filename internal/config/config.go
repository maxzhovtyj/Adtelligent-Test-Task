package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

const (
	configsDir = "configs"
	configFile = "config"

	httpKey = "http"
	authKey = "auth"
	dbKey   = "db"

	httpHostEnv      = "HTTP_HOST"
	passwordSaltEnv  = "PASSWORD_SALT"
	jwtSigningKeyEnv = "JWT_SIGNING_KEY"
	mySQLUserEnv     = "DB_USER"
	mySQLPasswordEnv = "DB_PASSWORD"
)

type (
	Config struct {
		HTTP HTTPConfig
		Auth AuthConfig
		DB   DBConfig
	}

	AuthConfig struct {
		JWT          JWTConfig
		PasswordSalt string
	}

	JWTConfig struct {
		AccessTokenTTL  time.Duration `mapstructure:"accessTokenTTL"`
		RefreshTokenTTL time.Duration `mapstructure:"refreshTokenTTL"`
		SigningKey      string
	}

	DBConfig struct {
		User     string
		Password string
		Database string `mapstructure:"name"`
	}

	HTTPConfig struct {
		Host               string        `mapstructure:"host"`
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderMegabytes"`
	}
)

func Init() (*Config, error) {
	var cfg Config

	if err := parseConfigFile(); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)

	if err := unmarshalConfig(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func setFromEnv(cfg *Config) {
	cfg.HTTP.Host = os.Getenv(httpHostEnv)

	cfg.Auth.PasswordSalt = os.Getenv(passwordSaltEnv)
	cfg.Auth.JWT.SigningKey = os.Getenv(jwtSigningKeyEnv)

	cfg.DB.User = os.Getenv(mySQLUserEnv)
	cfg.DB.Password = os.Getenv(mySQLPasswordEnv)
}

func unmarshalConfig(cfg *Config) error {
	if err := viper.UnmarshalKey(httpKey, &cfg.HTTP); err != nil {
		return err
	}

	if err := viper.UnmarshalKey(authKey, &cfg.Auth.JWT); err != nil {
		return err
	}

	if err := viper.UnmarshalKey(dbKey, &cfg.DB); err != nil {
		return err
	}

	return nil
}

func parseConfigFile() error {
	viper.AddConfigPath(configsDir)
	viper.SetConfigName(configFile)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return nil
}
