package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

/**
  技能
   1、第一次使用
   2、获取留言
   3、发送留言
*/
type Msg struct {
	From string `json:"from"` //发送者
	To   string `json:"to"`   //接收者
	Data string `json:"data"` //信息
}

type SkillService struct {
	aliasPath string //存放别名
	msgPath   string //存放消息的目录
}

func (this *SkillService) getAliasFileName(alias string) string {
	return fmt.Sprintf("%s/%s.alias", this.aliasPath, alias)
}

func (this *SkillService) getMsgDirectory(deviceId string) string {
	return fmt.Sprint("%s/%s/", this.msgPath, deviceId)
}

//通过别名获取设备id
func (this *SkillService) GetDeviceIdByAlias(alias string) string {
	fp := this.getAliasFileName(alias)
	if PathExists(fp) {
		data, _ := ioutil.ReadFile(fp)
		return string(data)
	}

	return ""
}

//创建别名，如果别名存在则创建失败
func (this *SkillService) CreateAlias(alias string, deviceId string) bool {
	fp := this.getAliasFileName(alias)
	if !PathExists(fp) {
		ioutil.WriteFile(fp, []byte(deviceId), os.ModeAppend)
		return true
	}
	return false
}

//根据设备Id获取留言
func (this *SkillService) GetMessages(deviceId string) []Msg {
	pathname := this.getMsgDirectory(deviceId)
	if PathExists(pathname) {
		dir_list, e := ioutil.ReadDir(pathname)
		if e != nil {
			fmt.Println("read dir error")
			return nil
		}
		for i, v := range dir_list {
			fmt.Println(i, "=", v.Name())
		}
	}
	return nil
}

func (this *SkillService) SendMessage(fromDeviceId string, toAlias string, msg string) {
	deviceId := this.GetDeviceIdByAlias(toAlias)
	if deviceId != "" {
		pathname := this.getMsgDirectory(deviceId)
		ioutil.WriteFile(pathname+fromDeviceId+".msg", []byte(msg), os.ModeAppend)

	}
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
