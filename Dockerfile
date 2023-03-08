FROM golang:1.19-alpine
WORKDIR /app
COPY . ./
RUN go mod tidy
EXPOSE 8080
RUN go build main.go
CMD [ "./main" ]