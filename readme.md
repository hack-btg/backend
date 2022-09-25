# Backend

Este projeto contem todos os microservicos desenvolvidos pelo squad 7 para o hackathon do BTG.

### Projeto `banking-service`

Contem a integracao com a API de bancos do BTG, informacoes do usuario e gerenciamento de crédito a ser tomado pelo usuário. Utiliza a porta 5001.

### projeto `credit-search-service`

Contem a logica de busca do melhor credito entre as diversas instituicoes financeiras (mockado). Retorna os 3 melhores bancos e suas taxas para o crédito a ser tomado. Utiliza a porta 5002.

### projeto `community-service`

Contem a landing page de exemplio do "crowdvesting" coletivo do Felipe Castanhari para o canal Nostalgia. Utiliza a porta 5003.


## Subindo os projetos localmente

Os projetos foram feitos com Docker, assim subir todos eles na sua máquina é tao simples quanto executar os seguintes comandos:

1. Buildar as imagens Docker
```shell
$ make build-all
```

2. Subir os servicos no terminal
```shell
$ make up
```

3. (opcional) subir destacado do terminal
```shell
$ make up-detatched
```

## Versao live

Todos os projetos aqui apresentados podem ser encontrados nos IPs abaixo:

- `banking-service`: http://3.236.209.27:5001/
- `credit-search-service`:http://3.236.209.27:5002/
- `community-service`:http://3.236.209.27:5003/

## Documentacao da API

A principal API do `banking-service` pode ser encontrada [aqui](https://github.com/hack-btg/backend/blob/master/banking-service/readme.md). 

Um arquivo completo com todas as requisicoes, incluindo o ambiente deployado em AWS pode ser encontrado nesta pasta com nome `insomnia.json`. O arquivo pode ser utilizado para teste das requisicoes no programa [Insomnia](https://insomnia.rest/).
