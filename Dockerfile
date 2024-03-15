# Establecer la imagen base
FROM golang:latest

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar el código fuente de la aplicación al contenedor
COPY order .

# Instalar las dependencias del proyecto
RUN go mod download

# Compilar la aplicación
RUN go build -o main ./cmd/main.go

# Exponer el puerto en el que se ejecuta la aplicación
EXPOSE 3000

# Comando para ejecutar la aplicación
CMD ["./main"]
