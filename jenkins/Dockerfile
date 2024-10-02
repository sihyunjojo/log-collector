# Jenkins LTS 이미지를 베이스로 사용
FROM jenkins/jenkins:lts

# Docker를 Jenkins 안에서 사용할 수 있도록 설치
USER root
RUN apt-get update && \
		apt-get install -y wget && \
    apt-get install -y docker.io && \
    apt-get clean && \
    wget https://go.dev/dl/go1.23.1.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.23.1.linux-amd64.tar.gz && \
    rm go1.23.1.linux-amd64.tar.gz

# 환경 변수 설정
ENV PATH="/usr/local/go/bin:${PATH}"

# Jenkins 실행 시 사용하는 SSH 디렉토리 설정
# (컴포즈에서는 /var/home/jenkins_ssh를 마운트했으므로 해당 부분은 컨테이너 실행 시 처리해야 함)
USER jenkins