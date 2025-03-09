  # Development stage
  FROM golang:1.21-alpine AS development
  
  WORKDIR /app
  
  # Install required build tools and make
  RUN apk add --no-cache gcc musl-dev make git && \
      go install github.com/google/wire/cmd/wire@latest
  
  # Copy go mod files
  COPY go.mod go.sum ./
  
  # Download dependencies
  RUN go mod download
  
  # Copy source code and Makefile
  COPY . .
  
  # Expose port
  EXPOSE 8080
  
  # Development command will use make run
  CMD ["make", "run"] 