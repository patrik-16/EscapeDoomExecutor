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
                echo 'Deploy...'

                sh 'chmod +x ./deployz1.sh'
                sh 'chmod +x ./deployz2.sh'

                sh './deployz1.sh'
                script {
                    sleep 5
                }

                sh './deployz2.sh &'
                script {
                    sleep 120
                }
                echo 'Done!'
            }
        }
    }
}
