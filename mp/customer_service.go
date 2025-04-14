package mp

// KeFuInfo 客服基本信息
type KeFuInfo struct {
	KfAccount     string `json:"kf_account"`         // 完整客服帐号，格式为：帐号前缀@公众号微信号
	KfNick        string `json:"kf_nick"`            // 客服昵称
	KfID          int    `json:"kf_id"`              // 客服编号
	KfHeadImgURL  string `json:"kf_headimgurl"`      // 客服头像
	KfWX          string `json:"kf_wx"`              // 如果客服帐号已绑定了客服人员微信号， 则此处显示微信号
	InviteWX      string `json:"invite_wx"`          // 如果客服帐号尚未绑定微信号，但是已经发起了一个绑定邀请， 则此处显示绑定邀请的微信号
	InviteExpTime int    `json:"invite_expire_time"` // 如果客服帐号尚未绑定微信号，但是已经发起过一个绑定邀请， 邀请的过期时间，为unix 时间戳
	InviteStatus  string `json:"invite_status"`      // 邀请的状态，有等待确认"waiting,rejected,expired
}

// KeFuInfoListRes 客服信息列表
type KeFuInfoListRes struct {
	WXError
	KfList []*KeFuInfo `json:"kf_list"` // 完整客服信息
}

// KeFuOnlineInfo 客服在线信息
type KeFuOnlineInfo struct {
	KfAccount    string `json:"kf_account"`
	Status       int    `json:"status"`
	KfID         int    `json:"kf_id"`
	AcceptedCase int    `json:"accepted_case"`
}

type KeFuOnlineListRes struct {
	WXError
	KfOnlineList []*KeFuOnlineInfo `json:"kf_online_list"`
}

// CreateSessionReq 创建会话
type CreateSessionReq struct {
	OpenID    string `json:"openid" binding:"required" msg:"openid required"`
	KfAccount string `json:"kf_account" binding:"required" msg:"kf_account required"`
}

// AddKfAccountReq 添加客服帐号
type AddKfAccountReq struct {
	KfAccount string `json:"kf_account" binding:"required" msg:"kf_account required"`
	Nickname  string `json:"nickname" binding:"required" msg:"nickname required"`
	Password  string `json:"password" binding:"required" msg:"password required"`
}

// UpdateKfStatusReq 下发客服输入状态
type UpdateKfStatusReq struct {
	ToUser  string `json:"touser" binding:"required" msg:"touser required"`
	Command string `json:"command" binding:"required" msg:"command required"`
}

// InviteWorkerReq 邀请绑定客服帐号
type InviteWorkerReq struct {
	KfAccount string   `json:"kf_account" binding:"required" msg:"kf_account required"`
	InviteWX  []string `json:"invite_wx" binding:"required" msg:"invite_wx required"`
}

// GetMsgRecordReq 获取客服消息记录
type GetMsgRecordReq struct {
	StartTime int64 `json:"starttime" binding:"required" msg:"starttime required"`
	EndTime   int64 `json:"endtime" binding:"required" msg:"endtime required"`
	MsgId     int64 `json:"msgid" binding:"required" msg:"msgid required"`   //消息id顺序从小到大，从1开始
	Number    int64 `json:"number" binding:"required" msg:"number required"` //每次获取条数，最多10000条
}

// GetKFMsgListRes 获取客服消息记录返回结果
type GetKFMsgListRes struct {
	WXError
	Number     int64         `json:"number"` //每次获取条数，最多10000条
	MsgId      int64         `json:"msgid"`  //消息id顺序从小到大，从1开始
	RecordList []KFMsgRecord `json:"recordlist"`
}

// KFMsgRecord 客服消息记录
type KFMsgRecord struct {
	Worker string `json:"worker"`   // 完整客服账号，格式为：账号前缀@公众号微信号
	OpenID string `json:"openid"`   // 用户标识
	Time   int64  `json:"time"`     // 操作时间，unix时间戳
	Text   string `json:"text"`     // 聊天记录
	OpCode int64  `json:"opercode"` // 操作码，2002（客服发送信息），2003（客服接收消息）
}

// GetKFSessionStatusRes 获取会话状态返回结果
type GetKFSessionStatusRes struct {
	WXError
	KfAccount  string `json:"kf_account"`
	CreateTime int64  `json:"create_time"`
}

// GetKFSessionListRes 获取客服会话列表返回结果
type GetKFSessionListRes struct {
	WXError
	SessionList []KFSession `json:"sessionlist"`
}

// KFSession 客服会话
type KFSession struct {
	OpenID     string `json:"openid"`
	CreateTime int64  `json:"createtime"`
}

// GetUnacceptedSessionListRes 获取未接入会话列表返回结果
type GetUnacceptedSessionListRes struct {
	WXError
	Count        int64 `json:"count"` // 未接入会话数量
	WaitCaseList []struct {
		OpenID     string `json:"openid"`      // 粉丝的openid
		LatestTime int64  `json:"latest_time"` // 粉丝的最后一条消息的时间，unix时间戳
	} `json:"waitcaselist"` // 未接入会话列表，最多返回100条数据，按照来访顺序
}

//================================================ 客服发送消息 =======================

