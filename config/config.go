package config

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
)

// Config is a structure for configuration
type Config struct {
	Database Database
	Service  Service
}

type Service struct {
	Port string
}

// Database is a structure for postgres database configuration
type Database struct {
	Name  string
	Host  string
	Port  string
	User  string
	Pass  string
	Ssl   string
	Ideal string
}

var (
	// Conf is a global variable for configuration
	Conf Config
	// TomlFile is a global variable for toml file path
	TomlFile string
	// DB Database client
	DB *gorm.DB
)

/*==========================Configuration using Env======================================*/
// ConfigurationWithEnv is a method to initialize configuration with environment variables
func ConfigurationWithEnv() {

}

/*==========================Configuration uing toml file======================================*/
// ConfigurationWithToml is a method to initialize configuration with toml file
func ConfigurationWithToml(filePath string) error {
	// set varible as file path if configuration is done using toml
	TomlFile = filePath
	log.Println(filePath)
	// parse toml file and save data config structure
	_, err := toml.DecodeFile(filePath, &Conf)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

/*==========================Database connection======================================*/
// DBConfig is a method that return postgres database connection string
func DBConfig() string {
	//again reset the config if any changes in toml file or environment variables
	//	SetConfig()
	// creating postgres connection string
	str := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		Conf.Database.Host,
		Conf.Database.Port,
		Conf.Database.User,
		Conf.Database.Pass,
		Conf.Database.Name,
		Conf.Database.Ssl)
	return str
}
