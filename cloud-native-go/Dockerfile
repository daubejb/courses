FROM golang:1.7.4-alpine
LABEL Authur="Jeffrey B. Daube"

ENV SOURCES /go/src/github.com/daubejb/courses/cloud-native-go/

COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go install

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT cloud-native-go