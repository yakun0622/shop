FROM golang:1.7.0 
MAINTAINER sunlinked

RUN mkdir -p /opt/cb/
ADD . /opt/cb
RUN mkdir -p /opt/cb/logs/
RUN cd /opt/cb && chmod +x cb
CMD ["/opt/cb/cb"]
