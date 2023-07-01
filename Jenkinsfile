pipeline {
    agent any

    tools {
        go 'go-1.20.4'
    }

    stages {
        stage('Git Checkout') {
            steps {
                echo 'Not needed cause git is already in settings'

                echo "Current Job Name: ${JOB_NAME}"
                echo "Current Build ID: ${BUILD_ID}"
                echo "Current Path: ${PATH}"

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
