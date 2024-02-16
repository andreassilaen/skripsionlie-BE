pipeline {
    agent {
        label 'kubernetes'
    }
    
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
                    def dockerImage = docker.build("skripsionline-be")
                    echo '> Configuring docker hub auth...'
                    withCredentials([string(credentialsId: 'DOCKER_HUB_TOKEN', variable: 'TOKEN')]){
                        sh 'docker login -u eximiusblitz -p $TOKEN'
                    }
                    echo '> Pushing image to docker hub...'
                    dockerImage.push()
                }
            }
        }  
        // stage('Push to registry') {
        //     steps {
        //         script{
        //             echo '> Creating image...'
        //             def dockerImage = docker.build("test-be")
        //         }
        //     }
        // }  
    }
}