// KfMessageComm 客服发送消息公共参数
type KfMessageComm struct {
	ToUser  string `json:"touser" binding:"required" msg:"touser required"`
	MsgType string `json:"msgtype" binding:"required" msg:"msgtype required"`

	// KF 客服帐号:如果要指定发送的客服账号，则必须填写此字段
	KF struct {
		KFAccount string `json:"kf_account" binding:"required" msg:"kf_account required"`
	} `json:"customservice,omitempty"`
}

// KfTextMessage 文本消息
type KfTextMessage struct {
	KfMessageComm
	Text struct {
		Content string `json:"content" binding:"required" msg:"content required"`
	} `json:"text" binding:"required" msg:"text required"` // 发送文本消息时，支持插入跳小程序的文字链
}

// KfImageMessage 图片消息
type KfImageMessage struct {
	KfMessageComm
	Image struct {
		MediaID string `json:"media_id" binding:"required" msg:"media_id required"`
	} `json:"image" binding:"required" msg:"image required"`
}

// KfVoiceMessage 语音消息
type KfVoiceMessage struct {
	KfMessageComm
	Voice struct {
		MediaID string `json:"media_id" binding:"required" msg:"media_id required"`
	} `json:"voice" binding:"required" msg:"voice required"`
}

// KfVideoMessage 视频消息
type KfVideoMessage struct {
	KfMessageComm
	Video struct {
		MediaID      string `json:"media_id" binding:"required" msg:"media_id required"`
		ThumbMediaID string `json:"thumb_media_id" binding:"required" msg:"thumb_media_id required"`
		Title        string `json:"title"`
		Description  string `json:"description"`
	} `json:"video" binding:"required" msg:"video required"`
}

// KfMusicMessage 音乐消息
type KfMusicMessage struct {
	KfMessageComm
	Music struct {
		Title        string `json:"title"`
		Description  string `json:"description"`
		MusicURL     string `json:"musicurl" binding:"required" msg:"musicurl required"`
		HQMusicURL   string `json:"hqmusicurl" binding:"required" msg:"hqmusicurl required"`
		ThumbMediaID string `json:"thumb_media_id" binding:"required" msg:"thumb_media_id required"`
	} `json:"music" binding:"required" msg:"music required"`
}

// KfNewsLinkToURLMessage 图文消息，点击跳转到外链
type KfNewsLinkToURLMessage struct {
	KfMessageComm
	News struct { // 图文消息，限制在1条以内，注意，如果图文数超过1，则将会返回错误码45008。
		Articles []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			URL         string `json:"url"`
			PicURL      string `json:"picurl"`
		} `json:"articles"`
	} `json:"news" binding:"required" msg:"news required"`
}

// KfNewsLinkToPageMessage 图文消息，点击跳转到图文消息页面
//
//  草稿接口灰度完成后，将不再支持此前客服接口中带 media_id 的 mpnews 类型的图文消息
type KfNewsLinkToPageMessage struct {
	KfMessageComm
	MPNews struct {
		MediaID string `json:"media_id" binding:"required" msg:"media_id required"`
	} `json:"mpnews" binding:"required" msg:"mpnews required"`
}

// KfNewsLinkToArticleMessage 图文消息，点击跳转到图文消息页面
type KfNewsLinkToArticleMessage struct {
	KfMessageComm
	MPNewsArticle struct {
		ArticleId string `json:"article_id" binding:"required" msg:"article_id required"`
	} `json:"mpnewsarticle" binding:"required" msg:"mpnewsarticle required"`
}

// KfMenuMessage 菜单消息
//
// 当用户点击后，微信会发送一条XML消息到开发者服务器，格式如下：
//
// 	<xml>
// 	<ToUserName><![CDATA[ToUser]]></ToUserName>
// 	<FromUserName><![CDATA[FromUser]]></FromUserName>
// 	<CreateTime>1500000000</CreateTime>
// 	<MsgType><![CDATA[text]]></MsgType>
// 	<Content><![CDATA[满意]]></Content>
// 	<MsgId>1234567890123456</MsgId>
// 	<bizmsgmenuid>101</bizmsgmenuid>
// 	</xml>
type KfMenuMessage struct {
	KfMessageComm
	MsgMenu struct {
		HeadContent string `json:"head_content" binding:"required" msg:"head_content required"`
		List        []struct {
			Id      string `json:"id" binding:"required" msg:"id required"`
			Content string `json:"content" binding:"required" msg:"content required"`
		} `json:"list" binding:"required" msg:"list required"`
		TailContent string `json:"tail_content"`
	} `json:"msgmenu" binding:"required" msg:"msgmenu required"`
}

// KfCardMessage 卡券消息
//
//  特别注意客服消息接口投放卡券仅支持非自定义Code码和导入code模式的卡券的卡券，详情请见：创建卡券。
type KfCardMessage struct {
	KfMessageComm
	WXCard struct {
		CardId string `json:"card_id" binding:"required" msg:"card_id required"`
	} `json:"wxcard" binding:"required" msg:"wxcard required"`
}

// KfMiniProgramMessage 小程序卡片消息
type KfMiniProgramMessage struct {
	KfMessageComm
	MiniProgramPage struct {
		AppID        string `json:"appid" binding:"required" msg:"appid required"`
		PagePath     string `json:"pagepath" binding:"required" msg:"pagepath required"`
		Title        string `json:"title"`
		ThumbMediaID string `json:"thumb_media_id" binding:"required" msg:"thumb_media_id required"`
	} `json:"miniprogrampage" binding:"required" msg:"miniprogrampage required"`
}
