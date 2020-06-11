package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var GLOBAL_CONFIG *Config

var conf_file = "config.json"

type Server struct {
	Location string
}
type Config struct {
	Port    int
	Servers []Server
}

func init() {
	filePtr, err := os.Open(conf_file)
	if err != nil {
		fmt.Println("Open file failed [Err:%s]", err.Error())
		return
	}
	defer filePtr.Close()
	var config Config
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Decoder failed", err.Error())
	} else {
		fmt.Println("Decoder success")
		fmt.Println(config.Port)
	}
	GLOBAL_CONFIG = &config
}
