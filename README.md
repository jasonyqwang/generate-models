# generate models
#### GO语言连接Mysql生成对应的model，包括对应字段类型、注释等。生成基础的结构体，不局限于某一个ORM。

 **生成示例---------**

```go 
package model

type Area struct {
	Id int `json:"id"` // ID
	Pid int `json:"pid"` // 父id
	Shortname string `json:"shortname"` // 简称
	Name string `json:"name"` // 名称
	MergerName string `json:"merger_name"` // 全称
	Level int `json:"level"` // 层级 0 1 2 省市区县
	Pinyin string `json:"pinyin"` // 拼音
	Code string `json:"code"` // 长途区号
	ZipCode string `json:"zip_code"` // 邮编
	First string `json:"first"` // 首字母
	Lng string `json:"lng"` // 经度
	Lat string `json:"lat"` // 纬度
}

func (Area) TableName() string {
	 return "lcc_area"
}
```

**参数配置--------conf.go**

```go 
package conf

//model生成的包名
const ModelPackageName = "model"
// model保存路径
const ModelPath = "./"+ModelPackageName+"/"
// 是否覆盖已存在model
const ModelReplace = true
// 数据库驱动
const DriverName = "mysql"

type DbConf struct {
	//HOST
	Host   string
	//端口
	Port   string
	//用户
	User   string
	//密码
	Pwd    string
	//数据库名称
	DbName string
	//数据库表的前缀
	TablePrefix string
}
// 数据库链接配置
var MasterDbConfig DbConf = DbConf{
	Host:   "",
	Port:   "3306",
	User:   "",
	Pwd:    "",
	DbName: "",
	TablePrefix: "",
}
```

**生成model--------**
```go
package main

import (
	"modeltools/dbtools"
	"modeltools/generate"
)


func main() {
	//初始化数据库
	dbtools.Init()
	generate.Genertate() //生成所有表信息
	//generate.Genertate("lcc_area") //生成指定表信息，可变参数可传入多个表名
}


```

## 参考链接
-- https://github.com/longzongqin/modeltools