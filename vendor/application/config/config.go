package config

import (
	"errors"
	"github.com/BurntSushi/toml"
	"os"
	"strings"
)

var Config map[string]interface{}

/*
Config get config value by path. if path not exists, return defaultVal

Config use TOML format, so you should know the real type that you get from Config
read more https://github.com/toml-lang/toml

[one]
[[one.two]]
key=val

use path "one.two.key" to get "val"
*/
func ConfigParse(path string, defaultVal ...interface{}) (interface{}, error) {

	if path == "" {
		return nil, errors.New("can't get config with empty path")
	}
	paths := strings.Split(path, ".")

	var ok bool
	var val interface{}
	mapData := Config
	for _, p := range paths {
		_, ok = mapData[p]
		if !ok {
			break
		}

		tmpMapData, ok := mapData[p].(map[string]interface{})
		if !ok {
			val = mapData[p]
			break
		}

		mapData = tmpMapData
	}

	//we had get last data
	if val != nil {
		return val, nil
	}

	if len(defaultVal) == 0 {
		return nil, errors.New("no config value " + path)
	}

	//use default value
	return defaultVal[0], nil
}

func ConfigMustString(path string, defaultVal ...string) string {
	defaulValInter := make([]interface{}, 0, len(defaultVal))
	for _, v := range defaultVal {
		defaulValInter = append(defaulValInter, v)
	}

	val, _ := ConfigParse(path, defaulValInter...)
	switch val.(type) {
	case string:
		return val.(string)
	}

	return ""
}

func GetConfig(configFilePath string) map[string]interface{} {
	if len(configFilePath) == 0 {
		panic("GetConfig: not found config file path")
	}

	Config = make(map[string]interface{})

	if string(configFilePath[0]) != "/" || string(configFilePath[0]) != "\\" {
		configFilePath = string(os.PathSeparator) + configFilePath
	}
	//init toml data

	fi, err := os.Open(configFilePath)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	_, err = toml.DecodeFile(configFilePath, &Config)

	if err != nil {
		panic(err)
	}

	return Config
}
