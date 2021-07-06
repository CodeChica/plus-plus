# https://github.com/microsoft/vscode-dev-containers/tree/v0.177.0/containers/ruby/.devcontainer/base.Dockerfile
FROM mcr.microsoft.com/vscode/devcontainers/ruby:3
ENV DEBIAN_FRONTEND noninteractive
RUN su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install lts/* 2>&1"
RUN apt-get update && apt-get -y install --no-install-recommends apt-transport-https exuberant-ctags fd-find fzf vim