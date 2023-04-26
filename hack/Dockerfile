FROM --platform=$TARGETPLATFORM docker.io/library/ubuntu

ARG TARGETPLATFORM
ARG TARGETARCH

ADD ./linux/${TARGETARCH}/enet-exporter /usr/local/bin/enet-exporter
ADD ./linux/${TARGETARCH}/enet /usr/local/bin/enet