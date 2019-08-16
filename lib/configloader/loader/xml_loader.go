package loader

import (
	"bufio"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"os"
)

//XMLLoader xml配置文件加载队对象
type XMLLoader struct {
	BaseConfLoader
	confPath string
}

//LoadConfigFromFile 读取配置
//TODO：xml读取
func (x *XMLLoader) LoadConfigFromFile(fileName string) error {
	x.BaseConfLoader.lock.Lock()
	defer x.BaseConfLoader.lock.Unlock()
	if fileName == "" {
		fileName = os.Getenv(EnvConfBasePath)
		if fileName == "" {
			fileName = EnvConfBasePath
		}
	}
	if info, _ := os.Stat(fileName); info == nil {
		return errors.New("load conf error")
	}
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	bs, err := ioutil.ReadAll(bufio.NewReader(file))
	if err != nil {
		return err
	}
	err = xml.Unmarshal(bs, &x.BaseConfLoader.confMap)
	if err != nil {
		return err
	}
	x.confPath = fileName
	return nil
}

//LoadConfigFromFileReader 从fileReader对象中读取对象
func (x *XMLLoader) LoadConfigFromFileReader(file *os.File) error {
	x.BaseConfLoader.init()
	x.BaseConfLoader.lock.Lock()
	defer x.BaseConfLoader.lock.Unlock()
	bs, err := ioutil.ReadAll(bufio.NewReader(file))
	if err != nil {
		return err
	}
	err = xml.Unmarshal(bs, &x.BaseConfLoader.confMap)
	if err != nil {
		return err
	}
	x.confPath = file.Name()
	return nil
}

func (x *XMLLoader) ReLoadConf() {
	x.LoadConfigFromFile(x.confPath)
}

func (x *XMLLoader) GetFileName() string {
	return x.confPath
}
