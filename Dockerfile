FROM golang:latest

LABEL maintainer="Robert Lin <robert@promotedai.dev>"
LABEL repository="https://go.promotedai.dev/gobenchdata"
LABEL homepage="https://promotedai.dev/r/gobenchdata"
LABEL version=v1

# set up git
RUN apt-get update && apt-get install -y --no-install-recommends git && rm -rf /var/lib/apt/lists/*

# set up code
WORKDIR /tmp/build
COPY . .

# set up gobenchdata
ENV GO111MODULE=on
RUN go build -ldflags "-X main.Version=$(git describe --tags)" -o /bin/gobenchdata
RUN rm -rf /tmp/build

# init entrypoint
WORKDIR /tmp/entrypoint
ADD entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
