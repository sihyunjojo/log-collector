# 1. Build 단계
FROM golang:1.23-alpine AS builder

# 필요한 패키지 설치
RUN apk add --no-cache build-base git

# 2. 작업 디렉터리 설정
WORKDIR /app

# 3. 모듈 다운로드
COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy

# 4. Swag 설치
RUN go install github.com/swaggo/swag/cmd/swag@v1.16.3

# 5. 애플리케이션 소스 코드 복사
COPY . .

# 6. Swag init 실행
RUN swag init

# 7. 포트 설정
EXPOSE 8089

# 8. 실행
CMD ["go", "run", "main.go"]


## 1. Build 단계
#FROM golang:1.23-alpine AS builder
#
## 필요한 패키지 설치
#RUN apk add --no-cache build-base git
#
## 2. 작업 디렉터리 설정
#WORKDIR /app
#
## 3. 모듈 다운로드 (go.mod, go.sum을 먼저 복사하여 캐시 활용)
#COPY go.mod go.sum ./
#RUN go mod download
#RUN go mod tidy
#
## 4. Swag 설치 (Swagger 문서 생성기)
#RUN go install github.com/swaggo/swag/cmd/swag@v1.16.3
#
## 5. 애플리케이션 소스 코드 복사
#COPY . .
#
## 6. Swag init 실행 (Swagger 문서 생성)
#RUN swag init
#
## 7. 애플리케이션 빌드
## 현재 작업 디렉터리 (/app) 내에서 Go 애플리케이션을 빌드하고, 빌드된 실행 파일을 main이라는 이름으로 생성합니다.
#RUN go build -o main .
#
## 8. 실행 단계 (더 경량화된 scratch 이미지 사용)
#FROM scratch
#
## 9. 빌드된 실행 파일 복사
#COPY --from=builder /app/main /main
#
## 10. .env 파일 복사 (필요한 경우)
#COPY --from=builder /app/.env .
#
## 11. 포트 설정
#EXPOSE 8089
#
## 12. 실행
#CMD ["/main"]
