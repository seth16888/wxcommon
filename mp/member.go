package mp

// MemberInfo 粉丝信息
type MemberInfo struct {
	IsSubscriber int    `json:"subscribe"` // 用户是否订阅该公众号标识, 值为0时, 代表此用户没有关注该公众号, 拉取不到其余信息
	OpenId       string `json:"openid"`    // 用户的标识, 对当前公众号唯一
	Nickname     string `json:"nickname"`  // 用户的昵称
	Sex          int    `json:"sex"`       // 用户的性别, 值为1时是男性, 值为2时是女性, 值为0时是未知
	Language     string `json:"language"`  // 用户的语言, zh_CN, zh_TW, en
	City         string `json:"city"`      // 用户所在城市
	Province     string `json:"province"`  // 用户所在省份
	Country      string `json:"country"`   // 用户所在国家

	// 用户头像, 最后一个数值代表正方形头像大小(有0, 46, 64, 96, 132数值可选, 0代表640*640正方形头像), 用户没有头像时该项为空
	HeadImageURL string `json:"headimgurl"`

	SubscribeTime int64  `json:"subscribe_time"`    // 用户关注时间, 为时间戳. 如果用户曾多次关注, 则取最后关注时间
	UnionId       string `json:"unionid,omitempty"` // 只有在用户将公众号绑定到微信开放平台帐号后, 才会出现该字段.
	Remark        string `json:"remark"`            // 公众号运营者对粉丝的备注, 公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
	GroupId       int64  `json:"groupid"`           // 用户所在的分组ID

	TagIdList      []int  `json:"tagid_list"`      // Tag List
	SubscribeScene string `json:"subscribe_scene"` // 返回用户关注的渠道来源
	QrScene        int    `json:"qr_scene"`        // 二维码扫码场景（开发者自定义）场景值ID，临时二维码时为32位非0整型，永久二维码时最大值为100000（目前参数只支持1--100000）
	QrSceneStr     string `json:"qr_scene_str"`    // 二维码扫码场景描述（开发者自定义）场景值ID（字符串形式的ID），字符串类型，长度限制为1到64
}

// GetMemberListRes 从WX获取用户openid列表返回结果
type GetMemberListRes struct {
	Total int `json:"total"`
	Count int `json:"count"`
	Data  struct {
		OpenIdList []string `json:"openid"`
	} `json:"data"`
	NextOpenId string `json:"next_openid"`
}

// Tag
type Tag struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

// TagMembersRes 从WX获取标签下的用户列表返回结果
type TagMembersRes struct {
	Count int `json:"count"`
	Data  struct {
		OpenIdList []string `json:"openid"`
	} `json:"data"`
	NextOpenId string `json:"next_openid"`
}

type BatchGetMemberInfoReq struct {
	AccessToken string              `json:"access_token" form:"access_token" binding:"required" msg:"access_token required"`
	UserList    []MemberInfoReqItem `json:"user_list" binding:"required,dive" msg:"user_list is required"`
}

type MemberInfoReqItem struct {
	OpenId string `json:"openid" binding:"required"`
	Lang   string `json:"lang"`
}
