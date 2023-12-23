package pconline

import (
	"context"
	"go.dtapp.net/gorequest"
)

func (c *Client) request(ctx context.Context, url string, param gorequest.Params) (gorequest.Response, error) {

	// 创建请求
	client := gorequest.NewHttp()

	// 设置请求地址
	client.SetUri(url)

	// 设置参数
	client.SetParams(param)

	// 发起请求
	request, err := client.Get(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if c.gormLog.status {
		go c.gormLog.client.Middleware(ctx, request)
	}

	return request, err
}
