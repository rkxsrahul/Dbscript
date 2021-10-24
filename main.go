package main

import (
	"flag"
	"log"
	"os"

	"dbscript/config"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	//conf flag is used to set the way user want configuration of service to be set
	// It can be via TOML or ENVIRONMENT
	conf := flag.String("conf", "environment", "set configuration from toml file or environment variables")
	//file flag is used to give path to toml file if conf is set with TOML
	file := flag.String("file", "", "set path of toml file")
	//Below function parses flags used as command line arguement
	flag.Parse()

	//check if conf is set to environment if yes call function to set configuration using environment
	if *conf == "environment" {
		log.Println("environment")
		config.ConfigurationWithEnv()
	} else if *conf == "toml" { //if conf is set to TOML set configuration using file
		log.Println("toml")
		if *file == "" {
			log.Println("Please pass toml file path")
			os.Exit(1)
		} else { // set configuration by calling below function
			err := config.ConfigurationWithToml(*file)
			if err != nil {
				log.Println("Error in parsing toml file")
				log.Println(err)
				os.Exit(1)
			}
		}
	} else {
		log.Println("Please pass valid arguments, conf can be set as toml or environment")
		os.Exit(1)
	}
	/*==========================Database======================================*/
	dbConfig := config.DBConfig()
	// connecting db using connection string
	db, err := gorm.Open("postgres", dbConfig)
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()
	config.DB = db
	/*==========================Routing======================================*/
	r := gin.Default()
	r.GET("/ping")
	r.Run(":" + config.Conf.Service.Port)

}
