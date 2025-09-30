package service

import (
	"context"
	"web/gopkg/gorms"
	"web/gopkg/services"
)

type Demo interface {
	CreateDemo(ctx context.Context, name string, fileType int, projectType int) (services.Result, error)
	DeleteDemo(ctx context.Context, demoId string) (services.Result, error)
	GetDemo(ctx context.Context, demoId string) (services.Result, error)

	UpdateDemo(ctx context.Context, writingOnlineId string, content string) (services.Result, error)
	PagingDemo(ctx context.Context, page gorms.Page) (services.Result, error)
}
