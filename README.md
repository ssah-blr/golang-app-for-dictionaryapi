# Jenkins Pipeline for Building Docker Image with Golang Web API

Welcome to the Jenkins Pipeline Automation repository for building a Docker Image with a Golang-based Web API. 
This repository contains the code and configuration necessary to create a CI/CD process for your Golang application, turning it into a Docker image and deploying it as a web API.

# Background

In this repository, we have a Jenkins pipeline script that automates the process of building a Docker image for a Golang application, which serves as a web API. With this setup, we can integrate and deploy Golang application, ensuring that it is consistently built and deployed whenever changes are pushed to the source code repository.

## Components -

    Jenkins Pipepine to build Docker Image
    GOLang Script which Works as a API and utilises https://dictionaryapi.com/ (https://dictionaryapi.com/products/api-collegiate-thesaurus)
    Docker build to use the GO Script


# STEPS :

To successfully build a Docker image for the GOLang script-based web API, follow the recommended stages below. Followings are stages in [Jenkinsfile](./Jenkinsfile) to used in Jenkins pipeline, to ensure a systematic and controlled deployment process.
1. Checkout Code

    The first stage is to checkout the code from version control system, such as Git, into the Jenkins workspace.

2. Build GOLang Application

    In this stage, we build the GOLang web application using the appropriate build commands. Ensure that all the dependencies are defined and properly set up. [api.go](go/api.go)

3. Unit Testing

    Write unit tests for the GOLang code to ensure the application functions correctly.[api_test.go](go/api_test.go)

4. Create Docker Image

    Use the [Dockerfile](go/Dockerfile) to create a Docker image. This file should specify the necessary dependencies and configurations for the GOLang-based web API.

5. Push Docker Image to Container Registry

    After successfully building the Docker image, push it to a container registry such as Docker Hub or a private repository. This step is essential for sharing and deploying the image.

6. Deploy Container

    Deploy the Docker container to the target environment, which could be a local Docker installation, a Kubernetes cluster, or any other container orchestration platform.

7. Integration Testing (Optional)

    If ther are integration tests, run them to ensure that the application behaves correctly in the target environment.

8. Cleanup and Reporting

    This final stage involves cleaning up any temporary resources and generating reports or notifications about the deployment status and any potential issues.

Getting Started

To get started with this Jenkins Pipeline automation, follow these steps:

    Install Jenkins on a server or use a Jenkins-based CI/CD service.

    Create a Jenkins Pipeline and link it to the Git repository.

    Configure the pipeline stages according to the specific requirements.

    Trigger the pipeline, and Jenkins will automate the process of building the Docker image for your GOLang script-based web API.

By following this automation process, we can save time and ensure consistency in our deployment process, making it easier to manage and scale the GOLang-based web application. Enjoy the benefits of a streamlined and automated CI/CD pipeline with Jenkins!

## Example Usage -
Example Docker Image - https://hub.docker.com/r/ssahblr/go-api
How to use -
```
> docker pull ssahblr/go-api
> docker run -p 5600:5600 ssahblr/go-api:latest <dictionaryapi.com-API-KEY>
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /api/parameter/:param     --> main.main.func1 (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :5600
---
exercise
---
the act of bringing into play or realizing in action : use


OR on Command Line -
> curl http://127.0.0.1:5600/api/parameter/exercise
{"exercise":"the act of bringing into play or realizing in action : use"}

```
