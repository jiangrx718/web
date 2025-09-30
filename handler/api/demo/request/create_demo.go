package request

type CreateDemoParams struct {
	Name        string `json:"name" binding:"required"`         // 文档标题
	FileType    int    `json:"file_type" binding:"required"`    // 文档类型,1可研报告,2需求报告
	ProjectType int    `json:"project_type" binding:"required"` // 项目类型,1产品购置类,2开发实施类,3数据工程类,4咨询、运维类
}
