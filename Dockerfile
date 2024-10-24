# Usando a imagem oficial do Golang 1.21 para construir o binário
FROM golang:1.21-alpine AS build

# Configura o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copia os arquivos go.mod e go.sum para o contêiner
COPY go.mod go.sum .env ./

# Baixa as dependências
RUN go mod download

# Copia o código do projeto
COPY . .

# Compila a aplicação
RUN go build -o portfolio cmd/main.go

# Fase final para criar a imagem leve de execução
FROM alpine:3.18

WORKDIR /app/portfolio

# Copia o binário compilado para a imagem final
COPY --from=build /app/.env .
COPY --from=build /app/portfolio .

# Porta exposta
EXPOSE 80

# Comando para rodar a aplicação
CMD ["./portfolio"]
