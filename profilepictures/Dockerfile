# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.6

# Install dependencies
RUN go get gopkg.in/mgo.v2
RUN go get github.com/gorilla/mux

# copy the local package files to the container workspace
ADD . /go/src/github.com/softwarearquitecturegroup1a/1A_profile_picture_ms/profilepictures

# Setting up working directory
WORKDIR /go/src/github.com/softwarearquitecturegroup1a/1A_profile_picture_ms/profilepictures

# Build the profilepictures command inside the container.
RUN go install github.com/softwarearquitecturegroup1a/1A_profile_picture_ms/profilepictures

# Run the profilepictures microservice when the container starts.
ENTRYPOINT /go/bin/profilepictures

# Service listens on port 8080.
EXPOSE 8080
