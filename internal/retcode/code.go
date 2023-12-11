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

	GetTagsFail   = NewCode("110001", "Get Tags Fail")
	CreateTagFail = NewCode("110002", "Create Tag Fail")
	UpdateTagFail = NewCode("110003", "Update Tag Fail")
	DeleteTagFail = NewCode("110004", "Delete Tag Fail")

	GetResourceFail    = NewCode("120001", "Get Resource Fail")
	GetResourcesFail   = NewCode("120002", "Get Resources Fail")
	CreateResourceFail = NewCode("120003", "Create Resource Fail")
	UpdateResourceFail = NewCode("120004", "Update Resource Fail")
	DeleteResourceFail = NewCode("120005", "Delete Resource Fail")

	GetOperationsFail   = NewCode("130001", "Get Operations Fail")
	CreateOperationFail = NewCode("130002", "Create Operation Fail")
	UpdateOperationFail = NewCode("130003", "Update Operation Fail")
	DeleteOperationFail = NewCode("130004", "Delete Operation Fail")

	GetSystemConfigFail = NewCode("140001", "Get System Config Fail")

	GetUsersFail   = NewCode("150001", "Get Users Fail")
	CreateUserFail = NewCode("150002", "Create User Fail")
	UpdateUserFail = NewCode("150003", "Update User Fail")
	DeleteUserFail = NewCode("150004", "Delete User Fail")

	GetRolesFail   = NewCode("160001", "Get Roles Fail")
	CreateRoleFail = NewCode("160002", "Create Role Fail")
	UpdateRoleFail = NewCode("160003", "Update Role Fail")
	DeleteRoleFail = NewCode("160004", "Delete Role Fail")

	GetRestsFail   = NewCode("170001", "Get Rests Fail")
	CreateRestFail = NewCode("170002", "Create Rest Fail")
	UpdateRestFail = NewCode("170003", "Update Rest Fail")
	DeleteRestFail = NewCode("170004", "Delete Rest Fail")

	UnknownError = NewCode("999999", "Unknown Error")
)
