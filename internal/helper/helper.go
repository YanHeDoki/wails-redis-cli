package helper

import (
	"changeme/internal/define"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

func GetConnection(identity string) (*define.Connection, error) {
	conf := new(define.Config)
	nowpath, _ := os.Getwd()
	data, err := ioutil.ReadFile(nowpath + string(os.PathSeparator) + define.ConfigName)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}
	for _, v := range conf.Connections {
		if v.Identity == identity {
			return v, err
		}
	}
	return nil, errors.New("连接数据不存在")
}
