FROM scratch

LABEL maintainer="Joey Espinosa <@particledecay>"

COPY kdummy /

WORKDIR /

ENTRYPOINT ["/kdummy"]
