package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
	return fmt.Sprintf("%s/%s/", this.msgPath, deviceId)
}

//通过别名获取设备id
func (this *SkillService) GetDeviceIdByAlias(alias string) string {
	fp := this.getAliasFileName(alias)
	if PathExists(fp) {
		fmt.Println("exist ", alias)
		data, err := ioutil.ReadFile(fp)
		if err != nil {
			fmt.Println("error:%s", err.Error())
		}
		fmt.Println(data)
		return string(data)
	}

	return ""
}

//创建别名，如果别名存在则创建失败
func (this *SkillService) CreateAlias(alias string, deviceId string) bool {
	fp := this.getAliasFileName(alias)
	if !PathExists(fp) {
		ioutil.WriteFile(fp, []byte(deviceId), 0666)
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
		result := []Msg{}
		for i, v := range dir_list {
			fmt.Println(i, "=", v.Name())
			msg := Msg{}
			content, e := ioutil.ReadFile(pathname + "/" + v.Name())
			if e != nil {
				fmt.Println(e.Error())
			}
			msg.Data = string(content)
			msg.From = strings.Replace(v.Name(), ".msg", "", 1)
			msg.To = deviceId
			result = append(result, msg)
		}

		return result
	}
	fmt.Println(deviceId, " msg not found!")
	return nil
}

func (this *SkillService) SendMessage(fromDeviceId string, toAlias string, msg string) {
	deviceId := this.GetDeviceIdByAlias(toAlias)
	if deviceId != "" {
		pathname := this.getMsgDirectory(deviceId)
		if !PathExists(pathname) {
			fmt.Println(os.Mkdir(pathname, os.ModePerm))
		}
		ioutil.WriteFile(pathname+fromDeviceId+".msg", []byte(msg), 0666)

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
