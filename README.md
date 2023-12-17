# Go for AWS Lambda

### Steps

01. Generate binary for AWS Lambda (Linux | MacOs):
```
GOOS=linux GOARCH=amd64 go build main.go
```

02. Generate zip file for AWS Lambda:
```
zip function.zip main
```

03. Deploy function to AWS Lambda
E possível fazer deploy da função para AWS Lambda utilizando a CLI do AWS ou utilizando a interface web.

Nota: é necessário alterar o handler para `main` nas configurações da função. O handler default é `hello.`

### Testing

É possível configurar uma lambda function para expor um endpoint.  
Exemplo de request:
```
curl --location 'https://endpoint.lambda-url.us-east-1.on.aws/' \
--header 'Content-Type: application/json' \
--data '{
    "name": "ViThor"
}'
```

Essa nossa função principal processaria o evento, retornando ele mesmo como resposta. Muito útil para entendermos os dados que temos disponíveis no evento.
Resposta:
```
{
    "body": "{\n    \"name\": \"ViThor\"\n}",
    "headers": {
        "accept": "*/*",
        "accept-encoding": "gzip, deflate, br",
        "cache-control": "no-cache",
        "content-length": "24",
        "content-type": "application/json",
        "host": "********.lambda-url.us-east-1.on.aws",
        "postman-token": "00daaa3c-8ff8-4694-becc-26bc8b2f2506",
        "user-agent": "PostmanRuntime/7.36.0",
        "x-amzn-tls-cipher-suite": "ECDHE-RSA-AES128-GCM-SHA256",
        "x-amzn-tls-version": "TLSv1.2",
        "x-amzn-trace-id": "Root=1-657f45ce-4cb3c14e00436ee751059ea2",
        "x-forwarded-for": "54.86.50.139",
        "x-forwarded-port": "443",
        "x-forwarded-proto": "https"
    },
    "isBase64Encoded": false,
    "queryStringParameters": {
        "query1": "teste"
    },
    "rawPath": "/test",
    "rawQueryString": "query1=teste",
    "requestContext": {
        "accountId": "anonymous",
        "apiId": "********",
        "domainName": "********.lambda-url.us-east-1.on.aws",
        "domainPrefix": "********",
        "http": {
            "method": "POST",
            "path": "/test",
            "protocol": "HTTP/1.1",
            "sourceIp": "54.86.50.139",
            "userAgent": "PostmanRuntime/7.36.0"
        },
        "requestId": "4e6ed033-4d0c-4584-bf20-270ef92fa21f",
        "routeKey": "$default",
        "stage": "$default",
        "time": "17/Dec/2023:19:02:38 +0000",
        "timeEpoch": 1702839758296
    },
    "routeKey": "$default",
    "version": "2.0"
}
```

### Help

https://github.com/aws/aws-lambda-go

https://github.com/aws/aws-lambda-go/tree/main/events