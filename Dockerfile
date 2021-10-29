FROM scratch

LABEL maintainer="Joey Espinosa <@particledecay>"

ARG TARGETARCH

WORKDIR /
COPY kdummy_linux_$TARGETARCH /kdummy

ENTRYPOINT ["/kdummy"]
