
# Build Stage
FROM golang:1.22.2 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /bande-a-part-api

## Test Stage
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy Stage
FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /
COPY --from=build-stage /bande-a-part-api /bande-a-part-api

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/bande-a-part-api" ]
