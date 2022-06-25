package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Port       string `json:"port"`
	ParserPort string `json:"ParserPort"`
	CrudPort   string `json:"CrudPort"`
}

func InitConfig() (*Config, error) {
	body, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		return nil, err
	}
	c := &Config{}
	if err := json.Unmarshal(body, c); err != nil {
		return c, err
	}
	return c, nil
}
