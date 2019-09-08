# My Service - testing API gateway Lambda with Golang

## Create project

    serverless create -t aws-go-mod -p myservice

## Build

    make

## Deploy

    sls deploy

## Test

    curl -X POST https://swvd26quzk.execute-api.ap-southeast-1.amazonaws.com/dev/hello -d 'Hello, world!' | jq




## Install plugins for Serverless Framework

### Create a package.json file

    npm init -y
    
### Install AWS SAM (Serverless Application Model) plugin

    sls plugin install --name serverless-sam
    
### Install prune (Serverless Application Model) plugin

    sls plugin install -n serverless-prune-plugin
    
### Install stage plugin

    sls plugin install --name serverless-stage-manager
