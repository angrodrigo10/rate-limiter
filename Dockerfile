FROM golang:1.23 AS builder

# Definir diretório de trabalho
WORKDIR /app

# Copiar go.mod e go.sum primeiro para instalar dependências
COPY go.mod go.sum ./

# Baixar as dependências
RUN go mod tidy

# Copiar o código fonte
COPY . .

# Compilar o binário
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/rate-limiter ./cmd/server/main.go

# Executar o binário
CMD ["/app/rate-limiter"]