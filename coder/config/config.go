/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2019-07-01 17:10:36
# File Name: config.go
# Description:
####################################################################### */

package config

import (
	"fmt"
	"sync"

	"github.com/ant-libs-go/util"
	"github.com/naoina/toml"
)

type Config struct {
	cfg  map[string]map[string]string
	path string
	file string
	lock sync.RWMutex
}

func (this *Config) SetPath(path string) error {
	this.path = path
	this.file = fmt.Sprintf("%s/.ant-coder.toml", this.path)
	return nil
}

func (this *Config) load(file string, cfg map[string]map[string]string) error {
	if exists, _, _ := util.PathExists(file); !exists {
		util.WriteFile("[basic]", file)
	}
	buff, err := util.ReadFile(file)
	if err != nil {
		return fmt.Errorf("Config file read fail, err: %s", err)
	}
	err = toml.Unmarshal([]byte(buff), cfg)
	if err != nil {
		return fmt.Errorf("Config file Unmarshal fail, err: %s", err)
	}
	return nil
}

func (this *Config) Load() (err error) {
	cfg := map[string]map[string]string{}
	if err = this.load(this.file, cfg); err != nil {
		return
	}
	this.lock.Lock()
	this.cfg = cfg
	this.lock.Unlock()
	return
}

func (this *Config) Save() error {
	buff, err := toml.Marshal(this.cfg)
	if err != nil {
		return fmt.Errorf("Config file marshal fail, err: %s", err)
	}
	err = util.WriteFile(string(buff), this.file)
	if err != nil {
		return fmt.Errorf("Config file write fail, err: %s", err)
	}
	return nil
}

func (this *Config) Get() map[string]map[string]string {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return this.cfg
}

func (this *Config) Set(key, key2, val string) *Config {
	this.lock.RLock()
	defer this.lock.RUnlock()
	if _, ok := this.cfg[key]; !ok {
		this.cfg[key] = map[string]string{}
	}
	this.cfg[key][key2] = val
	return this
}

var Default *Config = &Config{}

func SetPathAndLoad(path string) error {
	Default.SetPath(path)
	return Default.Load()
}

func Get(key, key2, defval string) string {
	t, ok := Default.Get()[key]
	if !ok {
		return defval
	}
	v, ok := t[key2]
	if !ok {
		return defval
	}
	return v
}

func Set(key, key2, val string) *Config {
	return Default.Set(key, key2, val)
}

func Save() error {
	return Default.Save()
}

/*
func main() {
	if err := SetPathAndLoad("/home/chenyu/.ant-coder"); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	//fmt.Println(Get())

	Set("niubi", "feichagn", "haha").Set("niubi2", "feichagn2", "haha2").Save()
}
*/
