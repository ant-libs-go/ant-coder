/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-09-12 15:11:55
# File Name: config.go
# Description:
####################################################################### */

package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"github.com/ant-libs-go/util"
	"github.com/naoina/toml"
)

type AppConfig struct {
	// basic info
	Basic *struct {
		Debug   bool
		Port    string
		LogFile string
	}
	Gc *struct {
		Percent int // default 100
	}
	Db map[string]*struct {
		Host string
		Port string
		User string
		Pawd string
		Name string
	}
}

type Config struct {
	cfg  *AppConfig
	file string
	lock sync.RWMutex
}

func (this *Config) SetFile(file string) error {
	this.file = file
	return nil
}

func (this *Config) load(file string, cfg *AppConfig) error {
	if len(file) == 0 {
		return fmt.Errorf("Config file not specified")
	}
	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("Config file open fail, err: %s", err)
	}
	defer f.Close()

	buff, err := ioutil.ReadAll(f)
	if err != nil {
		return fmt.Errorf("Config file read fail, err: %s", err)
	}

	err = toml.Unmarshal(buff, cfg)
	if err != nil {
		return fmt.Errorf("Config file Unmarshal fail, err: %s", err)
	}
	return nil
}

func (this *Config) Load() (err error) {
	var cfg AppConfig
	if err = Default.load(this.file, &cfg); err != nil {
		return
	}

	localConfigFile := this.GetLocalConfigFile()
	if len(localConfigFile) != 0 {
		if err = Default.load(localConfigFile, &cfg); err != nil {
			fmt.Println(fmt.Sprintf("Local %s", err))
		}
	}

	this.lock.Lock()
	this.cfg = &cfg
	this.lock.Unlock()
	return
}

func (this *Config) GetLocalConfigFile() string {
	var suffix = ".toml"
	f := fmt.Sprintf("%s-local%s", strings.TrimSuffix(this.file, suffix), suffix)
	exists, isdir, err := util.PathExists(f)
	if err != nil || exists == false || isdir == true {
		return ""
	}
	return f
}

func (this *Config) Get() *AppConfig {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return this.cfg
}

var Default *Config = &Config{}

func SetFile(file string) error {
	return Default.SetFile(file)
}

func SetFileAndLoad(file string) error {
	Default.SetFile(file)
	return Default.Load()
}

func Get() *AppConfig {
	return Default.Get()
}
