package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func InitConfig() (*Config, error) {
	body, err := ioutil.ReadFile("../config/config.json")
	if err != nil {
		return nil, err
	}
	c := &Config{}
	if err := json.Unmarshal(body, c); err != nil {
		return c, err
	}
	return c, nil
}
