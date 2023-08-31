# Use la imagen oficial de Go
FROM golang:latest

# Establece el directorio de trabajo
WORKDIR /go/src/app

# Inicializa un módulo de Go, creando un archivo go.mod
RUN go mod init

# Instala gin
RUN go install github.com/codegangsta/gin@latest

# Copia el código local al contenedor
COPY ./app .

# Expone el puerto 8080
EXPOSE 6061 8080 


