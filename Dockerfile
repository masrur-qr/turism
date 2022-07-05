FROM mongo
FROM golang:latest

COPY . /app
WORKDIR /
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
EXPOSE 5522
RUN go build -o /main
CMD [ "/main" ]
# # FROM golang:latest
# FROM mongo

# WORKDIR /app

# COPY go.mod ./
# COPY go.sum ./
# RUN go mod download

# COPY *.go ./

# RUN go build -o /docker-gs-ping

# CMD [ "/docker-gs-ping" ]
# #"""""""""""" new""""""""""""

# # RUN mkdir /app
# # ADD . /app
# # WORKDIR /app
# # # COPY go.mod .
# # # COPY go.sum .
# # # RUN go mod download
# # RUN go build -o main .
# # CMD [ "/app/main" ]
