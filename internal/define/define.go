package define

import "time"

var ConfigName = "redis-client.conf"

type Connection struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
}

type Config struct {
	Connections []*Connection `json:"connections"`
}

type DbItem struct {
	Key    string `json:"key"`
	Number int    `json:"number"`
}

type KeyListRequest struct {
	ConnIdentity string `json:"conn_Identity"`
	Db           int    `json:"db"`
	KeyWord      string `json:"keyWord"`
}
type KeyValueRequest struct {
	ConnIdentity string `json:"conn_Identity"`
	Db           int    `json:"db"`
	Key          string `json:"key"`
}

type KeyValueReply struct {
	Value string        `json:"value"`
	TTL   time.Duration `json:"ttl"`
	Type  string        `json:"type"`
}

type CreateKeyValueRequest struct {
	ConnIdentity string `json:"conn_Identity"`
	Db           int    `json:"db"`
	Key          string `json:"key"`
	KValue       string `json:"Kvalue"`
	Type         string `json:"type"`
}

type UpdateKeyValueRequest struct {
	ConnIdentity string        `json:"conn_Identity"`
	Db           int           `json:"db"`
	Key          string        `json:"key"`
	TTL          time.Duration `json:"ttl"`
	Value        string        `json:"value"`
}
