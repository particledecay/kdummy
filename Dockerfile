FROM scratch

LABEL maintainer="Joey Espinosa <@particledecay>"

ARG TARGETOS
ARG TARGETARCH

WORKDIR /
COPY kdummy_$TARGETOS_$TARGETARCH/* /kdummy

ENTRYPOINT ["/kdummy"]
