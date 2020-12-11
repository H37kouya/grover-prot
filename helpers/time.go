package helpers

import "time"

// GetTimeForFilename ファイル名のために時間を取得する
func GetTimeForFilename() string {
	t := time.Now()
	layout := "2006-01-02_15.04.05"
	return t.Format(layout)
}
