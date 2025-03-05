# Teste Técnico para Desenvolvedor Golang Pleno - Importador de Dados

Bem-vindo a minha resposta do teste técnico! O objetivo deste desafio é criar um importador de dados performático.

## Descrição do Desafio

Você deverá desenvolver um importador de dados para uma base de dados PostgreSQL, utilizando a linguagem Go (Golang). Esse importador será responsável por processar e armazenar os dados de um arquivo XLSX (com 80.000 linhas e 55 colunas) fornecido. Durante a avaliação, serão observados critérios como normalização dos dados e performance do importador.

Além disso, será necessário desenvolver uma API em Go, que deverá conter:
- Endpoint de autenticação
- Endpoints de consulta para acessar os dados importados no banco de dados PostgreSQL

## Prazo

A entrega do projeto deve ser realizada em até uma semana.

## Instruções para Rodar a Aplicação

Para rodar a aplicação, siga os passos abaixo:

1. Clone o repositório:
2. Execute o comando `docker compose up -d` para iniciar os serviços.
3. Acesse a aplicação em `http://localhost:80`.

Observação: Foi deixado propositalmente os arquivos de configuração do banco de dados e do ambiente para facilitar a execução da aplicação.


## Documentação

A documentação da API pode ser acessada em `http://localhost:80/api/swagger/index.html`.