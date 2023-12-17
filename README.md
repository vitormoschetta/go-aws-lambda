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


### Help

https://github.com/aws/aws-lambda-go

https://github.com/aws/aws-lambda-go/tree/main/events