package config

import (
	"encoding/json"
	"log"
	"os"
)

// FinancialConfig provides configuartion information about the financial application
type FinancialConfig struct {
	// HttpServer provides config info to set up HTTP server
	HTTPServer struct {
		APIPath string
		Port    int
	}
	Integration struct {
		// TraceTriggeredByClientEnabled enables client to activate trace in the request
		TraceTriggeredByClientEnabled bool
	}
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
	Config.HTTPServer.APIPath = "/financial/api"
	Config.HTTPServer.Port = 8080
	Config.Integration.TraceTriggeredByClientEnabled = true

	AppHome, AppHomeNotSet = os.LookupEnv("APP_HOME")
	log.Printf("From environment AppHome=%s\n", AppHome)
	log.Printf("AppHomeNotSet=%v", AppHomeNotSet)
	if !AppHomeNotSet {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		AppHome = dir
	}
	LoadConfigFromAppHome(AppHome)
}

func LoadConfigFromAppHome(appHome string) {
	log.Printf("Given AppHome=%s\n", appHome)
	if err := os.Chdir(appHome); err == nil {
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		AppHome = dir
		log.Printf("Absolute AppHome=%s\n", AppHome)
		defaultConfigFile := AppHome + "/configs/financial.json"
		if err := LoadConfig(defaultConfigFile); err != nil {
			log.Printf("Configuration file=[%s] not found", defaultConfigFile)
		}
	} else {
		panic(err)
	}
}

// LoadConfig loads configuration from given file
func LoadConfig(pathFileName string) error {
	f, err := os.Open(pathFileName)
	if err != nil {
		return err
	}
	log.Printf("Loading configuration from file=[%s]", pathFileName)
	dec := json.NewDecoder(f)
	if err := dec.Decode(&Config); err != nil {
		f.Close()
		return err
	}
	f.Close()
	return nil
}
