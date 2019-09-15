###########################
# Global build args
###########################
# change service and workdir with your 
ARG module=orov.io/micro
ARG workdir=/${module}
ARG key_rsa=./id_rsa
ARG git_host=bitbucket.org

############################
# STEP 1 build executable binary
############################

FROM golang AS build-env

ARG workdir
ARG git_host
ARG key_rsa

WORKDIR ${workdir}


# Install git.
## Git is required for fetching dependencies.
#RUN apt- update && apk add --no-cache git openssh-client

# --------- Start block: Uncomment this if you have private dependencies

## Update ssh hosts
#RUN mkdir -p /root/.ssh &&\
#    chmod 0700 /root/.ssh &&\
#    ssh-keyscan -t rsa ${git_host} > /root/.ssh/known_hosts

## Add the keys
#COPY ${key_rsa} /root/.ssh/id_rsa

# --------- End block


# Building the app
COPY . ${workdir}
RUN cd ${workdir} &&\
    go build -o app

############################
# STEP 2 build tiny executable container
############################

FROM alpine

ARG workdir

WORKDIR ${workdir}
COPY --from=build-env ${workdir}/app ${workdir}/
COPY --from=build-env ${workdir}/migrations/* ${workdir}/migrations/
RUN apk --update add ca-certificates

EXPOSE 8080
ENTRYPOINT ${workdir}/app