pipeline {
    agent any

    environment {
        DOCKER_IMAGE = "log-collector-app"
        DOCKER_TAG = "latest"
    }

    stages {
        stage('Checkout') {
            steps {
                git(
                    branch: 'main', // 브랜치 설정
                    url: 'https://github.com/TEAM-Joyride/logCollector.git', // 리포지토리 URL
                    credentialsId: '404e752c-fb33-4398-969b-11253d1777d1' // (옵션) 인증이 필요한 경우 자격 증명 설정
                )
            }
        }

        stage('Check') {
            steps {
                sh 'docker ps'
                sh 'ls /var/jenkins_home/workspace/log-server/docs'
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    // Docker 이미지 빌드
                    // ${DOCKER_IMAGE}:${DOCKER_TAG}: 빌드된 Docker 이미지의 이름과 태그를 지정합니다. 예를 들어, log-collector-app:latest 같은 형식입니다.
                    // .: 현재 디렉터리에서 Dockerfile을 찾고, 그 디렉터리의 모든 파일을 빌드 컨텍스트로 사용한다는 의미입니다.
                    sh 'docker build -t ${DOCKER_IMAGE}:${DOCKER_TAG} .'

                    // docker build -t ${DOCKER_IMAGE}:${DOCKER_TAG} -f docker/Dockerfile .
                        // docker/Dockerfile: docker/ 폴더 안에 있는 Dockerfile을 사용합니다.
                        // .: 빌드 컨텍스트는 현재 디렉터리 (.)로, 현재 디렉터리에 있는 모든 파일을 빌드 컨텍스트로 사용합니다.
                }
            }
        }

        stage('Run Tests') {
            steps {
                script {
                    sh 'ls /var/jenkins_home/workspace/log-server/docs'
                    // 테스트 실행
//                     sh 'go test /var/jenkins_home/workspace/log-server/...'
                    sh 'go test ./...'
                }
            }
        }

        stage('Run Docker Container') {
            steps {
                script {
                    // 기존 컨테이너가 있으면 삭제
                    sh 'docker rm -f log-collector || true'

                    // Docker 컨테이너 실행
                    sh 'docker run -d --name log-collector -p 8080:8080 ${DOCKER_IMAGE}:${DOCKER_TAG}'
                }
            }
        }
    }

    post {
        always {
            // 빌드 결과 확인
            sh 'docker ps -a'
        }
    }
}
