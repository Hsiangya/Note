package comm

import (
	"crypto/md5"
	"fmt"
	"io"
	"time"
)

const TimeLayout = "2006-01-02 15:04:05"

// Now time.Now
func Now() *time.Time {
	now := time.Now()
	return &now
}

// EncryptMd5 get md5 string
func EncryptMd5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// TimeFormat 时间转字符串
func TimeFormat(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format(TimeLayout)
}

// TimeParse 字符串转时间
func TimeParse(str string) *time.Time {
	if t, err := time.Parse(TimeLayout, str); err != nil {
		return nil
	} else {
		return &t
	}
}
