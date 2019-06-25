FROM registry.access.redhat.com/ubi7-dev-preview/ubi-minimal:7.6

ENV OPERATOR=/usr/local/bin/kubefed-operator \
    USER_UID=1001 \
    USER_NAME=kubefed-operator

LABEL com.redhat.delivery.appregistry=true

ADD deploy/olm-catalog/kubefed-operator /manifests

# install operator binary
COPY build/_output/bin/kubefed-operator ${OPERATOR}

COPY deploy /deploy

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
