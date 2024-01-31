pipeline {
    agent any
    
    stages {
        stage('Build') {
            steps {
                // Build the Go project
                sh 'go build -o myapp'
            }
        }
        
        stage('Dockerize') {
            steps {
                // Build the Docker image
                sh 'docker build -t myapp -f Dockerfile .'
            }
        }
        
        stage('Push to Registry') {
            steps {
                // Push the Docker image to a registry
                sh 'docker push myapp'
            }
        }
        
    }
}
