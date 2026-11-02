# CLAUDE.md

## Sobre o projeto

Projeto para gerenciamento de escritórios de advocacia, com controle financeiro (clientes, serviços, faturas e transações) e integração com WhatsApp via API Z-API para envio de mensagens e lembretes. Usa banco de dados PostgreSQL e sobe via Docker (Docker Compose com api, worker, postgres e redis).

## Stack

- Go (module `shizuoka/kakegawa`)
- `gin-gonic/gin` — REST API framework
- `gorm.io/gorm` + `gorm.io/driver/postgres` — ORM (Postgres)
- `atlasgo.io` — intended migration tool (not yet wired in); migrations currently live as plain SQL in `migrations/` (golang-migrate naming convention: `NNNNNN_name.up.sql` / `.down.sql`) and are applied manually (e.g. via the `migrate/migrate` Docker image) until a runner is wired in
- `caarlos0/env` + `joho/godotenv` — config loaded from `.env` / environment variables

## Commands

```bash
go run ./cmd/api          # run the API server (binds 0.0.0.0:$PORT, default 8080, see cmd/api/main.go)
go build ./...            # build everything
go vet ./...              # static checks
go test ./...             # run tests (no test files exist yet)
```

There is no Makefile, linter config, or CI pipeline yet — the commands above are the raw Go toolchain.

## Architecture

```
shizuoka_kakegawa/
├── cmd/
│   ├── api/
│   │   └── main.go              # sobe o servidor HTTP (Gin) e faz a injeção de dependências
│   └── worker/
│       └── main.go              # processa filas: envio de mensagens, lembretes de vencimento
│
├── internal/
│   ├── config/                  # leitura de variáveis de ambiente
│   │
│   ├── server/
│   │   ├── router.go            # registra as rotas de todos os módulos
│   │   └── middleware/          # auth JWT, request-id, logging, recovery, CORS, tenant
│   │
│   ├── auth/                    # usuários, login, escritórios (multi-tenant)
│   │
│   ├── client/
│   │   ├── handler.go           # camada HTTP (só Gin aqui)
│   │   ├── service.go           # regras de negócio
│   │   ├── repository.go        # interface + implementação (Postgres)
│   │   ├── model.go             # entidades do domínio
│   │   └── dto.go               # structs de request/response e validação
│   │
│   ├── whatsapp/
│   │   ├── provider.go          # interface Provider (SendText, SendTemplate...)
│   │   ├── cloudapi/            # implementação Meta Cloud API
│   │   ├── zapi/                # (opcional) outra implementação, ex.: Z-API / Evolution
│   │   ├── webhook.go           # recebe eventos, valida assinatura, enfileira
│   │   └── template.go          # templates de mensagem aprovados
│   │
│   └── platform/                # infraestrutura compartilhada
│       ├── database/            # conexão, helper de transação
│       ├── queue/               # fila (ex.: asynq com Redis)
│       ├── logger/              # slog/zap
│       └── crypto/              # criptografia de campos sensíveis
│
├── migrations/                  # SQL versionado (golang-migrate ou goose)
├── api/
│   └── openapi.yaml             # documentação da API
├── deployments/
│   ├── Dockerfile
│   └── docker-compose.yml       # api, worker, postgres, redis
├── test/                        # testes de integração / e2e
├── .env.example
├── Makefile
└── go.mod

## Domain model (from `schema.sql`)

Multi-tenant billing/service domain: a `tenant` has `clients`, who receive `services` and `invoices`, which generate `transactions`; `users` (tenant staff) have `role`s and their logins are tracked in `access_histories`.
