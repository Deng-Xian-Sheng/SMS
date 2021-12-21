package controller

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

func Pretreatment(DatabaseInfo map[string]string) (*casbin.Enforcer, error) {
	dsn := DatabaseInfo["UserName"] + ":" + DatabaseInfo["PassWd"] + "@tcp(" + DatabaseInfo["IP"] + ":" + DatabaseInfo["Port"] + ")/" + DatabaseInfo["NameDB"]
	a, err := gormadapter.NewAdapter("mysql", dsn, true)
	if err != nil {
		return nil, err
	}
	e, err := casbin.NewEnforcer("./controller/rbac_model.conf", a)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func Controller(DatabaseInfo map[string]string, SOA map[string]string) (bool, error) {
	e, err := Pretreatment(DatabaseInfo)
	if err != nil {
		return false, err
	}
	// sub 想要访问资源的用户。
	// obj 将被访问的资源。
	// act 用户对资源执行的操作。
	ok, err := e.Enforce(SOA["sub"], SOA["obj"], SOA["act"])
	if err != nil {
		// 处理err
		return false, err
	}
	if ok == true {
		// 允许alice读取data1
		return true, nil
	} else {
		// 拒绝请求，抛出异常
		return false, nil
	}
}

func AddRule(DatabaseInfo map[string]string, SOA map[string]string) (bool, error) {
	e, err := Pretreatment(DatabaseInfo)
	if err != nil {
		return false, err
	}
	//添加
	e.AddPolicy(SOA["sub"], SOA["obj"], SOA["act"])
	//验证
	hasPolicy := e.HasPolicy(SOA["sub"], SOA["obj"], SOA["act"])
	// 若显示为 true，也就是我们成功添加了策略
	return hasPolicy, nil
}

func DeleteRule(DatabaseInfo map[string]string, SOA map[string]string) (bool, error) {
	e, err := Pretreatment(DatabaseInfo)
	if err != nil {
		return false, err
	}
	// 移除一条策略，然后使用 HasPolicy() 来确认
	e.RemovePolicy(SOA["sub"], SOA["obj"], SOA["act"])
	hasPolicy := e.HasPolicy(SOA["sub"], SOA["obj"], SOA["act"])
	// 若显示为 false，也就是我们成功删除了策略
	return hasPolicy, nil
}
