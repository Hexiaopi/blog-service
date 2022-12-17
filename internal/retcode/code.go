package retcode

var (
	//router
	Success                 = NewCode("000000", "Success")
	RequestPathNotFound     = NewCode("010000", "Request Path Not Found")
	RequestMethodNotAllow   = NewCode("010001", "Request Method Not Allowed")
	RequestIllegal          = NewCode("010002", "Request Illegal")
	RequestUnMarshalError   = NewCode("010003", "Request Body UnMarshal Error")
	RequestNoPermission     = NewCode("010004", "Request No Permission")
	RequestAuthCheckFail    = NewCode("020001", "Request Auth Check Fail")
	GenerateAuthTokenFail   = NewCode("020002", "Generate Auth Token Fail")
	RequestTokenEmpty       = NewCode("020003", "Request Token Empty")
	RequestTokenAuthFail    = NewCode("020004", "Request Token Auth Fail")
	RequestTokenAuthTimeout = NewCode("020005", "Request Token Auth Timeout")

	//module
	GetArticleFail    = NewCode("100001", "Get Article Fail")
	GetArticlesFail   = NewCode("100002", "Get Articles Fail")
	CreateArticleFail = NewCode("100003", "Create Article Fail")
	UpdateArticleFail = NewCode("100004", "Update Article Fail")
	DeleteArticleFail = NewCode("100005", "Delete Article Fail")

	GetTagsFail   = NewCode("200001", "Get Tags Fail")
	CreateTagFail = NewCode("200002", "Create Tag Fail")
	UpdateTagFail = NewCode("200003", "Update Tag Fail")
	DeleteTagFail = NewCode("200004", "Delete Tag Fail")

	UnknownError = NewCode("999999", "Unknown Error")
)
