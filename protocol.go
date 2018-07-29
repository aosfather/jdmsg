package main


/**
  jd aplpha 协议

{{
     "request": {
    "dialogState": "STARTED",
    "requestId": "ec1b74c1-7f5a-429c-be0f-1c6d329b45c0_812_test",
    "timestamp": 1532826391222,
    "type": "LaunchRequest"
  },
  "session": {
    "application": {
      "applicationId": "jd.alpha.skill.d5ec16545aed4eba9684d416cbeb1249"
    },
    "device": {
      "deviceId": "deviceId_cfd4caab4e8747c49728a5f7877dbe12_dev_test"
    },
    "isNew": true,
    "sessionId": "75b5f8e4-205f-4ec4-9b48-b787e43dbfc7",
    "user": {
      "userId": "jd.account.2a2fc6110554047c86a390ba3e838e67"
    }
  },
  "version": "1.0"
}
 */

 type JDMessage struct {
    Request MsgRequest `json:"request"`
    Session MsgSession `json:"session"`
 	Version string `json:"version"`
 	
 }

 type MsgRequest struct {
 	Id string `json:"requestId"`
 	Type string `json:"type"`
 	TimeStamp int64 `json:"timestamp"`
 	State string `json:"dialogState"`
 	Intent _intent `json:"intent"`
 	Reason _reason `json:"reason"`
 }

 type MsgSession struct {
 	Application _application `json:"application"`
 	Device _device `json:"device"`
 	Contexts _contexts `json:"contexts"`
 	New bool `json:"isNew"`
 	Id string `json:"sessionId"`
 }

 type _application struct {
 	Id string `json:"applicationId"`
 }

 type _device struct {
 	Id string `json:"deviceId"`
 }

 type _contexts map[string]interface{}

 /**
 version	协议版本号，当前为“1.0”版本	string
intent	当前会话的意图名称。技能也可以改变意图名称，让新意图接管对话流程	string
contexts	技能存储会话上下文数据的空间，下次请求会带回当前响应存放的键值对数据，在不同会话中也是有效的。格式为 "contexts": {"string": " object "}	map <string, object>
response	回应信息	response object
directives	指令集合，可以存放多条指令。指令类型有以下几种：
1.对话指令
2.音频控制指令	directive array
shouldEndSession	是否结束会话，设置为false后，用户可以继续对话。	boolean
  */
 // 回应信息

 type JDMessageResponse struct {
   Intent string  `json:"intent"`
   Contexts _contexts `json:"contexts"`
   Directives []interface{} `json:"directives"`
   Response _response `json:"response"`
   Version string `json:"version"`
   ShouldEnd bool `json:"shouldEndSession"`
 }


 type _response struct {
    Output _response_output `json:"output"`
    Reprompt *_response_output `json:"reprompt"`
    Card _response_card `json:"card"`
 }

 type _response_output struct {
	 Type string `json:"type"`
	 Text string `json:"text"`
 }

 type _response_card struct {

 }


//意图
/**
intent object：

字段	说明	类型
name	意图名称	string
slots	槽位数组数据	map <string, slot object>
confirmResult	意图信息确认状态。可选项如下：
1.未确认
2.确认
3.否认
string：
1.NONE
2.CONFIRMED
3.DENIED
slot object：

字段	说明	类型
name	槽位名称	string
value	槽值	string
matched	是否匹配。可选项如下：
1.提取的槽值在槽位类型之中，匹配成功
2.提取的槽值不在槽位类型之中，匹配失败	boolean
confirmResult	槽位信息确认状态。可选项如下：
1.未确认
2.确认
3.否认	string：
1.NONE
2.CONFIRMED
3.DENIED
 */
 
 type _intent struct {
 	Name string
 	Slots map[string]_slot
 	ConfirmResult string `json:"confirmResult"`
 	
 }
 
 type _slot struct {
 	Name string `json:"name"`
 	Value string `json:"value"`
 	Matched bool `json:"matched"`
 	ConfirmResult string `json:"confirmResult"`
 }


 const (
 	//dialog state 对话状态,如果使用了对话指令,会存在对话数据以及状态信息
 	DS_STARTED="STARTED" //开始
 	DS_PROGRESS="IN_PROGRESS" //进行中
 	DS_COMPLETED="COMPLETED"//完成

 	//confirmResult 意图信息确认状态
 	CR_NONE="NONE"   //未确认
 	CR_CONFIRMED="CONFIRMED"//已确认
 	CR_DENIED="DENIED"//否认


 	//错误类型
    RT_NORMAL ="NOMAL"
    RT_ERROR="ERROR"

 )

 //reason
 type _reason struct {
 	Type string `json:"type"`
 	Message string `json:"message"`
 }