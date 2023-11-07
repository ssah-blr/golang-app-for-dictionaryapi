pipeline {
    agent any

    tools { go '1.21.3' }

    environment {
        DOCKER_IMAGE = 'ssahblr/go-api:latest'
        DOCKERHUB_CREDENTIALS = credentials('ssah-blr-docker')
    }

    stages {

        stage('Checkout') {
            steps {
                sh 'echo "For Checkout"'
            }
        }

        stage('Build Go Application') {
            steps {
                dir('go') {
                    sh 'go mod init test'
                    sh 'go get -u github.com/gin-gonic/gin'
                    sh ''
                }
            }
        }

        stage('Unit Tests') {
            steps {
                dir('go') {
                    sh 'go test'
                }
            }
        }

        stage('Build Docker Image') {
            steps {
                dir('go') {
                    sh 'docker build -t $DOCKER_IMAGE .'
                }
            }
        }

        stage('Push Docker Image to Docker Hub') {
            steps {
                sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'
                sh 'docker push $DOCKER_IMAGE'
            }
        }
    }
}
