FROM scratch

LABEL maintainer="Joey Espinosa <@particledecay>"

ARG TARGETOS
ARG TARGETARCH

WORKDIR /

COPY dist/kdummy_${TARGETOS}_${TARGETARCH}/* /kdummy

ENTRYPOINT ["/kdummy"]
