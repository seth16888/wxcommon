package mp

// NewsArticle 草稿箱的图文/图片消息
type NewsArticle struct {
	ArticleType        string `json:"article_type,omitempty"`
	Title              string `json:"title"`
	Author             string `json:"author,omitempty"`
	Digest             string `json:"digest,omitempty"`
	Content            string `json:"content"`
	ContentSourceURL   string `json:"content_source_url,omitempty"` // 图文消息的原文地址，当用户点击"阅读原文"时，将原文地址跳转至该地址
	ThumbMediaID       string `json:"thumb_media_id"`
	NeedOpenComment    int    `json:"need_open_comment,omitempty"`
	OnlyFansCanComment int    `json:"only_fans_can_comment,omitempty"`
	PicCrop2351        string `json:"pic_crop_2351,omitempty"`
	PicCrop11          string `json:"pic_crop_11,omitempty"`

	// 图片

	// 图片消息里的图片相关信息，图片数量最多为20张，首张图片即为封面图
	ImageInfo struct {
		ImageList []struct {
			ImageMediaId string `json:"image_media_id"` // 图片消息里的图片素材id（必须是永久MediaID）
		} `json:"image_list"`
	} `json:"image_info,omitempty"`

	CoverInfo struct {
		CropPercentList []struct {
			Ratio string `json:"ratio"`
			X1    string `json:"x1"`
			Y1    string `json:"y1"`
			X2    string `json:"x2"`
			Y2    string `json:"y2"`
		} `json:"crop_percent_list"`
	} `json:"cover_info,omitempty"`

	ProductInfo struct {
		FooterProductInfo struct {
			ProductKey string `json:"product_key"`
		} `json:"footer_product_info"`
	} `json:"product_info,omitempty"`
}

// GetProductDOMReq 获取商品DOM
type GetProductDOMReq struct {
	ProductId   string `json:"product_id" binding:"required" msg:"product_id required"`
	ArticleType string `json:"article_type" binding:"required" msg:"article_type required"`
	CardType    int    `json:"card_type" binding:"required" msg:"card_type required"`
}

// GetProductDOMRes 获取商品DOM结果
type GetProductDOMRes struct {
	Dom        string `json:"DOM"`
	ProductKey string `json:"product_key"`
}

// UpdateDraftReq 更新草稿
type UpdateDraftReq struct {
	MediaId  string `json:"media_id"`
	Index    int    `json:"index"`
	Articles struct {
		NewsArticle
	} `json:"articles"`
}

// BatchGetDraftReq 批量获取草稿
type BatchGetDraftReq struct {
	Offset    int `json:"offset"`
	Count     int `json:"count"`
	NoContent int `json:"no_content"`
}

// GetDraftRes 获取草稿
type GetDraftRes struct {
	WXError
	NewsItem []DraftRes `json:"news_item"`
}

type DraftRes struct {
	NewsArticle
	URL string `json:"url"`
}

// AddDraftReq 添加草稿
type AddDraftReq struct {
	Articles []NewsArticle `json:"articles"`
}

// GetDraftListRes 获取草稿列表
type GetDraftListRes struct {
	TotalCount int64              `json:"total_count"`
	ItemCount  int64              `json:"item_count"`
	Item       []GetDraftListItem `json:"item"`
}

// GetDraftListItem 获取草稿列表的单条数据
type GetDraftListItem struct {
	MediaId    string `json:"media_id"`
	UpdateTime int64  `json:"update_time"`
	Content    struct {
		NewsItem []DraftRes `json:"news_item"`
	} `json:"content"`
}
