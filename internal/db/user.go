package db

import "time"

type User struct {
	ID        string    `gorm:"primaryKey;comment:ユーザーのID"`
	Name      string    `gorm:"comment:ユーザーの氏名"`
	Email     string    `gorm:"comment:ユーザーのメールアドレス"`
	Password  string    `gorm:"comment:ユーザーのパスワード"`
	CreatedAt time.Time `gorm:"autoCreateTime;comment:作成日時"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;comment:更新日時"`
	Deleted   bool      `gorm:"default:false;comment:削除フラグ"`
}
