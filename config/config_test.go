package config

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/kekik/viper"
	yaml "gopkg.in/yaml.v2"
)

const customValues = `gofshare:
  basepath: /home/data
  baseuri: https://localhost
  format: /var/files/{EXT}/{FILENAME}
`

func createConfigFile(content string, name string, path string) string {
	if err := ioutil.WriteFile(path+"/"+name, []byte(content), 0600); err != nil {
		log.Fatal(err)
	}
	return path + "/" + name
}

func TestInitCustom(t *testing.T) {

	tmpfile := createConfigFile(customValues, "gofshare.yml", "/tmp")
	defer os.Remove(tmpfile) // clean up
	Init(tmpfile)

	cfg := values()
	var expected config
	err := yaml.Unmarshal([]byte(customValues), &expected)
	if err != nil {
		t.Fatal(err)
	}

	if cfg != expected {
		t.Error("values are not correct")
	}
	viper.Reset()
}

func TestInitEnvironment(t *testing.T) {
	var expected config
	err := yaml.Unmarshal([]byte(customValues), &expected)
	if err != nil {
		t.Fatal(err)
	}

	//The key consist of :
	//	Prefix : GFS
	//	SubObject: GOFSHARE
	//	KEY:	BASEPATH
	os.Setenv("GFS_GOFSHARE_BASEPATH", expected.GoFShare.BasePath)
	defer os.Unsetenv("GFS_GOFSHARE_BASEPATH")

	os.Setenv("GFS_GOFSHARE_BASEURI", expected.GoFShare.BaseURI)
	defer os.Unsetenv("GFS_GOFSHARE_BASEURI")

	os.Setenv("GFS_GOFSHARE_FORMAT", expected.GoFShare.Format)
	defer os.Unsetenv("GFS_GOFSHARE_FORMAT")

	//no config file, load from environment variables
	Init("")

	//load the current config
	cfg := values()

	if cfg != expected {
		t.Error("values are not correct")
	}
	viper.Reset()
}
