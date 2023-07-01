pipeline {
    agent any

    environment {
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }

    stages {
        stage('Git Checkout') {
            steps {
                echo 'Not needed cause git is already in settings'

                echo "Current Job Name2: ${JOB_NAME}"
                echo "Current Build ID2: ${BUILD_ID}"


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
