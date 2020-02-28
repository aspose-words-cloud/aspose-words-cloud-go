properties([
	gitLabConnection('gitlab'),
	parameters([string(defaultValue: 'refs/heads/master', description: 'the branch to build', name: 'branch', trim: true)])
])

node('windows2019') {
	try {
		gitlabCommitStatus("checkout") {
			stage('checkout'){
				checkout([$class: 'GitSCM', branches: [[name: params.branch]], doGenerateSubmoduleConfigurations: false, extensions: [[$class: 'LocalBranch', localBranch: "**"]], submoduleCfg: [], userRemoteConfigs: [[credentialsId: '361885ba-9425-4230-950e-0af201d90547', url: 'https://git.auckland.dynabic.com/words-cloud/words-cloud-go.git']]])
			}
		}
		gitlabCommitStatus("tests") {
			stage('tests') {
				withCredentials([usernamePassword(credentialsId: '6839cbe8-39fa-40c0-86ce-90706f0bae5d', passwordVariable: 'AppKey', usernameVariable: 'AppSid')]) {
					bat "echo {\"AppSid\":\"%AppSid%\",\"AppKey\":\"%AppKey%\",\"BaseUrl\":\"https://api-qa.aspose.cloud/v4.0\" } > config.json"
				}
				bat 'docker run -v %cd%:c:/sdk -w="c:/sdk" --rm -t golang:1.14.0-windowsservercore-1809 go test ./... -v'
			}
		}
	} finally {
		bat 'docker system prune -f'
		deleteDir()
	}
}