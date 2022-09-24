# Hackathon BTG

## endpoints


### /login

```shell
curl --request POST \
  --url http://3.236.209.27:5001/login \
  --header 'Content-Type: application/json' \
  --data '{
	"email":"aaa@gmail.com"
}'
```

### /users

| Método  | Endpoint      |  Descrição                                                                     | cURL |
| -------- | ------------- |  ------------------------------------------------------------------------------- | ---- |
| `POST` | /caixinha       |  Cria uma nova caixinha para um objetivo de emprestimo `usuario comum` ou  projeto `influenciador` |  (1)   |
| `GET` | /caixinhas        |  Retorna a lista de caixinhas criadas. | (2)    |
| `PUT` | /caixinhas/     |  Atualiza uma caixinha atraves do ID passado no corpo da requisicao. | (3)    |
| `DELETE`  | /caixinhas/:id         |  Remove do armazenamento uma caixinha de acordo com seu ID passado na URL. | (4)    |
| `GET`  | /caixinhas/:id |  Retorna os dados especificos de uma caixinha.                                           | (5)    |



#### (1) cURL POST /users/caixinha

```shell
curl --request POST \
  --url http://3.236.209.27:5001/users/caixinha \
  --header 'Content-Type: application/json' \
  --data '{
	"name":"novo studio Canal Nostalgia",
	"description":"O canal nostalgia surgiu blalbla",
	"total_value":100000.00,
	"community_total_value":0.0,
	"community_tax":0.13,
	"market_total_value":100000.00,
	"market_tax":0.35,
	"start_date":"2022-11-04",
	"is_community":true
}'
```

Resposta:

```json
{
	"id": 1,
	"name": "novo studio Canal Nostalgia",
	"description": "O canal nostalgia surgiu blalbla",
	"total_value": 100000,
	"community_total_value": 0,
	"community_tax": 0.13,
	"market_total_value": 100000,
	"market_tax": 0.35,
	"start_date": "2022-11-04",
	"created_at": "2022-09-24T21:49:21.540170766Z",
	"is_community": true
}
```


#### (2) cURL GET /users/caixinhas

```shell
curl --request GET \
  --url http://3.236.209.27:5001/users/caixinhas
```

Resposta: `201 Created`

```json
[
	{
		"id": 2,
		"name": "novo studio Canal Nostalgia",
		"description": "O canal nostalgia surgiu blalbla",
		"total_value": 100000,
		"community_total_value": 0,
		"community_tax": 0.13,
		"market_total_value": 100000,
		"market_tax": 0.35,
		"start_date": "2022-11-04",
		"created_at": "2022-09-24T21:49:21.540170766Z",
		"is_community": true
	},
	{
		"id": 3,
		"name": "novo studio Canal Nostalgia",
		"description": "O canal nostalgia surgiu blalbla",
		"total_value": 100000,
		"community_total_value": 0,
		"community_tax": 0.13,
		"market_total_value": 100000,
		"market_tax": 0.35,
		"start_date": "2022-11-04",
		"created_at": "2022-09-24T21:49:21.540170766Z",
		"is_community": true
	}
]
```

#### (3) cURL PUT /users/caixinhas/

```shell
curl --request PUT \
  --url http://3.236.209.27:5001/users/caixinhas/ \
  --header 'Content-Type: application/json' \
  --data '{
	"id":1,
	"name":"novo studio Canal Nostalgia",
	"total_value":120000.00,
	"community_total_value":0.0,
	"community_tax":0.13,
	"market_total_value":100000.00,
	"market_tax":0.35,
	"start_date":"2022-11-04",
	"is_community":true
}'
```

Resposta: `204 No Content`

#### (4) cURL DELETE /users/caixinhas/:id

```shell
curl --request DELETE \
  --url http://3.236.209.27:5001/users/caixinhas/1
```

Resposta: `204 No Content`

#### (5) cURL GET /users/caixinhas/:id

```shell
curl --request GET \
  --url http://3.236.209.27:5001/users/caixinhas/2
```


Resposta: `200 OK`

```json
{
	"id": 2,
	"name": "novo studio Canal Nostalgia",
	"description": "O canal nostalgia surgiu blalbla",
	"total_value": 100000,
	"community_total_value": 0,
	"community_tax": 0.13,
	"market_total_value": 100000,
	"market_tax": 0.35,
	"start_date": "2022-11-04",
	"created_at": "2022-09-24T21:49:21.540170766Z",
	"is_community": true
}
```
