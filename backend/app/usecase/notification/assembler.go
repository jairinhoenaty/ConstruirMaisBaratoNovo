package notification_usecase

type SendAppNotificationAssembler struct {
	UserIds []uint `json:"user_ids"`
	Title   string `json:"title"`
	Body    string `json:"body"`
}
