FROM python:3.10-bullseye

USER root
ENV DEBIAN_FRONTEND=noninteractive
ENV SHELL=/bin/bash

RUN apt-get update && apt-get install -y python3-pip python3-venv && apt-get clean && rm -rf /var/lib/apt/lists/*
RUN apt-get update && apt-get install sudo wget nano curl passwd git -y && apt-get clean && rm -rf /var/lib/apt/lists/*

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