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

node('sdk-linux') {
	try {
			stage('checkout'){
				checkout([$class: 'GitSCM', branches: [[name: params.branch]], doGenerateSubmoduleConfigurations: false, extensions: [[$class: 'LocalBranch', localBranch: "**"]], submoduleCfg: [], userRemoteConfigs: [[credentialsId: '361885ba-9425-4230-950e-0af201d90547', url: 'https://git.auckland.dynabic.com/words-cloud/words-cloud-go.git']]])

                sh 'git show -s HEAD > gitMessage'
                def commitMessage = readFile('gitMessage').trim()
                echo commitMessage
                needToBuild = params.ignoreCiSkip || !commitMessage.contains('[ci skip]')
                packageTesting = params.packageTesting
                sh 'git clean -fdx'
			}
        
        if (needToBuild) {
            withCredentials([usernamePassword(credentialsId: params.credentialsId, passwordVariable: 'ClientSecret', usernameVariable: 'ClientId')]) {
                if (packageTesting) {
                    stage('health check tests') {
                        try {
                            currentFolder = sh(returnStdout: true, script: "find . -maxdepth 1 -type d -name 'v*' -printf '%f\\n' | LC_ALL=C sort | tail -n 1").trim()
                            if (!currentFolder) {
                                error('No version directory found for health check tests')
                            }
                            withEnv(["SDK_TEST_FOLDER=${currentFolder}"]) {
                                sh 'bash Scripts/RunHealthCheckTestsInDocker.sh "$SDK_TEST_FOLDER"'
                            }
                        } finally {
                            junit '**/testReport.xml'
                        }
                    }
                } else {
                    stage('tests') {
                        try {
                            sh 'bash Scripts/RunTestsInDocker.sh'
                        } finally {
                            junit '**/testReport.xml'
                        }
                    }
                }
            }
        }
	} finally {
		deleteDir( )
		cleanWs()
	}
}
