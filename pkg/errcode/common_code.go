package errcode

var (
	Success = NewError(0, "成功")
	ServerError = NewError(10000000, "服务内部错误")
	InvalidParam = NewError(10000001, "入参错误")
	NotFound = NewError(10000002, "找不到")
	UnauthorizedAuthNotExist = NewError(10000003, "鉴权失败，找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError = NewError(10000004, "鉴权失败，Token错误")
	UnauthorizedTokenTimeout = NewError(10000005, "鉴权失败，Token超时")
	UnauthorizedTokenGenerate = NewError(10000006, "鉴权失败，Token生成失败")
	TooManyRequest = NewError(10000007, "请求过多")
	ErrorGetListFail = NewError(20010001, "获取列表失败")
	ErrorCreateFail = NewError(20010002, "创建失败")
	ErrorUpdateFail = NewError(20010003, "更新失败")
	ErrorDeleteFail = NewError(20010004, "删除失败")
	ErrorCountFail = NewError(20010005, "统计失败")
)

