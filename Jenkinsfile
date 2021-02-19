pipeline {
  agent any
  stages {
    stage('start1') {
      parallel {
        stage('start1') {
          agent any
          steps {
            echo 'test starting...'
            echo 'enter into func1'
          }
        }

        stage('build') {
          steps {
            echo 'building...'
          }
        }

        stage('end') {
          steps {
            echo 'ending...'
          }
        }

      }
    }

  }
}