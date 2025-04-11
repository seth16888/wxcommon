package paths

const (
	Path_Get_Stable_Access_Token = "/cgi-bin/stable_token"

	Path_Batch_Get_Material = "/cgi-bin/material/batchget_material"
	Path_Get_Material_Count = "/cgi-bin/material/get_materialcount"
	Path_Upload_Media       = "/cgi-bin/material/add_material"
	Path_Del_Material       = "/cgi-bin/material/del_material"

	Path_Upload_Temporary_Media = "/cgi-bin/media/upload"
	Path_Upload_News_Image      = "/cgi-bin/media/uploadimg"

	Path_Get_Member_Info       = "/cgi-bin/user/info"
	Path_Batch_Get_Member_Info = "/cgi-bin/user/info/batchget"
	Path_Get_Member_List       = "/cgi-bin/user/get"

	Path_Get_Member_Tags         = "/cgi-bin/tags/getidlist"
	Path_Update_Member_Remark    = "/cgi-bin/user/info/updateremark"
	Path_Get_Black_List          = "/cgi-bin/tags/members/getblacklist"
	Path_Batch_Add_Black_List    = "/cgi-bin/tags/members/batchblacklist"
	Path_Batch_Remove_Black_List = "/cgi-bin/tags/members/batchunblacklist"

	Path_Create_QRCode = "/cgi-bin/qrcode/create"

	// 创建标签
	Path_Get_Tags        = "/cgi-bin/tags/get"
	Path_Create_Tag      = "/cgi-bin/tags/create"
	Path_Update_Tag      = "/cgi-bin/tags/update"
	Path_Delete_Tag      = "/cgi-bin/tags/delete"
	Path_Get_Tag_Members = "/cgi-bin/user/tag/get"
	Path_Batch_Tagging   = "/cgi-bin/tags/members/batchtagging"
	Path_Batch_Untagging = "/cgi-bin/tags/members/batchuntagging"

	// Menu
	Path_Get_Current_SelfMenu   = "/cgi-bin/get_current_selfmenu_info"
	Path_Del_ConditionalMenu    = "/cgi-bin/menu/delconditional"
	Path_Try_MatchMenu          = "/cgi-bin/menu/trymatch"
	Path_Del_Menu               = "/cgi-bin/menu/delete"
	Path_Get_Menu               = "/cgi-bin/menu/get"
	Path_Create_Menu            = "/cgi-bin/menu/create"
	Path_Create_ConditionalMenu = "/cgi-bin/menu/addconditional"

	// Shorten 短key托管
	Path_Gen_Shorten   = "/cgi-bin/shorten/gen"
	Path_Fetch_Shorten = "/cgi-bin/shorten/fetch"

	// template message
	Path_Send_TemplateMessage = "/cgi-bin/message/template/send"
	Path_Send_TplSubscribeMsg = "/cgi-bin/message/template/subscribe"
	Path_Get_BlockedMsg       = "/wxa/sec/queryblocktmplmsg"
	Path_Set_Industry         = "/cgi-bin/template/api_set_industry"
	Path_Get_Industry         = "/cgi-bin/template/get_industry"
	Path_Get_TemplateId       = "/cgi-bin/template/api_add_template"
	Path_Get_AllPrivateTmpl   = "/cgi-bin/template/get_all_private_template"
	Path_Del_Template         = "/cgi-bin/template/del_private_template"

	// subscription message
	Path_Send_SubscribeMessage     = "/cgi-bin/message/subscribe/bizsend"
	Path_Get_Template_List         = "/wxaapi/newtmpl/gettemplate"
	Path_Add_Template              = "/wxaapi/newtmpl/addtemplate"
	Path_Del_Subscription_Template = "/wxaapi/newtmpl/deltemplate"
	Path_Get_Category              = "/wxaapi/newtmpl/getcategory"
	Path_Get_PubTpl_KeyWorks       = "/wxaapi/newtmpl/getpubtemplatekeywords"
	Path_Get_PubTpl_Titles         = "/wxaapi/newtmpl/getpubtemplatetitles"

	// customer service
	Path_Get_KfList                = "/cgi-bin/customservice/getkflist"
	Path_Get_OnlineKfList          = "/cgi-bin/customservice/getonlinekflist"
	Path_Update_KfStatus           = "/cgi-bin/message/custom/typing"
	Path_Add_KfAccount             = "/customservice/kfaccount/add"
	Path_Update_KfAccount          = "/customservice/kfaccount/update"
	Path_Del_KfAccount             = "/customservice/kfaccount/del"
	Path_Update_KfAvatar           = "/customservice/kfaccount/uploadheadimg"
	Path_Invite_Worker             = "/customservice/kfaccount/inviteworker"
	Path_Get_MsgRecord             = "/customservice/msgrecord/getmsglist"
	Path_Create_Session            = "/customservice/kfsession/create"
	Path_Close_Session             = "/customservice/kfsession/close"
	Path_Get_SessionStatus         = "/customservice/kfsession/getsession"
	Path_Get_KFSessionList         = "/customservice/kfsession/getsessionlist"
	Path_Get_UnacceptedSessionList = "/customservice/kfsession/getwaitcase"
	Path_KF_Send_Message           = "/cgi-bin/message/custom/send"

	// publish

	Path_Publish_Submit = "/cgi-bin/freepublish/submit"
	Path_Publish_Delete = "/cgi-bin/freepublish/delete"
	Path_Publish_Get    = "/cgi-bin/freepublish/getarticle"
	Path_Publish_List   = "/cgi-bin/freepublish/batchget"
	Path_Publish_Status = "/cgi-bin/freepublish/get"

	// draft

	Path_Draft_Switch      = "/cgi-bin/draft/switch"
	Path_Draft_Get_Prd_DOM = "/channels/ec/service/product/getcardinfo"
	Path_Draft_Delete      = "/cgi-bin/draft/delete"
	Path_Draft_Count       = "/cgi-bin/draft/count"
	Path_Draft_Get         = "/cgi-bin/draft/get"
	Path_Draft_Add         = "/cgi-bin/draft/add"
	Path_Draft_Update      = "/cgi-bin/draft/update"
	Path_Draft_List        = "/cgi-bin/draft/batchget"

	// comment

	Path_Comment_Open        = "/cgi-bin/comment/open"
	Path_Comment_Close       = "/cgi-bin/comment/close"
	Path_Comment_List        = "/cgi-bin/comment/list"
	Path_Comment_Markelect   = "/cgi-bin/comment/markelect"
	Path_Comment_UnMarkelect = "/cgi-bin/comment/unmarkelect"
	Path_Comment_Delete      = "/cgi-bin/comment/delete"
	Path_Comment_Reply       = "/cgi-bin/comment/reply/add"
	Path_Comment_DeleteReply = "/cgi-bin/comment/reply/delete"
)
