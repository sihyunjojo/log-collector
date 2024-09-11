package config

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"time"
)

// RotateLogger는 자정마다 새 로그 파일로 롤링하는 작업을 수행
func RotateLogger(logger *lumberjack.Logger, folderName string) {
	go func() {
		for {
			// 현재 시간
			now := time.Now()

			// 자정 시간 설정
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())

			// 자정까지 대기
			time.Sleep(time.Until(next))

			// 새 로그 파일 이름 (기본 파일명 + 날짜)
			baseDirectory := os.Getenv("LOG_DIRECTORY")
			if folderName != "" {
				baseDirectory = filepath.Join(baseDirectory, folderName)
			}
			newLogFile := filepath.Join(baseDirectory, "app-"+time.Now().Format("2006-01-02")+".log")

			// 롤링 수행
			logger.Rotate()
			logger.Filename = newLogFile
		}
	}()
}
