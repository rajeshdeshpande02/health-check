FROM golang:1.16-alpine
WORKDIR /app
COPY . ./
RUN go mod tidy
EXPOSE 8080
RUN go build main.go
CMD [ "/main" ]