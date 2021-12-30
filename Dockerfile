FROM ubuntu:21.04

LABEL maintainer="Joey Espinosa <@particledecay>"

ARG TARGETOS
ARG TARGETARCH
ARG TARGETPLATFORM
ARG TARGETVARIANT
ARG BUILDOS
ARG BUILDARCH
ARG BUILDPLATFORM
ARG BUILDVARIANT

WORKDIR /

RUN \
  echo TARGETOS=$TARGETOS && \
  echo TARGETARCH=$TARGETARCH && \
  echo TARGETPLATFORM=$TARGETPLATFORM && \
  echo TARGETVARIANT=$TARGETVARIANT && \
  echo BUILDOS=$BUILDOS && \
  echo BUILDARCH=$BUILDARCH && \
  echo BUILDPLATFORM=$BUILDPLATFORM && \
  echo BUILDVARIANT=$BUILDVARIANT

COPY kdummy_${TARGETOS}_${TARGETARCH}/* /kdummy

ENTRYPOINT ["/kdummy"]
