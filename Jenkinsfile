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
        echo ${PATH} | sed 's/:/:\\n/g'
        '''
      }
    }

    stage('Test') {
      steps {
        sh '''
        set +x; source ./env/env.docker.sh
        export BACK_TARGET=test
        export BACK_IMAGE=${BACK_IMAGE}_test
        echo ${BACK_TARGET}
        echo ${BACK_IMAGE}
        docker-compose build backend
        docker run ${BACK_IMAGE}:${BACK_TAG}
        '''
      }
    }

    stage('Docker Build') {
      steps {
        sh '''
        source ./env/env.docker.sh
        echo ${BACK_TARGET}
        echo ${BACK_IMAGE}
        '''
      }
    }

    stage('Deploy') {
      steps {
        sh '''
        source ./env/env.docker.sh
        '''
      }
    }
  }
}
