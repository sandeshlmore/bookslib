FROM golang:1.17-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# RUN apt update && apt upgrade && \
#     apt add --no-cache bash git openssh

RUN mkdir /app

WORKDIR /app

# # Copy go mod and sum files
COPY go.mod go.sum ./

# # Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . /app/


# Expose port 8080 to the outside world
EXPOSE 8080


RUN chmod +x ./docker-entrypoint.sh


# # Build the Go app
RUN go build ./cmd/main.go

# # Run the executable
ENTRYPOINT ["/bin/bash","docker-entrypoint.sh"]
# CMD ["./main"]
# CMD ["sleep", "6600"]