# 프로젝트 기술 사항
- log 수집
- swagger 사용

# 사용법

## 초기화
```bash
go mod init log-collector    # 프로젝트에 Go 모듈 초기화
go mod tidy                 # 필요한 의존성을 자동으로 다운로드
```

## 실행
```bash
swag init
go run main.go
```

### 종료
```plantext
ctrl + c
```

## 테스트
```bash
go test ./...
```


## swag 문서 생성
주석을 작성한 후 다음 명령어를 실행하여 Swagger 문서를 생성합니다.

```bash
swag init
```

# [swag](http://localhost:8089/swagger/index.html)