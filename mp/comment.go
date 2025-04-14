package mp

// OpenCommentReq 开启评论
type OpenCommentReq struct {
	MsgDataId int64 `json:"msg_data_id"` // 群发返回的msg_data_id
	Index     int64 `json:"index"`       // 多图文时，用来指定第几篇图文，从0开始，不带默认返回该msg_data_id的第一篇图文
}

// GetNewsCommentListReq 获取评论列表
type GetNewsCommentListReq struct {
	MsgDataId int64 `json:"msg_data_id"` // 群发返回的msg_data_id
	Index     int64 `json:"index"`       // 多图文时，用来指定第几篇图文，从0开始，不带默认返回该msg_data_id的第一篇图文
	Begin     int64 `json:"begin"`       // 起始位置
	Count     int64 `json:"count"`       // 获取数目（>=50会被拒绝）
	Type      int64 `json:"type"`        // type=0 普通评论+精选评论 type=1 普通评论 type=2 精选评论
}

// CommentData 评论数据结构
type CommentData struct {
	UserCommentId int64  `json:"user_comment_id"`
	OpenId        string `json:"openid"`
	CreateTime    int64  `json:"create_time"`
	Content       string `json:"content"`
	CommentType   int64  `json:"comment_type"`
	Reply         struct {
		Content    string `json:"content"`
		CreateTime int64  `json:"create_time"`
	} `json:"reply"`
}

// GetNewsCommentListRes 获取图文评论列表
type GetNewsCommentListRes struct {
	WXError
	Total   int64         `json:"total"`
	Comment []CommentData `json:"comment"`
}

// UserCommentIdData 评论ID
type UserCommentIdData struct {
	UserCommentId int64 `json:"user_comment_id"`
	Index         int64 `json:"index"`
	MsgDataId     int64 `json:"msg_data_id"`
}

// ReplyCommentData 回复评论
type ReplyCommentData struct {
	UserCommentIdData
	Content string `json:"content"`
}
