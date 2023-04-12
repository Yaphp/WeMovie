package model

import (
	"database/sql/driver"
	"strings"
	"time"
)

type LocalTime string

// Scan
//
//	@Description: 自定义从数据库读出来数据解析
//	@receiver receiver
//	@param value
//	@return error
func (receiver *LocalTime) Scan(value interface{}) error {
	timeValue := value.(time.Time)
	*receiver = LocalTime(timeValue.Format("2006-01-02 15:04:05"))
	return nil
}

// Value
//
//	@Description: 自定义存入数据库数据解析
//	@receiver receiver
//	@return driver.Value
//	@return error
func (receiver LocalTime) Value() (driver.Value, error) {
	if len(receiver) == 0 {
		return nil, nil
	}
	return string(receiver), nil
}

// LocalTime4Millisecond  时间戳带毫秒
type LocalTime4Millisecond string

// Scan
//
//	@Description: 自定义从数据库读出来数据解析
//	@receiver receiver
//	@param value
//	@return error
func (receiver *LocalTime4Millisecond) Scan(value interface{}) error {
	timeValue := value.(time.Time)
	*receiver = LocalTime4Millisecond(strings.Replace(timeValue.Format("2006-01-02 15:04:05.000"), ".", ":", 1))
	return nil
}

// Value
//
//	@Description: 自定义存入数据库数据解析
//	@receiver receiver
//	@return driver.Value
//	@return error
func (receiver LocalTime4Millisecond) Value() (driver.Value, error) {
	if len(receiver) == 0 {
		return nil, nil
	}
	return string(receiver), nil
}
