FROM golang:1.23.3 AS builder
ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /app
COPY . .
RUN go mod tidy
#WORKDIR /app/cmd/main
RUN go build -o main .


FROM mongo:focal
RUN apt-get update && apt-get install -y tzdata \
    && rm -rf /var/lib/apt/lists/*
ENV TZ=Asia/Shanghai
WORKDIR /root/
COPY --from=builder /app/main .
#COPY --from=builder /app/cmd/main/main .
#COPY --from=builder /app/config/config.yaml .
#COPY --from=builder /app/start.sh .
#RUN chmod +x start.sh
EXPOSE 27017 8000
CMD ["sh", "-c", "mongod --bind_ip_all & ./main"]