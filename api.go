package pconline

import (
	"context"
	"go.dtapp.net/gorequest"
)

func (c *Client) Get(ctx context.Context, _method string, notMustParams ...gorequest.Params) ([]byte, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+_method, params)
	// 定义
	return request.ResponseBody, err
}
