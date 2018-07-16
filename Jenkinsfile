#!groovy

def buildId
def commitId
def commitIdSubString

node('on-demand') {
  currentBuild.result = "SUCCESS"
  try {

    stage 'setup'
    def workspace = pwd()
    withEnv(["GOPATH=${workspace}"]) {
      dir('src/github.com/anubhavmishra/hello-oscon') {

        checkout scm
        sh "git rev-parse HEAD > commitid"

        commitId = readFile('commitid')
        commitIdSubString = commitId.substring(0, 5)
        buildId = "${env.BUILD_NUMBER}-${commitIdSubString}"
        if (env.BRANCH_NAME == 'master') {
          currentBuild.displayName = buildId
        } else {
          currentBuild.displayName = "${buildId}-${env.BRANCH_NAME}"
        }
        echo "commitId=${commitId}\n buildId=${buildId}"

        stage 'go test'
        sh '''
          make test
        '''

        if (env.BRANCH_NAME == 'master') {
          stage 'build'
          sh '''
            make build-service
          '''
        }
      }
    }
  } catch (error) {
    currentBuild.result = "FAILURE"
    throw error
  }
}