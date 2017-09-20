package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/kekik/viper"
)

type goFShareConfig struct {
	BasePath, BaseURI, Format string
}

type config struct {
	GoFShare goFShareConfig
}

// Init reads in config file and ENV variables if set.
func Init(cfgFile string) {

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	//This prefix will be used when setting environment variables
	//Ex: GFS_GOFSHARE_BASEPATH=VALUE
	viper.SetEnvPrefix("gfs")
	// name of config file (without extension)
	viper.SetConfigName("gofshare")
	// adding home directory as first search path
	viper.AddConfigPath("$HOME/.gofshare")
	// adding the current directory as a search path
	viper.AddConfigPath(".")
	//read in environment variables that match
	viper.AutomaticEnv()

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		fmt.Println("Loading from local file")
	}

	err := viper.ReadInConfig()
	if err != nil {
		e, ok := err.(viper.ConfigParseError)
		if ok {
			log.Fatalf("error parsing config file: %v", e)
		}
		log.Printf("No config file used")
	} else {
		log.Printf("Using config file: %v", viper.ConfigFileUsed())
	}

}

//Values return a config object with the loaded data
func values() config {
	return config{
		GoFShare: goFShareConfig{
			BasePath: viper.GetString("GoFShare.BasePath"),
			BaseURI:  viper.GetString("GoFShare.BaseURI"),
			Format:   viper.GetString("GoFShare.Format"),
		},
	}
}
