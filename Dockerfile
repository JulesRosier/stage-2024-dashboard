ARG GO_VERSION=1.22.1
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS build
WORKDIR /src

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

ARG TARGETARCH

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 GOARCH=$TARGETARCH go build -ldflags='-s -w -extldflags "-static"' -o /bin/server ./cmd/server

FROM alpine:3.19.1 AS final

ARG GIT_COMMIT=unspecified
LABEL org.opencontainers.image.version=$GIT_COMMIT
LABEL org.opencontainers.image.source=https://github.com/JulesRosier/stage-2024-dashboard

RUN --mount=type=cache,target=/var/cache/apk \
    apk --update add \
        ca-certificates \
        tzdata \
        && \
        update-ca-certificates

# Create a non-privileged user that the app will run under.
# See https://docs.docker.com/go/dockerfile-user-best-practices/
ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser
USER appuser

WORKDIR /app

RUN mkdir ./static

COPY --from=build /bin/server /app/

COPY ./static ./static

EXPOSE 3000

ENV SERVER_BIND=0.0.0.0

ENTRYPOINT [ "/app/server" ]
