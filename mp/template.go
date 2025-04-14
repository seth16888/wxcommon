package mp

// SetTemplateIndustryReq 设置所属行业
type SetTemplateIndustryReq struct {
	// 公众号模板消息所属行业编号
	IndustryId1 string `json:"industry_id1"`
	// 公众号模板消息所属行业编号
	IndustryId2 string `json:"industry_id2"`
}

// GetTemplateIndustryResp 获取设置的行业信息
type GetTemplateIndustryResp struct {
	PrimaryIndustry struct {
		FirstClass  string `json:"first_class"`
		SecondClass string `json:"second_class"`
	} `json:"primary_industry"`
	DeputyIndustry struct {
		FirstClass  string `json:"first_class"`
		SecondClass string `json:"second_class"`
	} `json:"secondary_industry"`
}

// GetTemplateIdReq 获取模板Id请求
type GetTemplateIdReq struct {
	TemplateIdShort string   `json:"template_id_short" binding:"required" msg:"template_id_short required"`
	KeywordNameList []string `json:"keyword_name_list" binding:"required,gt=0" msg:"keyword_name_list required"`
}

// GetTemplateIdRes 获取模板Id返回
type GetTemplateIdRes struct {
	WXError
	TemplateId string `json:"template_id"`
}

// GetAllPrivateTemplateRes 获取模板列表返回
type GetAllPrivateTemplateRes struct {
	TemplateList []struct {
		TemplateID      string `json:"template_id"`
		Title           string `json:"title"`
		Content         string `json:"content"`
		Example         string `json:"example"`
		PrimaryIndustry string `json:"primary_industry"`
		DeputyIndustry  string `json:"deputy_industry"`
	} `json:"template_list"`
}

// TemplateMessage 发送的模板消息内容
type TemplateMessage struct {
	ToUser      string                       `json:"touser"`                  // 必须, 接受者OpenID
	TemplateID  string                       `json:"template_id"`             // 必须, 模版ID
	URL         string                       `json:"url,omitempty"`           // 可选, 用户点击后跳转的URL, 该URL必须处于开发者在公众平台网站中设置的域中
	Color       string                       `json:"color,omitempty"`         // 可选, 整个消息的颜色, 可以不设置
	Data        map[string]*TemplateDataItem `json:"data"`                    // 必须, 模板数据
	ClientMsgID string                       `json:"client_msg_id,omitempty"` // 可选, 防重入ID

	MiniProgram struct {
		AppID    string `json:"appid"`    // 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系）
		PagePath string `json:"pagepath"` // 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar）
	} `json:"miniprogram,omitempty"` // 可选,跳转至小程序地址
}

// TemplateDataItem 模版内某个 .DATA 的值
type TemplateDataItem struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

// 发送订阅消息
type SubscribeMessage struct {
	TemplateMessage

	Scene string `json:"scene"`
	Title string `json:"title"`
}

// SendTemplateMessageRes 发送模版消息返回结果
type SendTemplateMessageRes struct {
	WXError

	MsgID int64 `json:"msgid"`
}

// GetBlockedMessagesReq 获取被屏蔽的消息请求
type GetBlockedMessagesReq struct {
	// 模板消息ID
	TmplMsgId string `json:"tmpl_msg_id" binding:"required" msg:"tmpl_msg_id required"`
	// 上一页查询结果最大的id，用于翻页，第一次传0
	LargestId int64 `json:"largest_id,omitempty"`
	// 单页查询的大小，最大100
	Limit int64 `json:"limit" binding:"required,gt=0,lte=100" msg:"limit required"`
}

// GetBlockedMessagesRes 获取被屏蔽的消息返回
type GetBlockedMessagesRes struct {
	WXError

	MsgInfo []BlockedMessageInfo `json:"msginfo"`
}

// BlockedMessageInfo 被屏蔽的消息
type BlockedMessageInfo struct {
	Id            int64  `json:"id"`
	TmplMsgId     string `json:"tmpl_msg_id"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	SendTimestamp int64  `json:"send_timestamp"`
	OpenId        string `json:"openid"`
}
