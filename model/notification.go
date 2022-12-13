package model

type NotificationItem struct {
	Title   string `json:"title" binding:"required"`
	Body    string `json:"body" binding:"required"`
	Payload struct {
		Id       int      `json:"id" binding:"required"`
		Type     string   `json:"type" binding:"required"`
		Date     int      `json:"date" binding:"required"`
		Title    string   `json:"title" binding:"required"`
		Content  string   `json:"content" binding:"required"`
		Image    string   `json:"image"`
		Url      string   `json:"url"`
		IsRead   bool     `json:"isRead"`
		FundCode []string `json:"fundCode" binding:"required"`
	} `json:"payload" binding:"required"`
	DeviceTokens []string `json:"tokens" binding:"required"`
}
type NotificationResponse struct {
	SuccessCount int `json:"successCount" binding:"required"`
	FailureCount int `json:"failureCount" binding:"required"`
}
