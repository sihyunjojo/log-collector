# 1. Build 단계
FROM golang:1.23-alpine AS builder

# 시간대 설정을 위한 패키지 설치
RUN apk add --no-cache tzdata

# Asia/Seoul 시간대로 설정
RUN cp /usr/share/zoneinfo/Asia/Seoul /etc/localtime && echo "Asia/Seoul" > /etc/timezone

# 필요한 패키지 설치
RUN apk add --no-cache build-base git

# 변경된 시간대 확인
RUN date

# 2. 작업 디렉터리 설정
WORKDIR /app

# 3. 모듈 다운로드 (go.mod, go.sum을 먼저 복사하여 캐시 활용)
COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy

# 4. Swag 설치 (Swagger 문서 생성기)
RUN go install github.com/swaggo/swag/cmd/swag@v1.16.3

# 5. 애플리케이션 소스 코드 복사
COPY . .

# 6. Swag init 실행 (Swagger 문서 생성)
RUN swag init

# 7. 애플리케이션 빌드
# 현재 작업 디렉터리 (/app) 내에서 Go 애플리케이션을 빌드하고, 빌드된 실행 파일을 main이라는 이름으로 생성합니다.
RUN go build -o main .

# 8. 실행 단계 (alpine 이미지 사용)
FROM alpine
#FROM scratch # 운영체제도, 쉘 명령어도, 유틸리티도 포함되지 않은 매우 경량화된 이미지입니다.

# Filebeat의 autodiscover 기능을 위한 레이블 추가
# 이 레이블들은 Filebeat가 이 컨테이너의 로그를 자동으로 수집하도록 지시합니다.
 # 이 컨테이너의 로그 수집을 활성화합니다.
LABEL co.elastic.logs/enabled="true"
# Filebeat의 Golang 모듈을 사용하여 로그를 파싱합니다.
LABEL co.elastic.logs/module="golang"

# ... (이하 내용 동일)
# tzdata 설치 (최종 실행 이미지에서도 필요)
############# 이걸 2번해줘야되드라?
RUN apk add --no-cache tzdata

# Asia/Seoul 시간대 설정
RUN cp /usr/share/zoneinfo/Asia/Seoul /etc/localtime && echo "Asia/Seoul" > /etc/timezone

# 빌드된 실행 파일 복사
COPY --from=builder /app/main /main

# .env 파일 복사 (필요한 경우)
COPY --from=builder /app/.env ./

RUN ls -l /main # 파일이 존재하는지 확인

# 11. 포트 설정
EXPOSE 8085

# 12. 실행
CMD ["/main"]


## 1. Build 단계
#FROM golang:1.23-alpine AS builder
#
## 필요한 패키지 설치
#RUN apk add --no-cache build-base git
#
## 2. 작업 디렉터리 설정
#WORKDIR /app
#
## 3. 모듈 다운로드
#COPY go.mod go.sum ./
#RUN go mod download
#RUN go mod tidy
#
## 4. Swag 설치
#RUN go install github.com/swaggo/swag/cmd/swag@v1.16.3
#
## 5. 애플리케이션 소스 코드 복사
#COPY . .
#
## 6. Swag init 실행
#RUN swag init
#
## 7. 포트 설정
#EXPOSE 8089
#
## 8. 실행
#CMD ["./main.exe"]


## 8. 실행 단계 (더 경량화된 scratch 이미지 사용)
#
## 9. 빌드된 실행 파일 복사
#COPY --from=builder /app/main ./main
#
## 10. .env 파일 복사 (필요한 경우)
#COPY --from=builder /app/.env .
#
#RUN ls -l ./main
#
## 11. 포트 설정
#EXPOSE 8089
#
## 12. 실행
#CMD ["./main"]
