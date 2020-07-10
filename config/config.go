// This package is config loading tools
package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// YAML struct
type Config struct {
	Database struct {
		Addr string `yarm: "addr"`
		Port int    `yarm: "port"`
		User string `yarm: "user"`
		Pass string `yarm: "pass"`
	} `yarm: "Database"`
	Slack struct {
		Token string `yarm: "token"`
	} `yarm: "Slack"`
}

func LoadConfigration(path string) Config {
	var config Config
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return config
}
