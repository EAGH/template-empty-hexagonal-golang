FROM golang:1.24rc1-bullseye

RUN apt-get update && apt-get install -y openssh-client git zsh pkg-config libseccomp-dev make gcc

RUN sh -c "$(wget -O- https://github.com/deluan/zsh-in-docker/releases/download/v1.1.5/zsh-in-docker.sh)" -- \
    -t robbyrussell \
    -p git \
    -p https://github.com/zsh-users/zsh-autosuggestions \
    -p https://github.com/zsh-users/zsh-completions
EXPOSE 3001