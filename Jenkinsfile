pipeline {
    agent any

    environment {
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }

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
