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

    stage('Pre Work') {
      steps {
        sh 'chmod +x ./mud_docker.sh'
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
