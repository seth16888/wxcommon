package mp

// UploadMediaRes 新增临时素材返回
type UploadMediaRes struct {
	Type      string `json:"type"`
	MediaID   string `json:"media_id"`
	CreatedAt int64  `json:"created_at"`
	// 缩略图media_id
	ThumbMediaId string `json:"thumb_media_id,omitempty"`
	URL          string `json:"url,omitempty"`
}

// UploadMaterialRes 新增永久素材返回
type UploadMaterialRes struct {
	MediaID string `json:"media_id"`
	Url     string `json:"url"`
}

// GetMaterialCountRes 从WX获取素材总数返回结果
type GetMaterialCountRes struct {
	VoiceCount int `json:"voice_count"`
	VideoCount int `json:"video_count"`
	ImageCount int `json:"image_count"`
	NewsCount  int `json:"news_count"`
}

// BatchGetMaterialRes 从WX获取素材列表返回结果(图片、语音、视频)
type BatchGetMaterialRes struct {
	TotalCount int            `json:"total_count"` // 该类型的素材的总数
	ItemCount  int            `json:"item_count"`  // 本次调用获取的素材的数量
	Items      []MaterialInfo `json:"item"`        // 本次调用获取的素材列表
}

// MaterialInfo 素材列表中素材的信息
// Note:
//	只适用：图片、语音、视频
type MaterialInfo struct {
	MediaId    string `json:"media_id"`    // 素材id
	Name       string `json:"name"`        // 文件名称
	UpdateTime int64  `json:"update_time"` // 最后更新时间
	URL        string `json:"url"`         // 当获取的列表是图片素材列表时, 该字段是图片的URL
}

// Article 图文素材
type Article struct {
	ThumbMediaId     string `json:"thumb_media_id"`               // 图文消息的封面图片素材id(必须是永久mediaID)
	Title            string `json:"title"`                        // 标题
	Author           string `json:"author,omitempty"`             // 作者
	Digest           string `json:"digest,omitempty"`             // 图文消息的摘要, 仅有单图文消息才有摘要, 多图文此处为空
	Content          string `json:"content"`                      // 图文消息的具体内容, 支持HTML标签, 必须少于2万字符, 小于1M, 且此处会去除JS
	ContentSourceURL string `json:"content_source_url,omitempty"` // 图文消息的原文地址, 即点击"阅读原文"后的URL
	ShowCoverPic     int    `json:"show_cover_pic"`               // 是否显示封面, 0为false, 即不显示, 1为true, 即显示
	URL              string `json:"url,omitempty"`                // !!!创建时不需要此参数!!! 图文页的URL, 文章创建成功以后, 会由微信自动生成
}

// BatchGetNewsRes 从WX获取图文列表返回结果
type BatchGetNewsRes struct {
	TotalCount int        `json:"total_count"` // 该类型的素材的总数
	ItemCount  int        `json:"item_count"`  // 本次调用获取的素材的数量
	Items      []NewsInfo `json:"item"`        // 本次调用获取的素材列表
}

// NewsInfo 图文列表中图文的信息
type NewsInfo struct {
	MediaId    string `json:"media_id"`    // 素材id
	UpdateTime int64  `json:"update_time"` // 最后更新时间
	Content    struct {
		Articles []Article `json:"news_item,omitempty"`
	} `json:"content"`
}

// VideoDescription 上传永久素材：视频的描述
type VideoDescription struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
