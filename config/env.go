package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
)

type Env struct {
	Model    string `yaml:"model"`
	Node     string `yaml:"node"`
	Newrelic struct {
		AppName    string `yaml:"app_name"`
		LicenseKey string `yaml:"license_key"`
	} `yaml:"newrelic"`
	Schedules []string `yaml:"schedules"`
}

var CurrentEnv Env

// 初始化环境
func InitEnv() {
	path_str, _ := filepath.Abs("config/env.yaml")
	content, err := ioutil.ReadFile(path_str)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = yaml.Unmarshal(content, &CurrentEnv)
	if err != nil {
		log.Fatal(err)
	}
}
