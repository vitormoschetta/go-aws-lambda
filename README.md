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
curl --location --request GET 'https://*****.lambda-url.us-east-1.on.aws' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Product A",
    "price": 9.99
}'
```

Essa nossa função principal processaria o evento, retornando ele mesmo como resposta. Muito útil para entendermos os dados que temos disponíveis no evento.
Resposta:
```
{
    "body": "{\n    \"name\": \"Product A\",\n    \"price\": 9.99\n}",
    "headers": {
        "accept": "*/*",
        "accept-encoding": "gzip, deflate, br",
        "cache-control": "no-cache",
        "content-length": "46",
        "content-type": "application/json",
        "host": "**************.lambda-url.us-east-1.on.aws",
        "postman-token": "16841192-061c-463e-b043-207a43740790",
        "user-agent": "PostmanRuntime/7.36.0",
        "x-amzn-tls-cipher-suite": "ECDHE-RSA-AES128-GCM-SHA256",
        "x-amzn-tls-version": "TLSv1.2",
        "x-amzn-trace-id": "Root=1-657f6094-5ffe54db5a6cea0972021be4",
        "x-forwarded-for": "54.86.50.139",
        "x-forwarded-port": "443",
        "x-forwarded-proto": "https"
    },
    "isBase64Encoded": false,
    "rawPath": "/",
    "rawQueryString": "",
    "requestContext": {
        "accountId": "anonymous",
        "apiId": "**************",
        "domainName": "**************.lambda-url.us-east-1.on.aws",
        "domainPrefix": "**************",
        "http": {
            "method": "GET",
            "path": "/",
            "protocol": "HTTP/1.1",
            "sourceIp": "54.86.50.139",
            "userAgent": "PostmanRuntime/7.36.0"
        },
        "requestId": "cb607c52-6b79-4f08-9442-c00379b8d5bc",
        "routeKey": "$default",
        "stage": "$default",
        "time": "17/Dec/2023:20:56:52 +0000",
        "timeEpoch": 1702846612619
    },
    "routeKey": "$default",
    "version": "2.0"
}
```

### Help

https://github.com/aws/aws-lambda-go

https://github.com/aws/aws-lambda-go/tree/main/events