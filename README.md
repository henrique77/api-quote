# Quotes API

API de cotação de Frete construida utilizando Golang, Fiber Framework, Docker, MySQL e integração com API externa.


### Executando a API

Usando o 'git clone' faça uma cópia do projeto e em seguida entre no diretório /api-quote
A aplicação e o banco de dados MySQL foram configurados para executar no Docker, então após configurar o .env conforme o arquivo .env-example, execute:
```bash
$ docker compose up --build
```

Com isso a API já estará pronta para uso e rodando localmente na porta definida.
Agora a API pode ser testada usando o [Postman](https://www.postman.com/) ou outra ferramenta de sua preferência
### As rotas implementadas foram:
```
- POST   /v1/quote
- GET    /v1/metrics?last_quotes={?}
```
#### POST
A rota __POST__ recebe os dados de entrada e realiza cotação com a [API da Frete Rápido](https://dev.freterapido.com/ecommerce/cotacao_v3/#).
Para essa rota é esperado um JSON de entrada como do exemplo:
 ```json
 {
	"recipient":{
        "address":{
            "zipcode":"29161376"
        }
	},
	"volumes":[
        {
            "category":7,
            "amount":2,
            "unitary_weight":5,
            "price":349,
            "sku":"abc-teste-123",
            "height":0.5,
            "width":0.3,
            "length":0.3
        },
        {
            "category":7,
            "amount":3,
            "unitary_weight":4,
            "price":556,
            "sku":"abc-teste-527",
            "height":0.7,
            "width":0.6,
            "length":0.15
        }
	]
}
 ```
#### GET
 A rota __GET__ consulta as metricas das cotações armazenadas no banco de dados. Opcionalmente nessa rota pode ser passado o parametro __last_quotes__ que seleciona a quantidade de cotações retornadas (ordem decrescente).

Para mais detalhes e informações de uso, acesse a documentação da API no link:
Obs.: O link considera que a porta da aplicação seja a 3000

<http://localhost:3000/v1/swagger/index.html>

## Sobre mim

Sou Henrique Caires, desenvolvedor de software. Estou a disposição para dúvidas, esclarecimentos e sugestões. Me encontre no linkedin: [Henrique Caires](https://www.linkedin.com/in/henrique-caires)
