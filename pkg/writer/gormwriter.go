package writer

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"reflect"
)

type GormWriter struct {
	db    *gorm.DB
	model reflect.Type
}

func NewGormWriter(db *gorm.DB, logInstance reflect.Type) *GormWriter {
	return &GormWriter{
		db:    db,
		model: logInstance,
	}
}

func (gw *GormWriter) Write(p []byte) (n int, err error) {
	logInstance := reflect.New(gw.model.Elem()).Interface()
	if err = json.Unmarshal(p, logInstance); err != nil {
		return 0, fmt.Errorf("log日志格式转换失败: %+v", err)
	}

	if err = gw.db.Model(&logInstance).Create(logInstance).Error; err != nil {
		return 0, fmt.Errorf("log日志写入失败: %+v", err)
	}

	return len(p), nil
}
