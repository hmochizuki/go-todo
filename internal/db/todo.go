package db

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID        uint      `gorm:"primaryKey;comment:TODOのID"`
	Name      string    `gorm:"size:100;comment:TODOの名称"`
	UserID    uuid.UUID `gorm:"not null;comment:ユーザーID(外部キー);index"`
	User      User      `gorm:"foreignKey:UserID;references:ID;comment:ユーザー情報"`
	Status    string    `gorm:"type:todo_status;default:todo;comment:TODOのステータス(todo,doing,pending,archived,done)"`
	CreatedAt time.Time `gorm:"autoCreateTime;comment:作成日時"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;comment:更新日時"`
	Deleted   bool      `gorm:"default:false;comment:削除フラグ"`
}
