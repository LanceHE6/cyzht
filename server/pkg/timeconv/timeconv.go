package timeconv

import (
	"fmt"
	"time"
)

// ParesStrToTime 将时间字符串转为Time类型
func ParesStrToTime(timeStr string) (*time.Time, error) {
	// 定义时间的布局，必须与时间字符串的格式相匹配
	layout := "2006-01-02 15:04:05"
	// 使用 Parse 函数将时间字符串转换为 time.Time 类型
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		fmt.Println("时间转换错误：", err)
		return nil, err
	}
	return &t, nil
}
