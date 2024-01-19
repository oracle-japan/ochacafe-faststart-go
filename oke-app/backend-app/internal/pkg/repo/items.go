// サンプルアプリのレポジトリ
package repo

import (
	"gorm.io/gorm"
)

// Items構造体定義
type Items struct {
	gorm.Model
	Name       string
	Date       string
	Topics     string
	Presenters string
}

