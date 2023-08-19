# Stage 1: Build the Go application
FROM golang:1.19 AS builder

WORKDIR /app

# Copia apenas os arquivos de dependências do Go para aproveitar o cache
COPY go.mod go.sum ./
RUN go mod download

# Copia o código-fonte do aplicativo
COPY . .

# Compila o aplicativo
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/app

# Stage 2: Build the final minimal image
FROM alpine:latest

# Instala dependências necessárias para executar o aplicativo
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copia o executável do estágio anterior
COPY --from=builder /app/app .

# Copia o arquivo .env para o contêiner
COPY .env .

# Define a porta na qual o aplicativo irá escutar
EXPOSE 8000

# Comando para executar o aplicativo
CMD ["./app"]
