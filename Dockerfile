FROM ubuntu:latest

# Atualiza os pacotes do sistema
RUN apt-get update && apt-get upgrade -y

# Instala as dependências necessárias
RUN apt-get install -y wget gnupg ca-certificates

# Adiciona o repositório do Chrome
RUN wget -q -O - https://dl.google.com/linux/linux_signing_key.pub | apt-key add -
RUN echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google.list

# Atualiza os pacotes novamente após adicionar o repositório
RUN apt-get update

# Instala o Chrome
RUN apt-get install -y google-chrome-stable

# Instala o Go 1.20
RUN wget https://golang.org/dl/go1.20.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.20.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"

# Daqui pra baixo é minha aplicação

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
# COPY application_default_credentials.json .

RUN CGO_ENABLED=0 go build -o server cmd/main.go


CMD ["./server"]