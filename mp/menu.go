/**
官网type:
view:  (url  )链接
text:  (value)返回文本
img:   (value)mediaID
photo:
video: (value)下载链接
voice: (value)mediaID
news:  (value)mediaID + 图文消息news_info => API下发卡片: article_id类型+article_id参数
*/

package mp

// CreateMenuReq 设置菜单请求数据
type CreateMenuReq struct {
	Button    []*Button  `json:"button,omitempty"`
	MatchRule *MatchRule `json:"matchrule,omitempty"`
}

// MenuRes 获取API设置的菜单返回数据
type MenuRes struct {
	WXError

	Menu struct {
		Button []Button `json:"button,omitempty"`
		MenuID int64    `json:"menuid,omitempty"`
	} `json:"menu,omitempty"`
	Conditionalmenu []conditionalMenuRes `json:"conditionalmenu,omitempty"`
}

// ConditionalMenu 个性化菜单返回结果
type conditionalMenuRes struct {
	Button    []Button  `json:"button"`
	MatchRule MatchRule `json:"matchrule"`
	MenuID    int64     `json:"menuid"`
}

// MatchRule 个性化菜单规则
type MatchRule struct {
	TagId              string `json:"tag_id,omitempty"`
	ClientPlatformType string `json:"client_platform_type,omitempty"`
}

// resMenuTryMatch 菜单匹配请求结果
type MenuTryMatchRes struct {
	WXError

	Button []*Button `json:"button,omitempty"`
}

// SelfMenuInfoRes 官网菜单配置详情
type SelfMenuInfoRes struct {
	WXError

	IsMenuOpen   int32 `json:"is_menu_open"`
	SelfMenuInfo struct {
		Button []SelfMenuButton `json:"button,omitempty"`
	} `json:"selfmenu_info,omitempty"`
}

// SelfMenuButton 官网菜单按钮配置
type SelfMenuButton struct {
	Type      string `json:"type,omitempty"`
	Name      string `json:"name"`
	Key       string `json:"key,omitempty"`
	URL       string `json:"url,omitempty"`
	Value     string `json:"value,omitempty"`
	SubButton struct {
		List []SelfMenuButton `json:"list,omitempty"`
	} `json:"sub_button,omitempty"`
	NewsInfo struct {
		List []ButtonNews `json:"list,omitempty"`
	} `json:"news_info,omitempty"`
}

// ButtonNews 图文消息菜单
type ButtonNews struct {
	Title      string `json:"title"`
	Author     string `json:"author"`
	Digest     string `json:"digest"`
	ShowCover  int32  `json:"show_cover"` //是否显示封面，0为不显示，1为显示
	CoverURL   string `json:"cover_url"`  //封面图片的URL
	ContentURL string `json:"content_url"` //正文的URL
	SourceURL  string `json:"source_url"` //原文的URL，若置空则无查看原文入口
}

// Button 菜单按钮
type Button struct {
	Type       string    `json:"type,omitempty"`
	Name       string    `json:"name,omitempty"`
	Key        string    `json:"key,omitempty"`
	URL        string    `json:"url,omitempty"`
	MediaID    string    `json:"media_id,omitempty"`
	AppID      string    `json:"appid,omitempty"`
	PagePath   string    `json:"pagepath,omitempty"`
  ArticleID  string    `json:"article_id,omitempty"`
	SubButtons []*Button `json:"sub_button,omitempty"`
}

// SetSubButton 设置二级菜单
func (btn *Button) SetSubButton(name string, subButtons []*Button) *Button {
	btn.Name = name
	btn.SubButtons = subButtons
	btn.Type = ""
	btn.Key = ""
	btn.URL = ""
	btn.MediaID = ""
	return btn
}

// SetClickButton btn 为click类型
func (btn *Button) SetClickButton(name, key string) *Button {
	btn.Type = "click"
	btn.Name = name
	btn.Key = key
	btn.URL = ""
	btn.MediaID = ""
	btn.SubButtons = nil
	return btn
}

// SetViewButton view类型
func (btn *Button) SetViewButton(name, url string) *Button {
	btn.Type = "view"
	btn.Name = name
	btn.URL = url
	btn.Key = ""
	btn.MediaID = ""
	btn.SubButtons = nil
	return btn
}

// SetScanCodePushButton 扫码推事件
func (btn *Button) SetScanCodePushButton(name, key string) *Button {
	btn.Type = "scancode_push"
	btn.Name = name
	btn.Key = key
	btn.URL = ""
	btn.MediaID = ""
	btn.SubButtons = nil
	return btn
}

// SetScanCodeWaitMsgButton 设置 扫码推事件且弹出"消息接收中"提示框
func (btn *Button) SetScanCodeWaitMsgButton(name, key string) *Button {
	btn.Type = "scancode_waitmsg"
	btn.Name = name
	btn.Key = key

	btn.URL = ""
	btn.MediaID = ""
	btn.SubButtons = nil
	return btn
}

