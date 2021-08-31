#!groovy

pipeline {
    agent any
    environment {
        registry = "kostiakorzh/demoshop-order-service-dev"
        registryCredential = 'dockerhub'
        dockerImage = ''
    }
    stages {
        stage('Build') {
            steps {
                script {
                    dockerImage = docker.build registry
                }
            }
        }
        stage('Push to dockerhub') {
            steps {
                script {
                    docker.withRegistry('', registryCredential) {
                        dockerImage.push("latest")
                    }
                }
            }
        }
        stage("Deploy to docker") {
            steps {
                sh 'docker rm -f order-dev-container'
                sh 'docker rmi kostiakorzh/demoshop-order-service-dev'
                sh 'docker run -p 1313:1313 -d --name order-dev-container -e MYSQL_URL=18.192.254.230 -e MYSQL_ROOT_PASSWORD=root1234 -e USER_SERVICE_URL=http://localhost:8082 kostiakorzh/demoshop-order-service-dev'
            }
        }
    }
}