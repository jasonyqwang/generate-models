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
