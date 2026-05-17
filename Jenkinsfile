@Library('sharedLib') _

pipeline {
  agent none

  environment {
    DOCKER_IMAGE = 'sveshnikov/app'
    DOCKER_TAG = "${BUILD_NUMBER}"
    DOCKER_CREDENTIALS = credentials('docker-hub-credentials')
    API_KEY = credentials('WEATHER_API_KEY')
  }

  stages {
    stage('Checkout') {
      agent { label 'staging' }
      steps { checkout scm }
    }
    stage('GitLab Status') {
      agent { label 'staging' }
      steps { updateGitlabCommitStatus name: 'pipeline', state: 'pending' }
    }
    stage('Check') {
      parallel {
        stage('Lint') {
          agent { label 'staging' }
          steps { script { sharedLib.lint() } }
        }
        stage('SAST') {
          agent { label 'staging' }
          steps { script { sharedLib.sast() } }
        }
        stage('Test') {
          agent { label 'staging' }
          steps { script { sharedLib.test() } }
        }
      }
    }
    stage('Docker Login') {
      agent { label 'staging' }
      steps {
        script {
          sh 'echo $DOCKER_CREDENTIALS_PSW | docker login -u $DOCKER_CREDENTIALS_USR --password-stdin'
        }
      }
    }
    stage('Build') {
      when { expression { changeRequest() || env.BRANCH_NAME == 'master' || env.TAG_NAME =~ /v.*/} }
      agent { label 'staging' }
      steps { script { sharedLib.build() } }
    }
    stage('Push') {
      when { expression {  env.BRANCH_NAME == 'master' || env.TAG_NAME =~ /v.*/ } }
      agent { label 'staging' }
      steps { script { sharedLib.push() } }
    }
    stage('Deploy Staging') {
      when { expression { env.BRANCH_NAME == 'master' } }
      agent { label 'staging' }
      steps { script { sharedLib.deploy(imageTag: env.DOCKER_TAG, environment: 'staging') } }
    }
    stage('Deploy Production') {
      when { expression { env.TAG_NAME =~ /v.*/ } }
      agent { label 'production' }
      steps { script { sharedLib.deploy(imageTag: env.DOCKER_TAG, environment: 'production') } }
    }
  }

  post {
    always {
      echo 'Pipeline finished!'
      node('staging') {
          cleanWs()
          sh 'docker system prune -a -f --volumes || true'
      }
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
