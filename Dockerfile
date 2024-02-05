FROM ubuntu

WORKDIR /app
COPY . /app

RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl && chmod +x ./kubectl && mv ./kubectl /usr/local/bin/kubectl

RUN go build -o myapp .

# Run the Go app
CMD ["./myapp"]
