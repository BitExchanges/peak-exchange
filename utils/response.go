package utils

const (
	AccessDenied       = "10001" //无访问权限
	NotFound           = "10002" //暂无数据
	ParamError         = "10003" //参数错误
	IllegalToken       = "10004" //非法token
	OperateError       = "10005" //操作失败
	UserNameOrPwdError = "10006" //用户名或密码错误
	UserNotFound       = "10007" //用户不存在
)

type Response struct {
	Head map[string]string `json:"head"`
	Body interface{}       `json:"body"`
}

type Page struct {
	CurrentPage  int         `json:"current_page"`  //当前页
	TotalPage    int         `json:"total_page"`    //总条数
	PerPage      int         `json:"per_page"`      //每页条数
	NextPage     int         `json:"next_page"`     //下一页
	PreviousPage int         `json:"previous_page"` //上一个
	Data         interface{} `json:"data"`
}

type PageResponse struct {
	Head map[string]string `json:"head"`
	Body interface{}       `json:"body"`
}

// 定义传统以及分页成功返回
var (
	SuccessBase         = map[string]interface{}{"code": "0", "msg": "success"}
	SuccessResponse     = Response{Head: map[string]string{"code": "0", "msg": "success"}}
	PageSuccessResponse = PageResponse{Head: map[string]string{"code": "0", "msg": "success"}}
)

// 组装错误返回
func BuildError(code, msg string) map[string]string {
	return map[string]string{"code": code, "msg": msg}
}

// 组装成功返回
func Success(data interface{}) map[string]interface{} {
	return map[string]interface{}{"code": "0", "msg": "success", "data": data}
}

// 分页初始化
func PageInit(data interface{}, page, count, perPage int) map[string]interface{} {
	//总条数 除以 每页条数 = 总页数
	//如果有余数 则总页数+1
	totalPage := count / perPage
	if (count % perPage) != 0 {
		totalPage += 1
	}

	//如果下一个大于总页数，那么下一个等于总页数
	nextPage := page + 1
	if nextPage > totalPage {
		nextPage = totalPage
	}

	//计算上一页，如果当前页减- 小于1 那么上一页等于第一页
	previousPage := page - 1
	if previousPage < 1 {
		previousPage = 1
	}

	body := Page{}
	body.TotalPage = totalPage
	body.PreviousPage = previousPage
	body.NextPage = nextPage
	body.CurrentPage = page
	body.Data = data
	body.PerPage = perPage
	return map[string]interface{}{"code": "0", "msg": "success", "data": body}
}
