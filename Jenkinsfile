/* groovylint-disable-next-line CompileStatic */

pipeline {
  agent { label 'worker' }

  environment {
    DOCKER_IMAGE = 'm.sveshnikov/weather-app'
    DOCKER_TAG = "${BUILD_NUMBER}"
    DOCKER_CREDENTIALS = credentials('docker-hub-credentials')

    PORT = '8000'
    VERSION = '1.0.0'
    AUTHOR = 'm.sveshnikov1'
    API_KEY = credentials('WEATHER_API_KEY')
  }

  stages {
    stage('Lint') {
      steps {
        echo 'lint'
      }
    }
    stage('Test') {
      steps {
        echo 'test'
      }
    }
    stage('Build') {
      steps {
        echo 'build'
      }
    }
    stage('Deploy') {
      when {
        branch 'master'
      }
      steps {
        echo 'deploy'
      }
    }
  }

  post {
    always {
        echo 'Pipeline finished!'
    }
    success {
        echo 'All stages passed!'
    }
    failure {
        echo 'ERROR!'
    }
}
}
