package eltype

// ConfigType 配置类型
type ConfigType int

const (
	// ConfigTypeGlobal 全局生效的配置类型
	ConfigTypeGlobal ConfigType = iota
	// ConfigTypeFriend 仅对好友消息生效的配置类型
	ConfigTypeFriend
	// ConfigTypeGroup 仅对群消息生效的配置类型
	ConfigTypeGroup
	//ConfigTypeCrontab 定时执行的配置
	ConfigTypeCrontab
	// ConfigTypeCounter 词频统计配置
	ConfigTypeCounter
)

// Config 配置对象
// @property	Type					ConfigType	配置类型
// @property	WhenMessageList			[]Message	作为触发条件的「消息」
// @property	WhenOperationList		[]Operation	作为触发条件的「事件/操作」
// @property	DoMessageList			[]Message	作为动作的「消息」
// @property	DoOperationList			[]Operation 作为动作的「操作」
// @property	Cron					string 		Cron 字符串
// @property	SenderList				[]Sender 	发送者列表
// @property	Receiver				[]Receiver 	接收者列表
type Config struct {
	Type              ConfigType
	IsCount           bool
	CountID           string
	Cron              string
	WhenMessageList   []Message
	WhenOperationList []Operation
	DoMessageList     []Message
	DoOperationList   []Operation
	SenderList        []Sender
	Receiver          []Receiver
}

// NewConfig 构造一个新的配置对象
// @param	configType				ConfigType	配置类型
// @param	WhenMessageList			[]Message	作为触发条件的「消息」
// @param	WhenOperationList		[]Operation	作为触发条件的「事件/操作」
// @param	DoMessageList			[]Message	作为动作的「消息」
// @param	DoOperationList			[]Operation 作为动作的「操作」
// @param	Cron					string 		Cron 字符串
func NewConfig(configType ConfigType,
	whenMessageList []Message,
	whenOperationList []Operation,
	doMessageList []Message,
	doOperationList []Operation,
	SenderList []Sender,
	Receiver []Receiver,
	Cron string,
	isCount bool,
	countID string) (Config, error) {
	var config Config
	config.Type = configType
	config.WhenMessageList = whenMessageList
	config.WhenOperationList = whenOperationList
	config.DoMessageList = doMessageList
	config.DoOperationList = doOperationList
	config.SenderList = SenderList
	config.Receiver = Receiver
	config.Cron = Cron
	config.IsCount = isCount
	config.CountID = countID
	return config, nil
}
