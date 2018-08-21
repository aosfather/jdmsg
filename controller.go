package main

import (
	"fmt"
	"github.com/aosfather/bingo/mvc"
)

/**


 */

type AlphaController struct {
	mvc.SimpleController
	Skill *SkillService `Inject:""`
}

func (this *AlphaController) GetUrl() string {
	return "/jd/msg"
}

func (this *AlphaController) GetParameType(method string) interface{} {
	return &JDMessage{}

}

func (this *AlphaController) Post(c mvc.Context, p interface{}) (interface{}, mvc.BingoError) {

	if value, ok := p.(*JDMessage); ok {
		fmt.Println(value)

		res := JDMessageResponse{}
		res.Version = "1.0"
		res.ShouldEnd = true
		res.Response.Output.Type = "PlainText"
		if value.Request.Type == "LaunchRequest" {

			res.Response.Output.Text = "欢迎使用我的留言！在使用的时候需要先绑定一个手机号，例如你可以说绑定13600，来绑定13600的手机号，这样其它人就可以通过13600来给你留言了；可以对我说获取留言，来查收其它给你的留言。最后你可以对我说给手机号说你想说的话，来给指定的手机号留言，例如给13600说你好棒，这样13600的用户就可以收到你好棒的留言了。玩的愉快！"

		} else if value.Request.Type == "IntentRequest" {
			name := value.Request.Intent.Name
			switch name {
			//绑定手机
			case "bind_mobile":
				res.Response.Output.Text = this.bindMobile(value.Session.Device.Id, value.Request.Intent)
			case "get_msg":
				res.Response.Output.Text = this.getMessages(value.Session.Device.Id)
			case "send_msg":
				res.Response.Output.Text = this.sendMessage(value.Session.Device.Id, value.Request.Intent)
			}

		}
		return &res, nil
	}
	return nil, nil
}

//绑定手机
func (this *AlphaController) bindMobile(deviId string, intent _intent) string {
	slot := intent.Slots["mobile"]
	if slot.ConfirmResult == "CONFIRMED" && slot.Matched {
		mobile := slot.Value
		if this.Skill.CreateAlias(mobile, deviId) {
			return "绑定成功"
		}

	}

	return "绑定失败，请重新绑定，并确认手机号输入正确，未被其它设备绑定！"

}

//获取留言
func (this *AlphaController) getMessages(deviId string) string {
	msgs := this.Skill.GetMessages(deviId)
	if len(msgs) > 0 {
		rmsg := fmt.Sprintf("您有%d条留言。", len(msgs))

		for index, msg := range msgs {
			rmsg += fmt.Sprintf("第%d条:%s对你说:%s。", index+1, msg.From, msg.Data)
		}

		return rmsg

	}

	return "您没有新的留言！"
}

//发送留言
func (this *AlphaController) sendMessage(devid string, intent _intent) string {
	target := intent.Slots["target"]
	msg := intent.Slots["values"]

	if this.Skill.GetDeviceIdByAlias(target.Value) != "" {

		this.Skill.SendMessage(devid, target.Value, msg.Value)

		return fmt.Sprintf("给%s留言成功！", target.Value)
	}

	return fmt.Sprintf("留言失败，%s不存在！", target.Value)
}
