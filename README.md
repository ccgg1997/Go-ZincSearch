# API Go con ZincSearch

Este proyecto es una API desarrollada en Go que utiliza ZincSearch.

## ¿Qué es ZincSearch?

ZincSearch es un motor de búsqueda de código abierto diseñado para ser rápido y fácil de usar. Proporciona una API RESTful para indexar y buscar documentos JSON. ZincSearch está diseñado para ser utilizado en aplicaciones que requieren búsqueda de texto completo y otras operaciones de búsqueda avanzadas.

## Requisitos

- Docker
- Docker Compose
- Go (opcional, solo necesario si quieres ejecutar la aplicación fuera de Docker)
- Tener disponibles los puertos 8080 (GO), 4080(Zincsearch). (Si necesitas otro puerto, modificalo en el archivo docker-compose.yml) 

## Instalación

1. Clona este repositorio en tu máquina local.  (git clone https://github.com/ccgg1997/Go-ZincSearch.git)
2. Navega a la carpeta del proyecto (/tucarpetaarchivos/ZINCSEARCH).
3. Construye y ejecuta los contenedores de Docker (docker compose up OR  docker-compose up).
4. Go ahora debería estar ejecutándose en `http://localhost:8080` , (Hola mundo en consola).
   ![image](https://github.com/ccgg1997/Go-ZincSearch/assets/89625031/ce88d228-c737-46a4-828c-1fd55f8840d3)

6. ZincSearch deberia estar ejecutando en  `http://localhost:4080`.
![image](https://github.com/ccgg1997/Go-ZincSearch/assets/89625031/7632a238-cd68-4e83-b79d-df15ef52bd93)

## Importante
-En la carpeta app/ se encuetra el archivo main.go






