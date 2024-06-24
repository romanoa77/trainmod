FROM alpine:latest

### Env var and cmd line arguments####
ARG srvd=/var/www
ARG repo=https://github.com/romanoa77/repdepo
ARG reponm=repdepo

ENV PAYLOAD=test.py
ENV PAYLOADNM=test

ENV GREET="Hello from Gunicorn server!"
######################################

## webserverver filesystem
RUN mkdir ${srvd}
######################################

RUN apk add  git &&apk cache clean

RUN apk add  py3-flask   &&apk cache clean
RUN apk add  py3-gunicorn   &&apk cache clean


RUN git clone ${repo}



ADD gunicorn.conf.py /etc
ADD  init.py ${srvd}

ADD ${reponm}/${PAYLOAD} ${srvd}

ADD ${reponm}/templates ${srvd}

WORKDIR ${srvd}



CMD [ "gunicorn", "-c","/etc/gunicorn.conf.py", "init:create_app()" ]

