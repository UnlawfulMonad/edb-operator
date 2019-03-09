FROM alpine:3.9

ENV OPERATOR=/usr/local/bin/edb-operator \
    USER_UID=1001 \
    USER_NAME=edb-operator

COPY build/bin /usr/local/bin
RUN apk update --no-cache && \
    apk upgrade --no-cache && \
    /usr/local/bin/user_setup

# install operator binary
COPY build/_output/bin/edb-operator ${OPERATOR}

USER ${USER_UID}
ENTRYPOINT ["/usr/local/bin/entrypoint"]
