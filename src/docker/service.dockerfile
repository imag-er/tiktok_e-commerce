FROM go-base:latest AS builder

WORKDIR /src/rpc/service/
COPY . .
RUN sh ./build.sh


FROM alpine:latest

COPY --from=builder /src/rpc/service/output/ /app/output
CMD ["sh","/app/output/bootstrap.sh"]
