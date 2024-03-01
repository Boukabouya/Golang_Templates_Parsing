FROM golang:1.22

WORKDIR /app

# Copy the built Go binary into the container
COPY go.mod .
COPY main.go .
COPY index.html .


# Make the binary executable (if needed)
RUN go build -o bin .

ENTRYPOINT [ "/app/bin" ]

