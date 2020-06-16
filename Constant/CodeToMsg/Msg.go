package CodeToMsg

var MsgFlags = map[int]string{
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "auth error， your token is invalid",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "auth error, your token expired",

	IP_FORBIDDEN: "your ip is in ipBlackList.",

	REQUEST_PARAM_ERROR: "your request has some bad parameters",

	SUCCESS:     "auth success",
	RPC_SUCCESS: "rpc dial success",
}

// 根据code获取Msg信息
func GetMsg(code int) string {

	s := MsgFlags[code]

	return s
}
