package task

type UpdateTaskRequest struct {
	Text   *string `json:"text"`
	IsDone *bool   `json:"is_done"`
}
