# Use the official Node.js image for building the frontend
FROM --platform=linux/amd64  node:18-alpine AS frontend-builder

# Set the working directory
WORKDIR /app-build

# Copy only the necessary files for installing dependencies
COPY app .

# Install dependencies
RUN npm i 
RUN npm run build


# Use a smaller base image for the server
FROM --platform=linux/amd64  golang:1.21-alpine3.19 AS server-builder

# Set the working directory
WORKDIR /build

# Copy the source files
COPY server/ ./

# downloading modules 
RUN go mod download

# Build the server app
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o server ./cmd/main.go



# Run from scratch

FROM --platform=linux/amd64 scratch 

WORKDIR /app 

COPY --from=frontend-builder /app-build/dist /app/dist 
COPY --from=server-builder /build/server /app/server 
COPY --from=server-builder /tmp /tmp 
COPY --from=server-builder /build/template/template.tmpl /app/template/template.tmpl
COPY --from=server-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt


# Expose the port your server is running on
EXPOSE 3000


# Command to run the server
CMD ["/app/server"]
