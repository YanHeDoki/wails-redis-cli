package main

import (
	"changeme/internal/define"
	"changeme/internal/service"
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// Connection 连接列表
func (a *App) ConnectionList() M {
	list, err := service.ConnectionList()
	if err != nil {
		return map[string]interface{}{
			"code": -1,
			"msg":  "Error" + err.Error(),
		}
	}
	return map[string]interface{}{
		"code": 200,
		"data": list.Connections,
	}
}

func (a *App) ConnectionCreate(connection *define.Connection) any {
	err := service.ConnectionCreate(connection)
	if err != nil {
		return M{"code": -1, "msg": "ERROR" + err.Error()}
	}
	return M{"code": 200, "msg": "新建成功"}
}

func (a *App) ConnectionEdit(connection *define.Connection) any {
	err := service.ConnectionEdit(connection)
	if err != nil {
		return M{"code": -1, "msg": "ERROR" + err.Error()}
	}
	return M{"code": 200, "msg": "编辑成功"}
}

func (a *App) ConnectionDelete(identity string) any {
	err := service.ConnectionDelete(identity)
	if err != nil {
		return M{"code": -1, "msg": "ERROR" + err.Error()}
	}
	return M{"code": 200, "msg": "删除成功成功"}
}

func (a *App) DbList(identity string) M {
	dbs, err := service.DbList(identity)
	if err != nil {
		return map[string]interface{}{
			"code": -1,
			"msg":  "Error" + err.Error(),
		}
	}
	return map[string]interface{}{
		"code": 200,
		"data": dbs,
	}
}

func (a *App) KeyList(req *define.KeyListRequest) M {
	if req.ConnIdentity == "" {
		return M{"code": -1, "msg": "连接的唯一标识不能为空"}
	}
	data, err := service.KeyList(req)
	if err != nil {
		return map[string]interface{}{
			"code": -1,
			"msg":  "Error" + err.Error(),
		}
	}
	return map[string]interface{}{
		"code": 200,
		"data": data,
	}
}

func (a *App) GetKeyValue(req *define.KeyValueRequest) M {
	if req.ConnIdentity == "" || req.Key == "" {
		return M{"code": -1, "msg": "连接的唯一标识不能为空"}
	}
	data, err := service.GetKeyValue(req)
	if err != nil {
		return map[string]interface{}{
			"code": -1,
			"msg":  "Error" + err.Error(),
		}
	}
	return map[string]interface{}{
		"code": 200,
		"data": data,
	}
}

func (a *App) DeleteKeyValue(req *define.KeyValueRequest) M {
	if req.ConnIdentity == "" || req.Key == "" {
		return M{"code": -1, "msg": "连接的唯一标识不能为空"}
	}
	err := service.DeleteKeyValue(req)
	if err != nil {
		return map[string]interface{}{
			"code": -1,
			"msg":  "Error" + err.Error(),
		}
	}
	return map[string]interface{}{
		"code": 200,
		"msg":  "删除成功",
	}
}

func (a *App) CreateKeyValue(req *define.CreateKeyValueRequest) M {
	if req.ConnIdentity == "" || req.Key == "" || req.Type == "" {
		return M{"code": -1, "msg": "连接的唯一标识或者参数不能为空"}
	}
	err := service.CreateKeyValue(req)
	if err != nil {
		return map[string]interface{}{
			"code": -1,
			"msg":  "Error" + err.Error(),
		}
	}
	return map[string]interface{}{
		"code": 200,
		"msg":  "新增成功",
	}
}

func (a *App) UpdataKeyValue(req *define.UpdateKeyValueRequest) M {
	if req.ConnIdentity == "" || req.Key == "" || req.Value == "" {
		return M{"code": -1, "msg": "连接的唯一标识或者参数不能为空"}
	}
	err := service.UpdateKeyValue(req)
	if err != nil {
		return map[string]interface{}{
			"code": -1,
			"msg":  "Error" + err.Error(),
		}
	}
	return map[string]interface{}{
		"code": 200,
		"msg":  "更新成功",
	}
}
