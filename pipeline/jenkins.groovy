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
                sh """
                    echo "Simulating docker build for tag: ${params.IMAGE_TAG}"
                    echo "Here we would run: make docker-build"
                """
            }
        }

        stage('Push Image') {
            steps {
                sh """
                    echo "Simulating docker push to ghcr.io/samtishka/telegram-bot:${params.IMAGE_TAG}"
                    echo "Here we would run: docker login && docker push"
                """
            }
        }
    }

    post {
        always {
            echo "Pipeline finished"
        }
    }
}
