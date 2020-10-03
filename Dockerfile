# This container only includes the default configuration file. In order to use
# cloud(K8s, AWS, GCE and others) integration, you need to use your own configuration file.
# So you should run this container like this:
#
# docker run -e OLRICD_CONFIG=/etc/olricd/config_ext/olricd.yaml  -v PATH_TO_LOCAL_CONFIG_FOLDER:/etc/olricd/config_ext olricio/olric-cloud-plugin:latest
#
# Please take a look at olricd-sample.yaml. Without this, olricd runs at standalone mode.
# Further details for configuration: https://github.com/buraksezer/olric-cloud-plugin

FROM golang:latest as build

WORKDIR /src/
COPY . /src/
RUN go mod download

RUN CGO_ENABLED=1 go build -ldflags="-s -w" -buildmode=plugin -o /usr/lib/olric-cloud-plugin.so

FROM olricio/olricd:v0.3.0-beta.4
COPY --from=build /usr/lib/olric-cloud-plugin.so /usr/lib/olric-cloud-plugin.so