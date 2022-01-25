pipeline {
  agent any
  
  stages {
    stage('Pre Check') {
      steps {
        sh '''
        id
        pwd
        git --version
        docker version
        '''
      }
    }

    stage('Test') {
      steps {
        sh '''
        set +x; source ./env/env.docker.sh; set -x
        export BACK_TARGET=test
        export BACK_IMAGE=${BACK_IMAGE}_test
        docker-compose build backend
        docker-compose up -d database
        docker-compose run backend
        docker-compose down || true
        '''
      }
    }

    stage('Docker Build') {
      steps {
        sh '''
        set +x; source ./env/env.docker.sh; set -x
        docker-compose build
        '''
      }
    }

    stage('Deploy') {
      steps {
        sh '''
        set +x; source ./env/env.docker.sh; set -x
        docker-compose down || true
        docker-compose up -d
        '''
      }
    }
  }
}
