package models

import "time"

// Role 角色对象
type Role struct {
	RecordID  string    `json:"record_id"`               // 记录ID
	Name      string    `json:"name" binding:"required"` // 角色名称
	Sequence  int       `json:"sequence"`                // 排序值
	Memo      string    `json:"memo"`                    // 备注
	Creator   string    `json:"creator"`                 // 创建者
	CreatedAt time.Time `json:"created_at"`              // 创建时间
}
