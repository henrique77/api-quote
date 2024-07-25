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
Como retorno é esperado um array de cotação como mostrado a seguir:
```json
[
    {
        "name": "AZUL CARGO",
        "service": "Convencional",
        "deadline": 2,
        "price": 39.95
    },
    {
        "name": "PRESSA FR (TESTE)",
        "service": "Normal",
        "deadline": 0,
        "price": 63.66
    }
]
```
#### GET
 A rota __GET__ consulta as métricas das cotações armazenadas no banco de dados. Opcionalmente nessa rota pode ser passado o parametro __last_quotes__ que seleciona a quantidade de cotações retornadas (ordem decrescente).

 Como retorno é esperado um objeto com algumas métricas de cotações. Veja o exemplo:
 ```json
{
    "results_per_carrier": {
        "AZUL CARGO": 2,
        "BTU BRASPRESS": 2,
        "JADLOG": 2,
        "PRESSA FR (TESTE)": 2,
        "RAPIDÃO FR (TESTE)": 2
    },
    "total_final_price": {
        "AZUL CARGO": 79.9,
        "BTU BRASPRESS": 191.98,
        "JADLOG": 224.16,
        "PRESSA FR (TESTE)": 127.32,
        "RAPIDÃO FR (TESTE)": 843.24
    },
    "average_final_price": {
        "AZUL CARGO": 39.95,
        "BTU BRASPRESS": 95.99,
        "JADLOG": 112.08,
        "PRESSA FR (TESTE)": 63.66,
        "RAPIDÃO FR (TESTE)": 421.62
    },
    "least_expensive_shipping": 39.95,
    "most_expensive_shipping": 421.62
}
```

Para mais detalhes e informações de uso, acesse a documentação da API no link abaixo:
Obs.: O link considera que a porta da aplicação seja a 3000

<http://localhost:3000/v1/swagger/index.html>

## Sobre mim

Sou Henrique Caires, desenvolvedor de software. Estou a disposição para dúvidas, esclarecimentos e sugestões. Me encontre no linkedin: [Henrique Caires](https://www.linkedin.com/in/henrique-caires)
