package config

import (
	"fmt"
	"time"
)

var SeoulLocation *time.Location

// GetSeoulTime은 현재 시간을 Asia/Seoul 시간대로 반환
func GetSeoulTime() time.Time {
	location, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		// 시간대 로드 실패 시 기본 시간대 사용
		return time.Now()
	}
	// 현재 시간을 Asia/Seoul 시간대로 변환하여 반환
	return time.Now().In(location)
}

// GetNextMidnightInSeoul은 Asia/Seoul 시간대의 자정 시간을 반환
func GetNextMidnightInSeoul() time.Time {
	now := GetSeoulTime()
	next := now.Add(time.Hour * 24)
	return time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, now.Location())
}

func init() {
	var err error
	SeoulLocation, err = time.LoadLocation("Asia/Seoul")
	if err != nil {
		fmt.Printf("서울 시간대를 로드하는데 실패했습니다: %v\n", err)
		SeoulLocation = time.FixedZone("KST", 9*60*60) // 실패 시 KST로 대체
	}
}
