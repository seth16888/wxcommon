package mp

// DeletePublishReq 删除已发布文章
type DeletePublishReq struct {
	ArticleID string `json:"article_id" binding:"required" msg:"article_id required"`
	Index     int64  `json:"index,omitempty"` //index 要删除的文章在图文消息中的位置，第一篇编号为1，该字段不填或填0会删除全部文章
}

// ListPublishReq 获取成功发布文章列表
type ListPublishReq struct {
	Offset    int64 `json:"offset" binding:"required" msg:"offset required"`
	Count     int64 `json:"count" binding:"required" msg:"count required"`
	NoContent bool  `json:"no_content,omitempty"`
}

// PublishedArticleRes 已发布文章
type PublishedArticleRes struct {
	Title              string `json:"title"`
	Author             string `json:"author"`
	Digest             string `json:"digest"`
	Content            string `json:"content"`
	ContentSourceURL   string `json:"content_source_url"`
	ThumbMediaId       string `json:"thumb_media_id"`
	ThumbURL           string `json:"thumb_url"`
	ShowCoverPic       int64  `json:"show_cover_pic"`
	NeedOpenComment    int64  `json:"need_open_comment"`
	OnlyFansCanComment int64  `json:"only_fans_can_comment"`
	URL                string `json:"url"`
	IsDeleted          bool   `json:"is_deleted"`
}

// GetPublishedArticleRes 获取已发布文章详情
type GetPublishedArticleRes struct {
	NewsItem []PublishedArticleRes `json:"news_item"`
}

// GetPublishedArticleListRes 获取已发布文章列表
type GetPublishedArticleListRes struct {
	TotalCount int64                         `json:"total_count"`
	ItemCount  int64                         `json:"item_count"`
	Item       []GetPublishedArticleListItem `json:"item"`
}

// GetPublishedArticleListItem 获取已发布文章列表的单条数据
type GetPublishedArticleListItem struct {
	ArticleId  string `json:"article_id"`
	UpdateTime int64  `json:"update_time"`
	Content    struct {
		NewsItem []PublishedArticleRes `json:"news_item"`
	} `json:"content"`
}

// GetPublishStatusRes 获取发布状态
type GetPublishStatusRes struct {
	PublishId     string `json:"publish_id"`
	PublishStatus int64  `json:"publish_status"`
	ArticleId     string `json:"article_id"`
	ArticleDetail struct {
		Count int64 `json:"count"`
		Item  []struct {
			Idx        int64  `json:"idx"`
			ArticleURL string `json:"article_url"`
		} `json:"item"`
	} `json:"article_detail"`
	FailIdx []int64 `json:"fail_idx"`
}
