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
    stage('GitLab Status') {
      steps {
        updateGitlabCommitStatus name: 'pipeline', state: 'pending'
      }
    }
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
        updateGitlabCommitStatus name: 'pipeline', state: 'success'
        echo 'All stages passed!'
    }
    failure {
        updateGitlabCommitStatus name: 'pipeline', state: 'failed'
        echo 'ERROR!'
    }
  }
}
