package timeconv

import (
	"fmt"
	"time"
)

// ParesStrToTime 将时间字符串转为Time类型
func ParesStrToTime(timeStr string) (*time.Time, error) {
	// 定义支持的两种时间格式
	layouts := []string{
		"2006-01-02 15:04:05", // 格式1: "2024-11-26 20:56:29"
		"2006-01-02T15:04:05", // 格式2: "2025-03-18T11:45:03"
	}

	// 尝试用每种格式解析时间字符串
	for _, layout := range layouts {
		t, err := time.Parse(layout, timeStr)
		if err == nil {
			return &t, nil // 解析成功，返回时间
		}
	}

	// 如果所有格式都解析失败，返回错误
	return nil, fmt.Errorf("时间格式不支持: %s", timeStr)
}
