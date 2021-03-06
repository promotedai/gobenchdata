FROM golang:latest

LABEL maintainer="Dan Hill <dhill@promotedai.ai>"
LABEL repository="https://github.com/promotedai/gobenchdata"
LABEL homepage="https://github.com/promotedai/gobenchdata"
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
