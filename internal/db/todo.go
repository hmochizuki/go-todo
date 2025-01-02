package db

import "time"

type Todo struct {
	ID        uint      `gorm:"primaryKey;comment:TODOのID"`
	Name      string    `gorm:"comment:TODOの名称"`
	UserID    string    `gorm:"not null;comment:ユーザーID(外部キー);index"`            // UserIDを外部キーとして設定
	User      User      `gorm:"foreignKey:UserID;references:ID;comment:ユーザー情報"` // UserIDを外部キーとしてUserテーブルを参照
	CreatedAt time.Time `gorm:"autoCreateTime;comment:作成日時"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;comment:更新日時"`
	Deleted   bool      `gorm:"default:false;comment:削除フラグ"`
}
