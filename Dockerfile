FROM golang:1.19
LABEL authors="andrusyakka@dopcore.com"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -o ./app
RUN chmod +x ./app

EXPOSE 8079

CMD ["/app/app"]
