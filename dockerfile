# Use una versión específica de Go
FROM golang:1.21.1

# Actualiza la lista de paquetes y instala bash
RUN apt-get update && apt-get install -y dash && rm -rf /var/lib/apt/lists/*

# Establece el directorio de trabajo
WORKDIR /go/src/app

# Inicializa un módulo de Go, creando un archivo go.mod
RUN go mod init

# Instala swag y gin
RUN go install github.com/swaggo/swag/cmd/swag@latest && \
    go install github.com/codegangsta/gin@latest

# Copia el código local al contenedor
COPY ./app .

# Genera la documentación de Swagger
WORKDIR /go/src/app/cmd/main
RUN swag init
WORKDIR /go/src/app

# Expone el puerto 8080 y 6061
EXPOSE 6061 8080 
