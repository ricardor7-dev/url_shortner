FROM golang:1.20-alpine AS build

#sets the working directory inside the container to `/app`.All subsequent instructions that use relative paths will be relative to this directory
WORKDIR /app

#Copy go.mod and go.sum files from the local machine to /app inside the container
COPY go.mod go.sum ./

#Download all dependencies
RUN go mod download

#Copy the rest of the source code from the local machine to /app inside the container
COPY . .

#Build the Go app and name the output binary "url_shortner"
RUN go build -o url_shortner .

#multi-stage building, to keep the image final image smaller
FROM alpine:latest

WORKDIR /app

#Copy the Pre-built binary file from the previous stage from /app inside the build container to /app inside the runtime container
COPY --from=build /app/url_shortner .    

#Expose port 8080 to the outside world
EXPOSE 8080

#Command to run the executable
CMD ["./url_shortner"]
