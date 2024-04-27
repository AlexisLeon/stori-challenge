package conf

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"os"
)

type Config struct {
	Mailer struct {
		ReplyTo string `yaml:"replyTo"`
		Source  string `yaml:"source"`
	}
	Database struct {
		Host     string `yaml:"host"`
		Database string `yaml:"database"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}

var conf = &Config{}

func LoadConfig(configFile string) {
	if configFile == "" {
		configFile = "config.yml"
	}

	f, err := os.Open(configFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to read conf file `config.yml`")
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&conf)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to decode conf file `config.yml`")
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func GetConfig() Config {
	if conf == nil {
		fmt.Fprintln(os.Stderr, "conf not loaded, please initialise with conf.LoadConfig()")
		os.Exit(1)
	}

	return *conf
}
