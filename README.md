# Secure-password

Esse repositório faz parte de um conjunto de desafios que estão disponíveis no <a href="https://github.com/backend-br/desafios/" target="_blank">backend-br</a>.

Esse desafio consiste em implementar um serviço que valide se uma senha é considerada segura com base em critérios pré-definidos.

O serviço irá receber uma requisição com a senha que deve ser validada, da seguinte forma:
```json
{
    "password": "vYQIYxO&p$yfI^r",
}
```

Resultado de uma requisição com senha insegura:
<img src="images/bad-request.png">

Resultado de uma requisição com senha segura:
<img src="images/no-content.png">