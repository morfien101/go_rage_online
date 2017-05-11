FROM scratch
MAINTAINER Randy Coburn
COPY ./go_rage_online /go_rage_online

CMD ["/go_rage_online"]
