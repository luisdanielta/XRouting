FROM python:3.12-bullseye

USER root
ENV DEBIAN_FRONTEND=noninteractive
ENV SHELL=/bin/bash

RUN apt-get update && apt-get install -y python3-pip python3-venv && apt-get clean && rm -rf /var/lib/apt/lists/*
RUN apt-get update && apt-get install sudo wget nano curl passwd git unzip -y && apt-get clean && rm -rf /var/lib/apt/lists/*

RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
RUN unzip awscliv2.zip
RUN ./aws/install --bin-dir /usr/local/bin --install-dir /usr/local/aws-cli --update

ARG USERNAME=dev
ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN groupadd --gid $USER_GID $USERNAME && \
    useradd -s /bin/bash --uid $USER_UID --gid $USER_GID -m $USERNAME && \
    echo "$USERNAME ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers.d/$USERNAME && \
    chmod 0440 /etc/sudoers.d/$USERNAME

RUN pip3 install --upgrade pip
RUN pip3 install --upgrade setuptools[core]
RUN pip3 install --upgrade build

WORKDIR /workspace

USER $USERNAME