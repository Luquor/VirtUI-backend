pipeline {
  // Run on an agent where we want to use Go
  agent any
git branch: 'main', credentialsId: '085ce31d-83d4-4a37-9ad2-6bbfd657cd26', url: 'https://github.com/Luquor/VirtUI-backend.git'
agent any
  // Ensure the desired Go version is installed for all stages,
  // using the name defined in the Global Tool Configuration
  tools { go '1.19' }

  stages {
    stage('Build') {
      steps {
        // Output will be something like "go version go1.19 darwin/arm64"
        sh 'go version'
      }
    }
  }
}

