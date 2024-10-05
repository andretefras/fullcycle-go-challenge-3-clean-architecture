# FullCycle - Desafio 3 - Clean Architecture

Implementação do terceiro desafio.

## Setup

O servidor de banco de dados local pode ser executado a partir do comando abaixo `docker-compose up`.

O banco de dados MySql será provisionado realizando a criação do banco de dados automaticamente.

Caso precise popular o banco de dados manualmente, basta importar o arquivo `configs/database.sql`.

Os arquivos .http encontram-se no diretório `api`.

## Interface HTTP

| Opção          | URI            | Porta |
|----------------|----------------|-------|
| Criar pedido   | /orders/create | 8000  |
| Listar pedidos | /orders/list   | 8000  |

## Interface gRPC

| Opção          | Service      | RPC         | Porta   |
|----------------|--------------|-------------|---------|
| Criar pedido   | OrderService | CreateOrder | 50051   |
| Listar pedidos | OrderService | ListOrders  | 50051   |

## Interface GraphQL

| Opção          | Type     | Operation   | Porta |
|----------------|----------|-------------|-------|
| Criar pedido   | mutation | createOrder | 8080  |
| Listar pedidos | query    | orders      | 8080  |