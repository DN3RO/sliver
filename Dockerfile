# STAGE: base
## Compiles Sliver for use
FROM --platform=linux/amd64 golang:1.21.4 as base

### Base packages
RUN apt-get update --fix-missing && apt-get -y install \
    git build-essential zlib1g zlib1g-dev wget zip unzip

### Add sliver user
RUN groupadd -g 999 sliver && useradd -r -u 999 -g sliver sliver
RUN mkdir -p /home/sliver/ && chown -R sliver:sliver /home/sliver

### Build sliver:
WORKDIR /go/src/github.com/bishopfox/sliver
ADD . /go/src/github.com/bishopfox/sliver/
RUN make clean-all 
RUN make 
RUN cp -vv sliver-server /opt/sliver-server 

# STAGE: test
## Run unit tests against the compiled instance
## Use `--target test` in the docker build command to run this stage
FROM --platform=linux/amd64 base as test

RUN apt-get update --fix-missing \
    && apt-get -y upgrade \
    && apt-get -y install \
    curl mingw-w64 binutils-mingw-w64 g++-mingw-w64 

RUN /opt/sliver-server unpack --force 

### Run unit tests
RUN /go/src/github.com/bishopfox/sliver/go-tests.sh

# STAGE: production
## Final dockerized form of Sliver
FROM --platform=linux/amd64 debian:bookworm-slim as production

### Install production packages
RUN apt-get update --fix-missing \
    && apt-get -y upgrade \
    && apt-get -y install \
    libxml2 libxml2-dev libxslt-dev locate gnupg \
    libreadline6-dev libcurl4-openssl-dev git-core \
    libssl-dev libyaml-dev openssl autoconf libtool \
    ncurses-dev bison curl xsel postgresql \
    postgresql-contrib postgresql-client libpq-dev \
    curl libapr1 libaprutil1 libsvn1 \
    libpcap-dev libsqlite3-dev libgmp3-dev \
    mingw-w64 binutils-mingw-w64 g++-mingw-w64 \
    nasm gcc-multilib

### Install MSF for stager generation
RUN curl https://raw.githubusercontent.com/rapid7/metasploit-omnibus/master/config/templates/metasploit-framework-wrappers/msfupdate.erb > msfinstall \
    && chmod 755 msfinstall \
    && ./msfinstall \
    && mkdir -p ~/.msf4/ \
    && touch ~/.msf4/initial_setup_complete 

### Cleanup unneeded packages
RUN apt-get remove -y curl gnupg \
    && apt-get autoremove -y \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

### Add sliver user
RUN groupadd -g 999 sliver \
    && useradd -r -u 999 -g sliver sliver \
    && mkdir -p /home/sliver/ \
    && chown -R sliver:sliver /home/sliver \
    && su -l sliver -c 'mkdir -p ~/.msf4/ && touch ~/.msf4/initial_setup_complete'

### Copy compiled binary
COPY --from=base /opt/sliver-server  /opt/sliver-server

### Unpack Sliver:
USER sliver
RUN /opt/sliver-server unpack --force 

WORKDIR /home/sliver/
VOLUME [ "/home/sliver/.sliver" ]
ENTRYPOINT [ "/opt/sliver-server" ]
