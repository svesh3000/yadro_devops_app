/* groovylint-disable-next-line CompileStatic */

pipeline {
  agent any

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
}
