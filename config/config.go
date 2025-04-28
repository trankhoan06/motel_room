package config

import (
	"errors"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	App      AppConfig
	Mysql    MysqlConfig
	AWS      AWSConfig
	Redis    RedisConfig
	Google   GoogleConfig
	Email    EmailSenderConfig
	CronSpec CronSpec
}

type CronSpec struct {
	UpdateStatusBooking string
}

type EmailSenderConfig struct {
	EmailSenderName     string
	EmailSenderAddress  string
	EmailSenderPassword string
}

type GoogleConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Scopes       []string
	Endpoint     any
	Audience     string
}

type AppConfig struct {
	Version string
	Port    string
	Mode    string
	Prefix  string
	Secret  string
	Host    string
	//VerifyKey    *rsa.PublicKey
	//SignKey      *rsa.PrivateKey
	MigrationURL string
}
type MysqlConfig struct {
	Host            string
	ContainerName   string
	Port            string
	User            string
	Password        string
	DBName          string
	MaxOpenConns    int
	MaxIdleConns    int
	MaxConnLifetime string
}

type AWSConfig struct {
	Region         string
	APIKey         string
	SecretKey      string
	S3Bucket       string
	S3Domain       string
	S3FolderImages string
}
type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

// location of the files used for signing and verification
//const (
//	privKeyPath = "secret/app.rsa"     // openssl genrsa -out app.rsa keysize
//	pubKeyPath  = "secret/app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
//)
//
//// initKeys : read the key files before starting http handlers
//func initKeys(cfg *Config) {
//	signBytes, err := os.ReadFile(privKeyPath)
//	fatal(err)
//
//	cfg.App.SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
//	fatal(err)
//
//	verifyBytes, err := os.ReadFile(pubKeyPath)
//	fatal(err)
//
//	cfg.App.VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
//	fatal(err)
//}

func LoadConfig() (*Config, error) {
	v := viper.New()

	v.AddConfigPath("config")
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		// check is not found file config
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	var c Config // Unmarshal data config have get in file config then get into c
	if err := v.Unmarshal(&c); err != nil {
		return nil, err
	}

	//initKeys(&c)

	return &c, nil
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
