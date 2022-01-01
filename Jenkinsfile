pipeline {
  agent any

  stages {
    stage('Check Tools') {
      steps {
        sh '''
        id
        pwd
        git --version
        docker version
        go version
        python3 --version
        '''
      }
    }

    stage('Test') {
      steps {
        sh './mud_docker.sh test'
      }
    }

    stage('Docker Build') {
      steps {
        sh './mud_docker.sh build'
      }
    }

    stage('Deploy') {
      steps {
        sh './mud_docker.sh up'
      }
    }
  }
}
