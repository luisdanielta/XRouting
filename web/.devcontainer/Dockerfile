FROM node:20-bullseye

USER root
ENV DEBIAN_FRONTEND=noninteractive
ENV SHELL=/bin/bash

RUN apt-get update && apt-get install sudo wget nano curl passwd git -y && apt-get clean && rm -rf /var/lib/apt/lists/*

EXPOSE 5173
EXPOSE 4173

RUN echo "node ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers.d/node && \
    chmod 0440 /etc/sudoers.d/node


WORKDIR /workspace
USER node