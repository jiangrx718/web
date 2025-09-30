package demo

import (
	"context"
	"web/gopkg/log"
	"web/internal/g"
	"web/internal/model"

	"web/gopkg/utils"

	"go.uber.org/zap"
)

func (d *Dao) Create(ctx context.Context, name string, fileType int, projectType int, content string) (*model.Demo, error) {
	logPrefix := "/internal/dao/demo: Dao.Create()"
	log.Sugar().Info(logPrefix, "操作记录：", name)
	demoItem := model.Demo{
		DemoId:      utils.GenUUIDWithoutUnderline(),
		Name:        name,
		FileType:    fileType,
		ProjectType: projectType,
		Content:     content,
	}
	if err := g.Demo.Debug().Create(&demoItem); err != nil {
		log.Sugar().Error(ctx, logPrefix, zap.Any("demo record", demoItem))
		return nil, err
	}
	return &demoItem, nil
}
