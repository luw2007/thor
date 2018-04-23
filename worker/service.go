package service

import (
	"context"

	"github.com/luw2007/thor"
	"github.com/luw2007/thor/res"
)

// WorkerService 工作服务
type WorkerService interface {
	PostResource(ctx context.Context, metas []res.Meta) (reply thor.Reply)
	PostJob(ctx context.Context, job thor.Job) (reply thor.Reply)
}