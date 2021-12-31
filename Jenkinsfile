pipeline {
  agent any

  stages {
    stage('Check Tools') {
      steps {
        sh '''
        git --version
        docker version
        go version
        node -v
        npm -v
        python3 --version
        '''
      }
    }

    stage('Build') {
      steps {
        echo 'Hello World'
      }
    }

    stage('Test') {
      steps {
        echo 'Hello World'
      }
    }

    stage('Deploy') {
      steps {
        echo 'Hello World'
      }
    }
  }
}
