FROM scratch

LABEL maintainer="Joey Espinosa <@particledecay>"

ARG TARGETOS
ARG TARGETARCH

WORKDIR /
COPY dist/kdummy_$TARGETOS_$TARGETARCH/* /kdummy

ENTRYPOINT ["/kdummy"]
