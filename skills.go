package main
/**
  技能
   1、第一次使用
   2、获取留言
   3、发送留言
 */
type Msg struct {
	From string //发送者
	To string  //接收者
	Data string //信息
}

 type SkillService struct {
 	aliasPath string //存放别名
 	msgPath string  //存放消息的目录
 }

//通过别名获取设备id
 func (this *SkillService)GetDeviceIdByAlias(alias string) string {


 	return ""
 }

//创建别名，如果别名存在则创建失败
 func (this *SkillService)CreateAlias(alias string,deviceId string) bool {

 	return false
 }

