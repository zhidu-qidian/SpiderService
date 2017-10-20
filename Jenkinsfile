pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                sh 'go version'
                sh 'go build'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}