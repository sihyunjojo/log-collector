package writer

import (
	"log"  // 기본 로깅 기능을 제공하는 패키지
	"time" // 시간과 날짜 관련 기능을 제공하는 패키지
)

// ScheduleLogRotation 자정마다 로그 파일을 변경하는 함수
func ScheduleLogRotation() {
	// 매일 자정에 로그 파일을 변경하도록 스케줄링
	// 무한 루프를 사용하여 매일 자정에 작업을 반복
	for {
		// 현재 시간
		// 현재 시간을 가져옴
		now := time.Now()

		// 다음 자정까지 남은 시간 계산
		// 현재 시간에서 다음 자정 시간(00:00:00)을 계산
		nextMidnight := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())

		// 자정까지 남은 시간을 계산
		durationUntilMidnight := nextMidnight.Sub(now)

		// 자정까지 대기
		// 자정까지 남은 시간 동안 대기 (Sleep)
		time.Sleep(durationUntilMidnight)

		// 자정에 파일 이름 변경
		// 자정이 되면 로그 파일을 어제 날짜로 변경하는 함수를 호출
		err := MoveLogAtMidnight("watch_log.txt")
		if err != nil {
			// 파일 이름 변경 중 오류가 발생하면 로그에 출력
			log.Printf("Error rotating log file: %v", err)
		}
	}
}
