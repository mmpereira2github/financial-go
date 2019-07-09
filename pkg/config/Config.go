package config

import (
	"encoding/json"
	"log"
	"os"
)

// FinancialConfig provides configuartion information about the financial application
type FinancialConfig struct {
	DaoConfig struct {
		IndexDaoImplName string
	}
}

// AppHomeNotSet says whether AppHome was specified or not in the environment variable
var AppHomeNotSet = false

// AppHome specifies the directory where the financial application is installed
var AppHome = "."

// Config provides access to the configuration of the financail application
var Config = FinancialConfig{}

func init() {
	AppHome, AppHomeNotSet = os.LookupEnv("APP_HOME")
	log.Printf("AppHome=%s\n", AppHome)
	log.Printf("AppHomeNotSet=%v", AppHomeNotSet)
	if !AppHomeNotSet {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		AppHome = dir
	}
	defaultConfigFile := AppHome + "/configs/financial.json"
	if err := LoadConfig(defaultConfigFile); err != nil {
		log.Printf("Default configuration file=[%s] not found", defaultConfigFile)
	}
}

// LoadConfig loads configuration from given file
func LoadConfig(pathFileName string) error {
	f, err := os.Open(pathFileName)
	if err != nil {
		return err
	}
	dec := json.NewDecoder(f)
	if err := dec.Decode(&Config); err != nil {
		f.Close()
		return err
	}
	f.Close()
	return nil
}
