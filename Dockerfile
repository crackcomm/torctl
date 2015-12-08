FROM golang

RUN apt-get update -qy
RUN apt-get upgrade -qy
RUN apt-get install -qy wget git build-essential automake libevent-dev libssl-dev python-setuptools

ENV PROJECT /go/src/bitbucket.org/8us/tower
ENV GOPATH  /go:$PROJECT/Godeps/_workspace
ENV TOR_BIN $PROJECT/bin/linux/amd64/tor
ENV TOWER_GEO_DB $PROJECT/GeoLite2-Country.mmdb

# Download MaxMind GeoLite2 Country database
# RUN wget http://geolite.maxmind.com/download/geoip/database/GeoLite2-Country.mmdb.gz
# RUN gzip -d GeoLite2-Country.mmdb.gz
# RUN mv GeoLite2-Country.mmdb $PROJECT/

ADD . $PROJECT
WORKDIR $PROJECT

RUN ./scripts/install
