package demo

import (
	"github.com/gin-gonic/gin"
	"web/gopkg/gins"
	"web/handler/api/demo/request"
)

// DeleteDemo 删除
// @Tags APi:示例
// @Summary 删除
// @Description 删除
// @Produce	json
// @Param param body request.DeleteDemo true "请求参数"
// @Success 200 {object} services.BaseResult{data=view.WritingOnline} "操作成功"
// @Router /api/demo/delete [post]
func (h *Handler) DeleteDemo(ctx *gin.Context) {
	var req request.DeleteDemo
	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	res, err := h.demoService.DeleteDemo(ctx, req.DemoId)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	gins.StatusOK(ctx, res)
}
