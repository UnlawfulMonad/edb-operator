FROM alpine:3.9

ENV OPERATOR=/usr/local/bin/edb-operator \
    USER_UID=1001 \
    USER_NAME=edb-operator

# install operator binary
COPY build/_output/bin/edb-operator ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
