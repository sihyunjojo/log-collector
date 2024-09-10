package logwriter

import (
    "os"
)

// 공통적인 로그 저장 함수
func SaveLog(logData, logFilePath string) error {
    // 로그 폴더 경로 추출
    logFolder := "./logs"

    // 로그 폴더가 없으면 생성
    if _, err := os.Stat(logFolder); os.IsNotExist(err) {
        os.Mkdir(logFolder, 0755)
    }

    // 로그 파일에 데이터 기록 (추가 모드)
    f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer f.Close()

    // 로그 데이터를 파일에 저장
    if _, err := f.WriteString(logData); err != nil {
        return err
    }
    return nil
}
