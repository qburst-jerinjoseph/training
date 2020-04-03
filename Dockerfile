### Stage-1 Build (base)
FROM golang:1.12.0 AS build_base

# Move to project root
WORKDIR /training

COPY go.mod go.sum ./

# Download dependancies
RUN go mod download

### Stage-2 Build (server)
FROM build_base AS build_server

# Copy in files
COPY main.go ./
COPY internal internal

# Install binary (after fetching dependencies)
RUN ["go", "install", "."]


# Install server application
CMD ["training"]