// SetPicSysPhotoButton 设置弹出系统拍照发图按钮
func (btn *Button) SetPicSysPhotoButton(name, key string) *Button {
	btn.Type = "pic_sysphoto"
	btn.Name = name
	btn.Key = key

	btn.URL = ""
	btn.MediaID = ""
	btn.SubButtons = nil
	return btn
}

// SetPicPhotoOrAlbumButton 设置弹出拍照或者相册发图类型按钮
func (btn *Button) SetPicPhotoOrAlbumButton(name, key string) *Button {
	btn.Type = "pic_photo_or_album"
	btn.Name = name
	btn.Key = key

	btn.URL = ""
	btn.MediaID = ""
	btn.SubButtons = nil
	return btn
}

// SetPicWeixinButton 设置弹出微信相册发图器类型按钮
func (btn *Button) SetPicWeixinButton(name, key string) *Button {
	btn.Type = "pic_weixin"
	btn.Name = name
	btn.Key = key

	btn.URL = ""
	btn.MediaID = ""
	btn.SubButtons = nil
	return btn
}

// SetLocationSelectButton 设置 弹出地理位置选择器 类型按钮
func (btn *Button) SetLocationSelectButton(name, key string) *Button {
	btn.Type = "location_select"
	btn.Name = name
	btn.Key = key

	btn.URL = ""
	btn.MediaID = ""
	btn.SubButtons = nil
	return btn
}

// SetMediaIDButton  设置 下发消息(除文本消息) 类型按钮
func (btn *Button) SetMediaIDButton(name, mediaID string) *Button {
	btn.Type = "media_id"
	btn.Name = name
	btn.MediaID = mediaID

	btn.Key = ""
	btn.URL = ""
	btn.SubButtons = nil
	return btn
}

// SetViewLimitedButton  设置 跳转图文消息URL 类型按钮
func (btn *Button) SetViewLimitedButton(name, mediaID string) *Button {
	btn.Type = "view_limited"
	btn.Name = name
	btn.MediaID = mediaID

	btn.Key = ""
	btn.URL = ""
	btn.SubButtons = nil
	return btn
}

// SetMiniprogramButton  设置 跳转小程序 类型按钮 (公众号后台必须已经关联小程序)
func (btn *Button) SetMiniprogramButton(name, url, appID, pagePath string) *Button {
	btn.Type = "miniprogram"
	btn.Name = name
	btn.URL = url
	btn.AppID = appID
	btn.PagePath = pagePath

	btn.Key = ""
	btn.MediaID = ""
	btn.SubButtons = nil
	return btn
}

// NewSubButton 二级菜单
func NewSubButton(name string, subButtons []*Button) *Button {
	return (&Button{}).SetSubButton(name, subButtons)
}

// NewClickButton btn 为click类型
func NewClickButton(name, key string) *Button {
	return (&Button{}).SetClickButton(name, key)
}

// NewViewButton view类型
func NewViewButton(name, url string) *Button {
	return (&Button{}).SetViewButton(name, url)
}

// NewScanCodePushButton 扫码推事件
func NewScanCodePushButton(name, key string) *Button {
	return (&Button{}).SetScanCodePushButton(name, key)
}

// NewScanCodeWaitMsgButton 扫码推事件且弹出"消息接收中"提示框
func NewScanCodeWaitMsgButton(name, key string) *Button {
	return (&Button{}).SetScanCodeWaitMsgButton(name, key)
}

// NewPicSysPhotoButton 弹出系统拍照发图按钮
func NewPicSysPhotoButton(name, key string) *Button {
	return (&Button{}).SetPicSysPhotoButton(name, key)
}

// NewPicPhotoOrAlbumButton 弹出拍照或者相册发图类型按钮
func NewPicPhotoOrAlbumButton(name, key string) *Button {
	return (&Button{}).SetPicPhotoOrAlbumButton(name, key)
}

// NewPicWeixinButton 弹出微信相册发图器类型按钮
func NewPicWeixinButton(name, key string) *Button {
	return (&Button{}).SetPicWeixinButton(name, key)
}

// NewLocationSelectButton  弹出地理位置选择器 类型按钮
func NewLocationSelectButton(name, key string) *Button {
	return (&Button{}).SetLocationSelectButton(name, key)
}

// NewMediaIDButton  下发消息(除文本消息) 类型按钮
func NewMediaIDButton(name, mediaID string) *Button {
	return (&Button{}).SetMediaIDButton(name, mediaID)
}

// NewViewLimitedButton  跳转图文消息URL 类型按钮
func NewViewLimitedButton(name, mediaID string) *Button {
	return (&Button{}).SetViewLimitedButton(name, mediaID)
}

// NewMiniprogramButton  跳转小程序 类型按钮 (公众号后台必须已经关联小程序)
func NewMiniprogramButton(name, url, appID, pagePath string) *Button {
	return (&Button{}).SetMiniprogramButton(name, url, appID, pagePath)
}
