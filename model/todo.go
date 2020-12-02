package model

type Todo struct { // 待办事项做为事务存到数据库中
	ID     int    `json:"id"` // 与前端交互用json数据格式
	Title  string `json:"title"`
	Status bool   `json:"status"` // 待办事项完成与否的状态
}

func (Todo) TableName() string {
	return "shuwen_checklist_todos"
}
