# Use the official latest Ubuntu release.
FROM ubuntu:latest

RUN apt-get update && apt-get install -y \
    g++ \
    clang \
    clang-tools \
    make \
    gdb \
    cmake \
    vim \
    && rm -rf /var/lib/apt/lists/*