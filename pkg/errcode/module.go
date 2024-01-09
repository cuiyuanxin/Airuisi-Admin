package errcode

var (
	ErrorUserUsernameNotExist  = NewError(20010000, "账号不存在")
	ErrorUserAccountIsDisabled = NewError(20010001, "账号已禁用")
	ErrorUserAccountIsDelete   = NewError(20010002, "账号已删除")
	ErrorUserAccountIsLocked   = NewError(20010003, "账号已锁定")

	ErrorUsersCheckMobileIsNo = NewError(200100000, "手机号不存在")
	ErrorUsersSendCodeFail    = NewError(200100020, "手机验证码发送失败")
	ErrorUsersCodeFault       = NewError(200100030, "验证码错误")
	ErrorUsersCodeOverdue     = NewError(200100040, "验证码已过期")
	ErrorUsersCodeStatus      = NewError(200100050, "验证码还在有效期内，请勿多次申请")
)
