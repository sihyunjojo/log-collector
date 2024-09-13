package config

import (
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
	"strconv"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

// SetupLogger는 환경 변수를 사용하여 로그 파일 설정을 반환합니다.
func SetupLogger(folderName string, fileName string) *lumberjack.Logger {
	// 환경 변수 읽기
	baseDirectory := os.Getenv("LOG_DIRECTORY") // 기본 로그 디렉토리
	logFileName := os.Getenv("LOG_FILE_NAME")
	logMaxSize, _ := strconv.Atoi(os.Getenv("LOG_MAX_SIZE"))
	logMaxBackups, _ := strconv.Atoi(os.Getenv("LOG_MAX_BACKUPS"))
	logMaxAge, _ := strconv.Atoi(os.Getenv("LOG_MAX_AGE"))
	logCompress, _ := strconv.ParseBool(os.Getenv("LOG_COMPRESS"))

	// 이 코드는 logDirectory라는 변수를 선언하고, baseDirectory의 값을 logDirectory에 복사하는 작업을 수행합니다.
	logDirectory := baseDirectory
	if folderName != "" {
		logDirectory = baseDirectory + "/" + folderName
		// 폴더가 존재하지 않으면 생성
		if _, err := os.Stat(logDirectory); os.IsNotExist(err) {
			err := os.MkdirAll(logDirectory, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}
	}

	// 파일명 처리
	if fileName != "" {
		logFileName = fileName
	}

	// 로그 설정 반환
	return &lumberjack.Logger{
		Filename:   logDirectory + "/" + logFileName,
		MaxSize:    logMaxSize,
		MaxBackups: logMaxBackups,
		MaxAge:     logMaxAge,
		Compress:   logCompress,
	}
}

// GetLogger는 writer에서 설정한 로그 작성자 (Fiber의 로거)를 반환함
func GetLogger() logger.Config {
	userLogger := SetupLogger("status", "status")
	seoulLocation, _ := time.LoadLocation("Asia/Seoul")

	return logger.Config{
		Format:     "[${time}] ${status} - ${method} ${path} ${latency} - ${ip}\n",
		Output:     userLogger,
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   seoulLocation.String(),
		TimeFunc: func() time.Time {
			return time.Now().In(SeoulLocation)
		},
	}
}
