# Establecer la imagen base
FROM golang:1.22 AS builder

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /usr/src/app

# Copiar el código fuente de la aplicación al contenedor
COPY . .

# Instalar las dependencias del proyecto
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o order ./cmd/main.go

FROM scratch
COPY --from=builder /usr/src/app/order ./order
CMD ["./order"]