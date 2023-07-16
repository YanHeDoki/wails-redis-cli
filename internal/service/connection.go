package service

import (
	"changeme/internal/define"
	"encoding/json"
	"errors"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"os"
)

func ConnectionList() (*define.Config, error) {
	nowpath, _ := os.Getwd()
	data, err := ioutil.ReadFile(nowpath + string(os.PathSeparator) + define.ConfigName)
	if err != nil {
		return nil, err
	}
	conf := new(define.Config)
	err = json.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func ConnectionCreate(connection *define.Connection) error {
	if connection.Addr == "" {
		return errors.New("地址不能为空")
	}
	if connection.Name == "" {
		connection.Name = connection.Addr
	}
	if connection.Port == "" {
		connection.Port = "6379"
	}
	connection.Identity = uuid.NewV4().String()
	conf := new(define.Config)
	nowpath, _ := os.Getwd()
	data, err := ioutil.ReadFile(nowpath + string(os.PathSeparator) + define.ConfigName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			conf.Connections = []*define.Connection{connection}
			data, _ := json.Marshal(conf)
			os.MkdirAll(nowpath, 0600)
			ioutil.WriteFile((nowpath + string(os.PathSeparator) + define.ConfigName), data, 0600)
			return nil
		}
		return err
	}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		return err
	}
	conf.Connections = append(conf.Connections, connection)
	data, err = json.Marshal(&conf)
	if err != nil {
		return err
	}
	ioutil.WriteFile((nowpath + string(os.PathSeparator) + define.ConfigName), data, 0600)
	return nil
}

//编辑

func ConnectionEdit(connection *define.Connection) error {
	if connection.Identity == "" {
		return errors.New("连接唯一标识不能为空")
	}
	if connection.Addr == "" {
		return errors.New("地址不能为空")
	}
	if connection.Name == "" {
		connection.Name = connection.Addr
	}
	if connection.Port == "" {
		connection.Port = "6379"
	}
	conf := new(define.Config)
	nowpath, _ := os.Getwd()
	data, err := ioutil.ReadFile(nowpath + string(os.PathSeparator) + define.ConfigName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		return err
	}
	for i, v := range conf.Connections {
		if v.Identity == connection.Identity {
			conf.Connections[i] = connection
		}
	}
	data, err = json.Marshal(&conf)
	if err != nil {
		return err
	}
	ioutil.WriteFile((nowpath + string(os.PathSeparator) + define.ConfigName), data, 0600)
	return nil
}

func ConnectionDelete(identity string) error {
	if identity == "" {
		return errors.New("连接唯一标识不能为空")
	}
	conf := new(define.Config)
	nowpath, _ := os.Getwd()
	data, err := ioutil.ReadFile(nowpath + string(os.PathSeparator) + define.ConfigName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		return err
	}
	for i, v := range conf.Connections {
		if v.Identity == identity {
			conf.Connections = append(conf.Connections[:i], conf.Connections[i+1:]...)
		}
	}
	data, err = json.Marshal(&conf)
	if err != nil {
		return err
	}
	ioutil.WriteFile((nowpath + string(os.PathSeparator) + define.ConfigName), data, 0600)
	return nil
}
