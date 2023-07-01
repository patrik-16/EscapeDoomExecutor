pipeline {
    agent any

    stages {
        stage('Git Checkout') {
            steps {
                echo 'Not needed cause git is already in settings'
                sh 'go version'
            }
        }

        stage('Build') {
            steps {
                echo 'Building...'
                sh 'go build'
                sh 'go version'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying...'

                echo 'Done!'
            }
        }
    }
}
