package retcode

var (
	//router
	Success                = NewCode("000000", "Success")
	RequestPathNotFound    = NewCode("010000", "Request Path Not Found")
	RequestMethodNotAllow  = NewCode("010001", "Request Method Not Allowed")
	RequestIllegal         = NewCode("010002", "Request Illegal")
	RequestUnMarshalError  = NewCode("010003", "Request Body UnMarshal Error")
	RequestNoPermission    = NewCode("010004", "Request No Permission")
	RequestAuthCheckFail   = NewCode("020001", "Request Auth Check Fail")
	GenerateAuthTokenFail  = NewCode("020002", "Generate Auth Token Fail")
	RequestTokenEmpty      = NewCode("020003", "Request Token Empty")
	RequestTokenWrong      = NewCode("020004", "Request Token Wrong")
	RequestTokenAuthFail   = NewCode("020005", "Request Token Auth Fail")
	RequestTokenAuthExpire = NewCode("020006", "Request Token Expire")
	RequestUserGetFail     = NewCode("020007", "Get User Fail")

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

	GetResourceFail    = NewCode("300001", "Get Resource Fail")
	GetResourcesFail   = NewCode("300002", "Get Resources Fail")
	CreateResourceFail = NewCode("300003", "Create Resource Fail")
	UpdateResourceFail = NewCode("300004", "Update Resource Fail")
	DeleteResourceFail = NewCode("300005", "Delete Resource Fail")

	GetOperationsFail   = NewCode("400001", "Get Operations Fail")
	CreateOperationFail = NewCode("400002", "Create Operation Fail")
	UpdateOperationFail = NewCode("400003", "Update Operation Fail")
	DeleteOperationFail = NewCode("400004", "Delete Operation Fail")

	GetSystemConfigFail = NewCode("500001", "Get System Config Fail")

	GetUsersFail   = NewCode("600001", "Get Users Fail")
	CreateUserFail = NewCode("600002", "Create User Fail")
	UpdateUserFail = NewCode("600003", "Update User Fail")
	DeleteUserFail = NewCode("600004", "Delete User Fail")

	GetRolesFail   = NewCode("700001", "Get Roles Fail")
	CreateRoleFail = NewCode("700002", "Create Role Fail")
	UpdateRoleFail = NewCode("700003", "Update Role Fail")
	DeleteRoleFail = NewCode("700004", "Delete Role Fail")

	UnknownError = NewCode("999999", "Unknown Error")
)
