pipeline {
    agent any
    
    stages {
        stage('Build') {
            steps {
                // Build the Go project
                sh 'make build'
            }
        }
        
        stage('Dockerize') {
            steps {
                script{
                    echo '> Creating image...'
                    def dockerImage = docker.build("test-be")
                }
            }
        }  
    }
}
