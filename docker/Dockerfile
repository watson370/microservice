FROM golang:1.13.8-buster AS builder

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go build -v ./...



FROM debian:buster-slim
# root access for development
RUN echo "root:password" | chpasswd
# RUN set -ex; \     This is a template to test for libs and install them if missing, copied in for future reference
# 	if ! command -v gpg > /dev/null; then \
# 		apt-get update; \
# 		apt-get install -y --no-install-recommends \
# 			gnupg \
# 			dirmngr \
# 		; \
# 		rm -rf /var/lib/apt/lists/*; \
# 	fi

RUN set -eux; \
    groupadd -r microservice --gid=999; \
    useradd -r -g microservice --uid=999 --home-dir=/var/lib/microservice --shell=/bin/bash microservice; \
    mkdir -p /var/lib/microservice; \
    chown -R microservice:microservice /var/lib/microservice

COPY --chown=microservice:microservice --from=builder /go/src/app/microservice /var/lib/microservice
RUN chmod +x /var/lib/microservice/microservice

EXPOSE 8080

USER microservice

ENTRYPOINT [ "/var/lib/microservice/microservice" ]