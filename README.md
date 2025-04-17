# API CRUD em Go com Docker e PostgreSQL

## 🛢️ Modelagem SQL

<img src="https://github.com/user-attachments/assets/fa899bb2-1b5b-4768-b203-6c0d493f669d" width="700" />

## ⚙️ Requisitos

- [Docker](https://www.docker.com/) instalado
- [Git](https://git-scm.com/downloads) instalado

## 🚀 Como executar

1. Abra o terminal

2. Clone o repositório:

```bash
git clone https://github.com/luizpaulo73/car-store-crud.git
cd car-store-crud
```

3. Suba os containers com Docker Compose:

```bash
docker build -t crud-go .
docker compose up -d
docker compose start
docker start go-app
```

A API estará disponível em: http://localhost:8000
O banco de dados PostgreSQL estará acessível na porta padrão 5432.
