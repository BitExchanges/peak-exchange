package utils

import "fmt"

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
	SuccessResponse     = Response{Head: map[string]string{"code": "0", "msg": "success"}}
	PageSuccessResponse = PageResponse{Head: map[string]string{"code": "0", "msg": "success"}}
)

func BuildError(code string) Response {
	return Response{Head: map[string]string{"code": code}}
}

func (errResponse Response) Error() string {
	return fmt.Sprintf("code:%s;msg:%s", errResponse.Head["code"], errResponse.Head["msg"])
}

func (pageRes PageResponse) Init(data interface{}, page, count, perPage int) {
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
	pageRes.Body = body
}
