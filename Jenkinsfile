pipeline {
  agent any
  stages {
    stage('start1') {
      agent {
        docker {
          image 'alpine:3.13.2'
        }

      }
      steps {
        echo 'test starting...'
        echo 'enter into func1'
      }
    }

  }
}