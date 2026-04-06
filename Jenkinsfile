pipeline {
  agent { label 'worker' }

  environment {
    DOCKER_IMAGE = 'm.sveshnikov/weather-app'
    DOCKER_TAG = "${BUILD_NUMBER}"
    DOCKER_CREDENTIALS = credentials('docker-hub-credentials')
    API_KEY = credentials('WEATHER_API_KEY')
  }

  stages {
    stage('GitLab Status') {
      steps {
        updateGitlabCommitStatus name: 'pipeline', state: 'pending'
      }
    }
    stage('Check') {
      parallel {
        stage('Lint') {
          steps {
            sh '''
              docker run --rm -v $PWD:/app -w /app golang:1.21 sh -c "go vet ./..."
            '''
          }
        }
        stage('Test') {
          steps {
            sh '''
              docker run --rm -e API_KEY=$API_KEY -v $PWD:/app -w /app golang:1.21 \
              sh -c "go test -v ./..."
            '''
          }
        }
      }
    }
    stage('Build') {
      steps {
        sh "docker build -t ${DOCKER_IMAGE}:${DOCKER_TAG} ."
        sh "docker tag ${DOCKER_IMAGE}:${DOCKER_TAG} ${DOCKER_IMAGE}:latest"
      }
    }
    stage('Deploy') {
      when {
        branch 'master'
      }
      steps {
        sh '''
          echo $DOCKER_CREDENTIALS_PSW | docker login -u $DOCKER_CREDENTIALS_USR --password-stdin
          docker push ${DOCKER_IMAGE}:${DOCKER_TAG}
          docker push ${DOCKER_IMAGE}:latest
        '''
      }
    }
  }

  post {
    always {
        echo 'Pipeline finished!'
    }
    success {
        updateGitlabCommitStatus name: 'pipeline', state: 'success'
        echo 'All stages passed!'
    }
    failure {
        updateGitlabCommitStatus name: 'pipeline', state: 'failed'
        echo 'ERROR!'
    }
  }
}
