FROM golang:1.21 as base

FROM base as build_app
WORKDIR /app
COPY . ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /the-app

FROM build_app as dev_app

FROM build_app as non_root_user
ARG USERNAME=nonroot
ARG USER_UID=1000
ARG USER_GID=$USER_UID
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME
USER $USERNAME

FROM non_root_user as test_app
RUN go test ./...

FROM non_root_user as run_app
CMD ["/the-app"]
