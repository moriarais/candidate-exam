# SAP CxP Multi Cloud - Candidate Exam

Dear candidate,

Welcome to the SAP CxP Multi Cloud exam. In this exam you will demonstrate your knowledge and understanding of the concepts of Cloud Foundry, the GO programming language, Linux, GitHub, and more.

Each task in the exam has a standard task and an additional **expert** task. If you are applying for the expert role, you should also complete the expert tasks.

The exam assumes basic knowledge of Cloud Foundry and its APIs, and basic knowledge of the GO programming language. You can read more about those topics in the following guides.

* [Cloud Foundry overview](https://docs.cloudfoundry.org/concepts/overview.html)
* [Cloud Foundry CLI](https://docs.cloudfoundry.org/cf-cli/)
* [A tour of the Go programming language](https://tour.golang.org/)

## Creating a working environment

If you are doing the exam on an SAP computer, the working environment has already been set up for you, and you can skip to the [Exam](#exam) section. If you are doing the test at home, you will need to complete the following steps before starting the test.

1. Install the [Go Programming Language](https://golang.org/dl/)

1. Install [Git](https://git-scm.com/downloads)

1. Install the [Cloud Foundry Command Line Interface (CF-CLI)](https://github.com/cloudfoundry/cli#downloads)

1. [Optional] Install [IntelliJ Idea Community Edition](https://www.jetbrains.com/idea/#chooseYourEdition) or another IDE of your choice

1. Create the folder ~/go

1. Set the system environment variable GOPATH to "~/go"

1. Get some required Go dependencies (takes a few minutes)

 ```
 $ go get github.com/cloudfoundry-community/go-cfclient
 $ go get github.com/go-martini/martini
 ```

## Exam

### Prerequisites
 
1. Make sure you got the following from your instructor:

  * URL for the Cloud Foundry API
  * Username and password for Cloud Foundry

1. Fork [this repository](https://github.com/sapmulticloud/candidate-exam) to your own GitHub workspace

1. Clone the forked repository to your computer

1. Copy the **cloudapp** folder from the GIT repository to **~/go/src**.

1. Go to ~/go/src/cloudapp and run main.go

  `$ go run main.go`

1. Use your favorite browser to navigate to [localhost:3000](http://localhost:3000)
  
### Task 1: List Cloud Foundry Buildpacks

* Extend ClientService.go to retrieve the list of available buildpacks from Cloud Foundry

* Add the **/buildpacks** endpoint to return a JSON containing the list of buildpacks. The /buildpacks endpoint should return a JSON similar to the example below.

  ```
  [
    {
      "buildpack": {
        "name": "java_buildpack",
        "enabled": true,
        "locked": false,
        "filename": "java-buildpack-v3.7.1.zip"
      }
    },
    {
      "buildpack": {
        "name": "ruby_buildpack",
        "enabled": true,
        "locked": false,
        "filename": "ruby_buildpack-cached-v1.6.19.zip"
      }
    },
    ...
  ]
  ```
  
  Reference: https://github.com/cloudfoundry-community/go-cfclient
  
* **EXPERT:** Extend and use the buildpacks.go model to marshal and unmarshal the JSON data

### Task 2: Allow shuffling and sorting of the buildpacks list

* Add the URL parameter **/sort**

  The URL **/buildpacks/sort?shuffle** should return a shuffled list  
  The URL **/buildpacks/sort?byName** should return a list sorted by name
  
* **EXPERT:** Add unit tests for the shuffling and sorting methods

### Task 3: Show a welcome message

* The **/** endpoint should return a warm welcome message with the current date and time

* **EXPERT:** Instead of the current date and time, the **/** endpoint should return a message stating the active Cloud Foundry ORG and SPACE

### Task 4: Deploy the application to Cloud Foundry

* Use the CF CLI to deploy the application to your development space

  Reference: [CF CLI](https://docs.cloudfoundry.org/cf-cli/)

  Note: You will need to install and use [Godep](https://github.com/tools/godep) to update your Go dependencies before pushing the application.
  
  Tip: Note that the godep binary will be installed in **~/go/bin**.

* **EXPERT:** Give your application a different name (including its deployed URL prefix)

### Task 5: Submit your test

* Push your code back to your forked repository

* Create a pull request to the original exam repository

  Note: In the pull request, add the URL of your deployed application

## Good luck!
