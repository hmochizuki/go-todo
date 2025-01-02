package db

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string    `gorm:"comment:ユーザーの氏名"`
	Email     string    `gorm:"unique;comment:ユーザーのメールアドレス"`
	Password  string    `gorm:"comment:ユーザーのパスワード"`
	CreatedAt time.Time `gorm:"autoCreateTime;comment:作成日時"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;comment:更新日時"`
	Deleted   bool      `gorm:"default:false;comment:削除フラグ"`
}
