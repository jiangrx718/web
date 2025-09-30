package request

type DeleteDemo struct {
	DemoId string `json:"demo_id" binding:"required"` // 案例ID
}
