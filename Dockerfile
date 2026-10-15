#
# Контейнер сборки
#
FROM golang:1.26 as builder


ENV CGO_ENABLED=0

COPY . /go/src/github.com/web-rabis/reader-registration-api
WORKDIR /go/src/github.com/web-rabis/reader-registration-api
RUN \
    version=git describe --abbrev=6 --always --tag; \
    echo "version=$version" && \
    cd cmd/apiserver && \
    go build -a -tags reader_registration_api -installsuffix reader_registration_api -ldflags "-X main.version=${version} -s -w" -o /go/bin/reader-registration-api -mod vendor

#
# Контейнер для получения актуальных SSL/TLS сертификатов
#
FROM alpine:3.16 as alpine
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
RUN addgroup -S reader && adduser -S reader -G reader

# копируем документацию
RUN mkdir -p /usr/share/api
COPY --from=builder /go/src/github.com/web-rabis/reader-registration-api/api /usr/share/api
RUN chown -R reader:reader /usr/share/api

ENTRYPOINT [ "/bin/reader-registration-api" ]

#
# Контейнер рантайма
#
FROM scratch
COPY --from=builder /go/bin/reader-registration-api /bin/reader-registration-api

# копируем сертификаты из alpine
COPY --from=alpine /etc/ssl/certs /etc/ssl/certs

# копируем документацию
COPY --from=alpine /usr/share/api /usr/share/api

# копируем пользователя и группу из alpine
COPY --from=alpine /etc/passwd /etc/passwd
COPY --from=alpine /etc/group /etc/group

USER reader

ENTRYPOINT ["/bin/reader-registration-api"]
