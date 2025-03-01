FROM golang:1.24.0

# Essential tools
RUN apt update && apt install -y curl git nano sudo bash && apt-get clean && rm -rf /var/lib/apt/lists/*

RUN echo "export PATH=$PATH:$GOPATH/bin:/usr/local/go/bin" >> ~/.bashrc

RUN curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s

# Ports for development
EXPOSE 3000
EXPOSE 8000

ARG USERNAME=dev
ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN groupadd --gid $USER_GID $USERNAME && \
    useradd -s /bin/bash --uid $USER_UID --gid $USER_GID -m $USERNAME && \
    echo "$USERNAME ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers.d/$USERNAME && \
    chmod 0440 /etc/sudoers.d/$USERNAME


USER $USERNAME

WORKDIR /workspace