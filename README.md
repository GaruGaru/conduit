# Conduit 

## Minimal cli for transferring messages on sqs 

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


