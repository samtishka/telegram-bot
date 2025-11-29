pipeline {
    agent any

    parameters {
        string(name: 'IMAGE_TAG', defaultValue: 'v1.0.0-linux-amd64', description: 'Docker image tag')
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build Image') {
            steps {
                sh 'make docker-build'
            }
        }

        stage('Push Image') {
            steps {
                withCredentials([string(credentialsId: 'ghcr_token', variable: 'TOKEN')]) {
                    sh """
                        echo \$TOKEN | docker login ghcr.io -u samtishka --password-stdin

                        docker tag samtishka/telegram-bot:${params.IMAGE_TAG} ghcr.io/samtishka/telegram-bot:${params.IMAGE_TAG}
                        docker push ghcr.io/samtishka/telegram-bot:${params.IMAGE_TAG}
                    """
                }
            }
        }
    }

    post {
        always {
            echo "Pipeline finished"
        }
    }
}
