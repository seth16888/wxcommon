package mp

// AddTemplateReq 选择订阅通知模板
type AddTemplateReq struct {
	Tid string `json:"tid" binding:"required"` //模板标题 id
	// 开发者自行组合好的模板关键词列表，关键词顺序可以自由搭配（例如 [3,5,4] 或 [4,5,3]），最多支持5个，最少2个关键词组合
	KidList []int `json:"kidList" binding:"required"` //关键词 id 列表

	SceneDesc string `json:"sceneDesc" binding:"required"` //服务场景描述，15个字以内
}

// GetSubscriptionCategoryRes 获取订阅通知公众号类目
type GetSubscriptionCategoryRes struct {
	WXError
	Data []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
}

// GetPubTemplateKeyWordsRes 获取模板中的关键词
type GetPubTemplateKeyWordsRes struct {
	WXError
	Count int `json:"count"`
	Data  []struct {
		Kid     int    `json:"kid"`
		Name    string `json:"name"`
		Rule    string `json:"rule"`
		Example string `json:"example"`
	} `json:"data"`
}

// GetPubTemplateTitlesRes 获取类目下的公共模板
type GetPubTemplateTitlesRes struct {
	WXError
	Count int `json:"count"`
	Data  []struct {
		Tid        string `json:"tid"`
		Title      string `json:"title"`
		Type       int    `json:"type"`
		CategoryId string `json:"categoryId"`
	} `json:"data"`
}

// GetPrivateTemplateListRes 获取当前帐号下的个人模板列表
type GetPrivateTemplateListRes struct {
	WXError
	Data []struct {
		PriTmplId string `json:"priTmplId"`
		Title     string `json:"title"`
		Content   string `json:"content"`
		Example   string `json:"example"`
		Type      int    `json:"type"`
	} `json:"data"`
}

// SendSubscriptionMessageReq 发送订阅消息
type SendSubscriptionMessageReq struct {
	Touser      string                                   `json:"touser" binding:"required"`      // 接收者（用户）的 openid
	TemplateId  string                                   `json:"template_id" binding:"required"` // 所需下发的订阅模板id
	Page        string                                   `json:"page,omitempty"`                 // 点击模板卡片后的跳转页面
	Data        map[string]*SendSubscribeMessageDataItem `json:"data" binding:"required"`        // 模板内容，格式形如 { "key1": { "value": any }, "key2": { "
	MiniProgram struct {
		AppID    string `json:"appid"`    // 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系）
		PagePath string `json:"pagepath"` // 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar）
	} `json:"miniprogram,omitempty"` // 可选,跳转至小程序地址
}

type SendSubscribeMessageDataItem struct {
	Value string `json:"value"`
}
