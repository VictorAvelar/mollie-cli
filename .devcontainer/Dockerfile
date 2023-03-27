ARG GO_VERSION=1-bullseye

ARG VARIANT=1-bullseye
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${GO_VERSION}

# Installs devtooling and font-utils
RUN apt-get update && \
    apt install -y exa bat curl fontconfig

# Install shell superset https://starship.rs
RUN sh -c "$(curl -fsSL https://starship.rs/install.sh)" -y -f

# Install NerdFonts to enabled ligatures and glyphs.
RUN mkdir /root/.fonts && \
    wget https://github.com/ryanoasis/nerd-fonts/releases/download/v2.3.3/FiraCode.zip && \
    unzip FiraCode.zip && \
    rm FiraCode.zip && \
    fc-cache /root/.fonts
