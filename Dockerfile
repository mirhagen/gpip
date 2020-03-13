# golang:1.14.0-alpine3.11 SHA256 digest.
FROM golang@sha256:6578dc0c1bde86ccef90e23da3cdaa77fe9208d23c1bb31d942c8b663a519fa5 AS builder

ENV USER=gpip
ENV UID=10001

# Create a user.
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nohome" \
    --no-create-home \
    --shell "/sbin/nologin" \
    --uid "${UID}" \
    "${USER}"

# Build final image.
FROM scratch
# Set port and version.
ARG PORT=5050
ARG VERSION=0.1.0
# Set version as environment variable.
ENV GPIP_VERSION=${VERSION}

# Get user settings.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

COPY release/bin/linux/gpip /usr/bin/gpip

EXPOSE ${PORT}

USER gpip:gpip

ENTRYPOINT [ "/usr/bin/gpip" ]
