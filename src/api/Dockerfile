FROM go-base:latest AS builder

WORKDIR /src/api
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api .


FROM alpine:latest
WORKDIR /app
COPY --from=builder /src/api/api .
EXPOSE 8888
CMD ["/app/api"]
