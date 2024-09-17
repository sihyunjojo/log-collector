package config

import (
	"time"
)

var SeoulLocation *time.Location

// GetSeoulTime은 현재 시간을 Asia/Seoul 시간대로 반환
func GetSeoulTime() time.Time {
	location, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		location = time.FixedZone("KST", 9*60*60)
	}
	return time.Now().In(location)
}

// GetNextMidnightInSeoul은 Asia/Seoul 시간대의 자정 시간을 반환
func GetNextMidnightInSeoul() time.Time {
	now := GetSeoulTime()

	// 다음 자정 시간을 계산 (현재 날짜의 자정)
	next := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	return next
}

// GetSeoulTimeYesterday는 어제 날짜를 Asia/Seoul 시간대로 반환
func GetSeoulTimeYesterday() time.Time {
	location, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		location = time.FixedZone("KST", 9*60*60)
	}
	// 하루(-1)를 더해 어제 날짜로 설정
	return time.Now().In(location).AddDate(0, 0, -1)
}
