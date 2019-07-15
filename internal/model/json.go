package model

import (
	"strconv"
	"time"
)

// TimeImplementedMarshaler 是自定义日期类型，用于格式化日期
type TimeImplementedMarshaler time.Time

// MarshalJSON json自定义时间转换
func (obj TimeImplementedMarshaler) MarshalJSON() ([]byte, error) {
	seconds := time.Time(obj).Unix()
	return []byte(strconv.FormatInt(seconds, 10)), nil
}
