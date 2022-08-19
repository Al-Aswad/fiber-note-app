package requests

type CreateTodo struct {
	ActivityGroupID int    `validate:"required" json:"activity_group_id"`
	Title           string `validate:"required" json:"title"`
	IsActive        bool   `validate:"required" json:"is_active"`
	Priority        string `validate:"required" json:"priority"`
}
