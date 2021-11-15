package retcode

var (
	//router
	Success                 = NewError("000000", "Success")
	RequestPathNotFound     = NewError("010000", "Request Path Not Found")
	RequestMethodNotAllow   = NewError("010001", "Request Method Not Allowed")
	RequestIllegal          = NewError("010002", "Request Illegal")
	RequestUnMarshalError   = NewError("010003", "Request Body UnMarshal Error")
	RequestNoPermission     = NewError("010004", "Request No Permission")
	RequestAuthNotExists    = NewError("020001", "Request Auth Not Exists")
	GenerateAuthTokenFail   = NewError("020002", "Generate Auth Token Fail")
	RequestTokenEmpty       = NewError("020003", "Request Token Empty")
	RequestTokenAuthFail    = NewError("020004", "Request Token Auth Fail")
	RequestTokenAuthTimeout = NewError("020005", "Request Token Auth Timeout")

	//module
	GetArticleFail    = NewError("100001", "Get Article Fail")
	GetArticlesFail   = NewError("100002", "Get Articles Fail")
	CreateArticleFail = NewError("100003", "Create Article Fail")
	UpdateArticleFail = NewError("100004", "Update Article Fail")
	DeleteArticleFail = NewError("100005", "Delete Article Fail")

	GetTagsFail   = NewError("200001", "Get Tags Fail")
	CreateTagFail = NewError("200002", "Create Tag Fail")
	UpdateTagFail = NewError("200003", "Update Tag Fail")
	DeleteTagFail = NewError("200004", "Delete Tag Fail")

	UnknownError = NewError("999999", "Unknown Error")
)
