# Conduit 

## Minimal cli for transferring messages on sqs 

[![Go Report Card](https://goreportcard.com/badge/github.com/GaruGaru/flaw)](https://goreportcard.com/report/github.com/GaruGaru/flaw)

### Usage 

#### Transfer messages from queues 
```bash
conduit transfer --source=<source_sqs_url> --destination=<destination_sqs_url> --concurrency=10
```

#### Clone messages from queues 
```bash
conduit transfer --source=<source_sqs_url> --destination=<destination_sqs_url> --delete=false 
```

#### Publish messages to queues 
```bash
conduit publish --destination=<destination_sqs_url> "message body" 
```

#### Publish messages to queues using pipes  
```bash
cat file.txt | xargs conduit publish --destination=<destination_sqs_url> 
```


### Run with docker

```bash
docker run \
 -e AWS_REGION=<region> \
 -e AWS_ACCESS_KEY_ID=<access-key> \
 -e AWS_SECRET_ACCESS_KEY=<secret> \
 garugaru/conduit <command>
```

### Install from source


```bash
go get -u github.com/garugaru/conduit
```bash