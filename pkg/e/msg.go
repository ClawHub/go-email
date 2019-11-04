package e

var MsgFlags = map[int]string{
	SUCCESS:          "ok",
	ERROR:            "fail",
	INVALID_PARAMS:   "请求参数错误",
	INVALID_TEMPLATE: "模板不支持",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
