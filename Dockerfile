FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./ 
RUN go mod download

COPY . .

RUN go build -o /app/api-gate cmd/main.go

FROM alpine

COPY --from=0 /app/api-gate /api-gate

COPY --from=0 /app/pkg/config/envs/dev.env /envs/dev.env

EXPOSE 3001

CMD [ "/api-gate" ]