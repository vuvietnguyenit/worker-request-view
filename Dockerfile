FROM golang:1.21-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
COPY info.user.csv ./
# Copy data file
# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /worker-request-view
# Run
CMD ["/worker-request-view"]