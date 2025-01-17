package errmsg

const (
	SUCCESE = 200
	ERROR   = 500
	//code = 1000  用户模块的报错
	ERROR_USERNAME_USED     = 1001
	ERROR_PASSWORD_WRONG    = 1002
	ERROR_USER_NOT_EXIST    = 1003
	ERROR_TOKEN_EXIST       = 1004
	ERROR_TOKEN_RUNTIME     = 1005
	ERROR_TOKEN_WRONG       = 1006
	ERROR_TOKEN_TYPE_WRONG  = 1007
	ERROR_USER_NO_RIGHT     = 1008
	ERROR_EMPTY_USERNAME    = 1009
	ERROR_APPLY_NEW_ACCOUNT = 1100

	//code = 2000 文章模块的错误
	ERROR_ART_NO_EXIST = 2001

	//code = 3000 分类模块的错误
	ERROR_CATENAME_USED  = 3001
	ERROR_CATE_NOT_EXIST = 3002

	// code = 5000 队伍模块的报错
	ERROR_TEAM_NOT_EXIT = 5001

	// code = 5100 队伍目录模块的报错
	ERROR_TEAM_ITEM_EXIT = 5101

	// code = 6000 Admin 模块的报错
	ERROR_EXITING_ITEM      = 6001
	ERROR_CONNECTING_SERVER = 6002
	ERROR_ARG               = 6003

	//code = 7000 Lang 语言模块的报错
	ERROR_LANG_USED      = 7001
	ERROR_LANG_NOT_EXIST = 7002
	ERROR_LANG_EMPTY     = 7003

	// code = 10000 条件和需求
	REQUIRE_PASSWORD_PERMISSION = 10307
	FUNCTION_UNDER_DEVELOP      = 10308

	//code = 8000 avatar
	ERROR_URL_EMPTY = 8001
)

var codeMsg = map[int]string{
	SUCCESE:                     "OK",
	ERROR:                       "FAIL",
	ERROR_USERNAME_USED:         "用户名已存在!",
	ERROR_PASSWORD_WRONG:        "密码错误",
	ERROR_USER_NOT_EXIST:        "用户不存在",
	ERROR_TOKEN_EXIST:           "TOKEN 不存在",
	ERROR_TOKEN_RUNTIME:         "TOKEN 已过期",
	ERROR_TOKEN_WRONG:           "Token 不正确",
	ERROR_TOKEN_TYPE_WRONG:      "TOKEN 格式错误",
	ERROR_ART_NO_EXIST:          "文章不存在",
	ERROR_CATENAME_USED:         "该分类已存在",
	ERROR_CATE_NOT_EXIST:        "该查询不存在",
	ERROR_USER_NO_RIGHT:         "该用户没有权限",
	ERROR_TEAM_NOT_EXIT:         "团队不存在",
	ERROR_TEAM_ITEM_EXIT:        "团队主题已经存在",
	REQUIRE_PASSWORD_PERMISSION: "需要项目密码",
	ERROR_EXITING_ITEM:          "该用户名下还有项目，不允许删除。请先将其项目删除或者重新分配/转让",
	ERROR_APPLY_NEW_ACCOUNT:     "ADMINISTRATOR 没有开启申请账号权限",
	ERROR_LANG_USED:             "该语言已存在",
	ERROR_LANG_NOT_EXIST:        "该语言查询不存在",
	ERROR_LANG_EMPTY:            "请选择其中一种语言",
	FUNCTION_UNDER_DEVELOP:      "功能还在开发中",
	ERROR_CONNECTING_SERVER:     "无法连接LDAP",
	ERROR_ARG:                   "参数错误",
	ERROR_URL_EMPTY:             "跳转域名不能为空",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
