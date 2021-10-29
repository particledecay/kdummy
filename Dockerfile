FROM scratch

LABEL maintainer="Joey Espinosa <@particledecay>"

ARG TARGETARCH

WORKDIR /app
COPY dist/kdummy_linux_$TARGETARCH/kdummy /

ENTRYPOINT ["/app/kdummy"]
