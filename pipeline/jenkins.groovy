pipeline {
    agent any

    parameters {
        choice(
            name: 'PLATFORM',
            choices: ['linux/amd64', 'linux/arm64'],
            description: 'Цільова платформа для збірки контейнера'
        )
        string(
            name: 'IMAGE_TAG',
            defaultValue: 'v1.0.0-linux-amd64',
            description: 'Тег Docker-образу'
        )
    }

    stages {
        stage('Checkout') {
            steps {
                git branch: 'main',
                    url: 'https://github.com/samtishka/telegram-bot.git'
            }
        }

        stage('Build Image') {
            steps {
                script {
                    sh """
                      docker buildx build --platform \${params.PLATFORM} \
                        -t samtishka/telegram-bot:\${params.IMAGE_TAG} \
                        --load .
                    """
                }
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
