package models

import (
	"context"

	"gorm.io/gorm"
)

type ctxKey string

const (
	ctxKeyOperator ctxKey = "operator"
)

func GetCurrentOperator(db *gorm.DB) map[string]any {
	if db != nil && db.Statement != nil && db.Statement.Context != nil {
		if v := db.Statement.Context.Value(ctxKeyOperator).(map[string]any); v != nil {
			return v
		}
	}
	return nil
}

func SetCurrentOperator(params map[string]any) context.Context {
	return context.WithValue(context.Background(), ctxKeyOperator, params)
}

type AuditFields struct {
	CreateBy string `gorm:"size:100;comment:创建人;" json:"createBy"`
	UpdateBy string `gorm:"size:100;comment:修改人;"  json:"updateBy"`
	DeleteBy string `gorm:"size:100;comment:删除人;"  json:"deleteBy"`
}

type Page struct {
	Current uint64 `gorm:"-" json:"current" `
	Size    uint64 `gorm:"-" json:"size" `
	Total   uint64 `gorm:"-" json:"total" `
	Pages   uint64 `gorm:"-" json:"pages" `
}
