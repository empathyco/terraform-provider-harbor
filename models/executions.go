package models

var PathExecution = "/replication/executions"

type ExecutionBody struct {
	Description string `json:"description,omitempty"`
	Status string `json:"status,omitempty"`
	StatusText string `json:"status_text,omitempty"`
	Trigger string `json:"trigger,omitempty"`
	StartTime string `json:"start_time,omitempty"`
	Failed int `json:"failed,omitempty"`
	Succeed int `json:"succeed,omitempty"`
	Stopped int `json:"stopped,omitempty"`
	EndTime string `json:"end_time,omitempty"`
	InProgress int `json:"in_progress,omitempty"`
	Total int `json:"total,omitempty"`
	Id int `json:"id,omitempty"`
	PolicyId int `json:"policy_id,omitempty"`
}
