def gitrepository = "${OPTION_GIT_REPOSITORY}"
def gitcredentials = "${OPTION_GIT_CREDENTIALS}"
def gitbranch = "${OPTION_GIT_BRANCH}"
def imagerepository = "https://${OPTION_NEXUS_REPOSITORY}/"
def repocredentials = "${OPTION_NEXUS_CREDENTIALS}"

pipeline {
    stage('checkout') {
        git credentialsId: gitcredentials, url: gitrepository, branch: gitbranch
    }

    stage('build') {
        docker.withRegistry(nexusrepository, nexuscredentials) {
            sh 'docker build -t "${OPTION_NEXUS_REPOSITORY}/${OPTION_IMAGE_NAME}:jenkins-${BUILD_NUMBER}" .'
        }
    }

    stage('promotion') {
        input "push image?"
    }

    stage('push') {
        docker.withRegistry(nexusrepository, nexuscredentials) {
            sh 'docker push "${OPTION_NEXUS_REPOSITORY}/${OPTION_IMAGE_NAME}:jenkins-${BUILD_NUMBER}"'
        }
    }
}
