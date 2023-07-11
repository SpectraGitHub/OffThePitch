# syntax=docker/dockerfile:1

FROM golang:1.17

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

# Copy code folders
COPY constants ./constants
COPY dbaxs ./dbaxs
COPY handlers ./handlers
COPY internal ./internal

# copy main server function
COPY *.go ./

#ENV GOPATH /app


# build docker image
RUN go build -o /otp-app

EXPOSE 8080
EXPOSE 80

CMD [ "/otp-app" ]