properties([
	gitLabConnection('gitlab'),
	[$class: 'ParametersDefinitionProperty', 
		parameterDefinitions: [
			[$class: 'StringParameterDefinition', name: 'branch', defaultValue: 'master', description: 'the branch to build'],
			[$class: 'StringParameterDefinition', name: 'apiUrl', defaultValue: 'https://api-qa.aspose.cloud', description: 'api url'],
            [$class: 'BooleanParameterDefinition', name: 'ignoreCiSkip', defaultValue: false, description: 'ignore CI Skip'],
            [$class: 'StringParameterDefinition', name: 'credentialsId', defaultValue: '6839cbe8-39fa-40c0-86ce-90706f0bae5d', description: 'credentials id'],
            [$class: 'BooleanParameterDefinition', name: 'packageTesting', defaultValue: false, description: 'Testing package from repository without local sources. Used for prodhealthcheck'],
		]
	]
])

def needToBuild = false
def packageTesting = false
def currentFolder = "dev"

node('win2019') {
	try {
		gitlabCommitStatus("checkout") {
			stage('checkout'){
				checkout([$class: 'GitSCM', branches: [[name: params.branch]], doGenerateSubmoduleConfigurations: false, extensions: [[$class: 'LocalBranch', localBranch: "**"]], submoduleCfg: [], userRemoteConfigs: [[credentialsId: '361885ba-9425-4230-950e-0af201d90547', url: 'https://git.auckland.dynabic.com/words-cloud/words-cloud-go.git']]])

                bat 'git show -s HEAD > gitMessage'
                def commitMessage = readFile('gitMessage').trim()
                echo commitMessage
                needToBuild = params.ignoreCiSkip || !commitMessage.contains('[ci skip]')
                packageTesting = params.packageTesting
                bat 'git clean -fdx'
			}
		}
        
        if (needToBuild) {
            if (packageTesting) {
                gitlabCommitStatus("health check tests") {
                    stage('health check tests'){
                        currentFolder = powershell(returnStdout: true, script:"echo (Get-ChildItem -Path v* -Directory | Sort-Object -Property Name | Select -Last 1).Name").trim()
                        bat "Scripts\\RunHealthCheckTestsInDocker.bat ${currentFolder}"
                    }
                }
            }
            else {
                gitlabCommitStatus("tests") {
                    stage('tests') {
                        withCredentials([usernamePassword(credentialsId: params.credentialsId, passwordVariable: 'ClientSecret', usernameVariable: 'ClientId')]) {
                            try {
                                bat "Scripts\\RunTestsInDocker.bat"
                            } 
                            finally {
                                junit '**\\testReport.xml'
                            }
                        }
                    }
                }
            }
        }
	} finally {
		deleteDir( )
	}
}