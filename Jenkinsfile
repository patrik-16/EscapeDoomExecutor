1. stage
  pull git

2. stage
  build docker container

3. stage
  deploy


pipeline {
    agent any

    stages {
        stage('Git Checkout') {
            steps {
                echo 'Not needed cause git is already in settings'
            }
        }

        stage('Build') {
            steps {
                echo 'Building...'
                sh 'go build'
            }
        }



    }
}
