package errcode

var (
	Success               = NewError(0, "成功")
	Fail                  = NewError(1000000, "失败")
	InvalidParams         = NewError(1000001, "无效参数")
	Unauthorized          = NewError(1000002, "认证错误")
	NotFound              = NewError(1000003, "没有找到")
	UnKnow                = NewError(1000004, "未知")
	DeadLineExceded       = NewError(1000005, "超出最后访问期限")
	AccessDenied          = NewError(1000006, "访问被拒绝")
	LimitExceed           = NewError(1000007, "访问限制")
	MethodNotAllowed        = NewError(1000008, "不支持该方法")
)