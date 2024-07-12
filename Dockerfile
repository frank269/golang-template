FROM golang:1.22-rc AS build-stage

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /runner cmd/main.go


# Run the tests in the container
# FROM build-stage AS run-test-stage
# RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /
COPY --from=build-stage /runner /

USER nonroot:nonroot
ENV TZ=Asia/Ho_Chi_Minh

ARG PORT=8080
EXPOSE ${PORT}
ENTRYPOINT ["/runner"]