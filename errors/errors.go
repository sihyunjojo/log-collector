package customerrors

import (
	"fmt"
)

// PortNotFoundError 정의
type PortNotFoundError struct {
	EnvVar string
}

// 에러 메시지를 반환하는 메서드
func (e *PortNotFoundError) Error() string {
	return fmt.Sprintf("%s 환경변수가 존재하지 않습니다.", e.EnvVar)
}
