package models

// 省市区
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
	Adcode string `json:"adcode"` // 高德code
}

func (t *Area) TableName() string {
	 return "lc_area"
}