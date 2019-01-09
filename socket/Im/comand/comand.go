package comand

type Parse interface {
	Parse() error
}

//发送接口
type Send interface {
	Send() error
}

//广播
type Broad interface {
	Broad() error
}

//退出
type Exit interface {
	Exit() error
}

//登录
type Login interface {
	Login() error
}

//消息+签名的方式验证
type Message struct {
	Cmd     string `json:"cmd"`     //命令
	Name    string `json:"name"`    //用户名
	Context string `json:"content"` //内容区
}
