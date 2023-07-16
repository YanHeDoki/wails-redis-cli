package service

import (
	"changeme/internal/define"
	"changeme/internal/helper"
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"strconv"
	"strings"
)

func DbList(identity string) ([]*define.DbItem, error) {
	connection, err := helper.GetConnection(identity)
	if err != nil {
		return nil, err
	}
	//获取数据库连接对象
	rdb := redis.NewClient(&redis.Options{
		Addr:     connection.Addr + ":" + connection.Port,
		Username: connection.UserName,
		Password: connection.PassWord,
	})
	keyspace, err := rdb.Info(context.Background(), "keyspace").Result()
	if err != nil {
		return nil, err
	}

	m := make(map[string]int)
	v := strings.Split(keyspace, "\n")
	for i := 1; i < len(v)-1; i++ {
		database := strings.Split(v[i], ":")
		if len(database) < 2 {
			continue
		}
		vv := strings.Split(database[1], ",")
		if len(vv) < 1 {
			continue
		}
		keysnumber := strings.Split(vv[0], "=")
		if len(keysnumber) < 2 {
			continue
		}
		number, err := strconv.Atoi(keysnumber[1])
		if err != nil {
			return nil, err
		}
		m[database[0]] = number
	}
	//获取数据库的个数
	databasesRes, err := rdb.ConfigGet(context.Background(), "databases").Result()
	if err != nil {
		return nil, err
	}
	if len(databasesRes) < 2 {
		return nil, errors.New("连接数据异常")
	}
	dbnumber, err := strconv.Atoi(databasesRes[1].(string))
	if err != nil {
		return nil, err
	}
	data := make([]*define.DbItem, dbnumber)

	for i := range data {
		data[i] = &define.DbItem{
			Key:    "db" + strconv.Itoa(i),
			Number: 0,
		}
		if n, ok := m["db"+strconv.Itoa(i)]; ok {
			data[i].Number = n
		}
	}

	return data, nil
}

func KeyList(req *define.KeyListRequest) ([]string, error) {
	connection, err := helper.GetConnection(req.ConnIdentity)
	if err != nil {
		return nil, err
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     connection.Addr + ":" + connection.Port,
		Username: connection.Name,
		Password: connection.PassWord,
		DB:       req.Db,
	})
	count := 100
	if req.KeyWord != "" {
		count = 2000
	}
	result, _, err := rdb.Scan(context.Background(), 0, "*"+req.KeyWord+"*", int64(count)).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetKeyValue(req *define.KeyValueRequest) (*define.KeyValueReply, error) {
	connection, err := helper.GetConnection(req.ConnIdentity)
	if err != nil {
		return nil, err
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     connection.Addr + ":" + connection.Port,
		Username: connection.Name,
		Password: connection.PassWord,
		DB:       req.Db,
	})
	_type, err := rdb.Type(context.Background(), req.Key).Result()
	if err != nil {
		return nil, err
	}
	result, err := rdb.Get(context.Background(), req.Key).Result()
	if err != nil {
		return nil, err
	}
	ttl, err := rdb.TTL(context.Background(), req.Key).Result()
	if err != nil {
		return nil, err
	}
	return &define.KeyValueReply{
		Value: result,
		TTL:   ttl,
		Type:  _type,
	}, nil
}

func DeleteKeyValue(req *define.KeyValueRequest) error {
	connection, err := helper.GetConnection(req.ConnIdentity)
	if err != nil {
		return err
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     connection.Addr + ":" + connection.Port,
		Username: connection.Name,
		Password: connection.PassWord,
		DB:       req.Db,
	})
	_, err = rdb.Del(context.Background(), req.Key).Result()
	if err != nil {
		return err
	}

	return nil
}

func CreateKeyValue(req *define.CreateKeyValueRequest) error {
	connection, err := helper.GetConnection(req.ConnIdentity)
	if err != nil {
		return err
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     connection.Addr + ":" + connection.Port,
		Username: connection.Name,
		Password: connection.PassWord,
		DB:       req.Db,
	})
	if req.KValue == "" {
		req.KValue = "value"
	}
	_, err = rdb.Set(context.Background(), req.Key, req.KValue, -1).Result()
	return err
}

func UpdateKeyValue(req *define.UpdateKeyValueRequest) error {
	connection, err := helper.GetConnection(req.ConnIdentity)
	if err != nil {
		return err
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     connection.Addr + ":" + connection.Port,
		Username: connection.Name,
		Password: connection.PassWord,
		DB:       req.Db,
	})
	if req.TTL == 0 {
		req.TTL = -1
	}
	err = rdb.Set(context.Background(), req.Key, req.Value, req.TTL).Err()
	if err != nil {
		return err
	}
	return nil
}
