![Release](https://github.com/yunussandikci/cloudflare-dynamic-dns/workflows/Release/badge.svg)
[![Cloudflare Dynamic DNS](https://img.shields.io/docker/image-size/yunussandikci/cloudflare-dynamic-dns)](https://github.com/yunussandikci/cloudflare-dynamic-dns)
# Cloudflare Dynamic DNS

**Cloudflare Dynamic DNS** is a minimal project that updates your dns records on cloudflare dynmically with its own ip adress within the period you specified. 
Docker image is only ~4 megabyte, and supports `amd64`, `arm/v7` and `arm64` archictectures. Yes, it **supports Raspberry Pi!** 

## Usage
You can with using following example to run it on your Docker
1. Docker Run
```
docker run \
  -e CLOUDFLARE_TOKEN=YOUR_CLOUDFLARE_TOKEN \
  -e INTERVAL=TIME_INTERVAL_IN_MINUTES \ //like 5
  -e DOMAINS=COMMA_SEPERATED_DOMAINS \ //like foo.bar.com, baz,bar.com
   yunussandikci/cloudflare-dynamicdns:1.0.0
```
2. Docker Compose
```
cloudflare-ddns:
  image: yunussandikci/cloudflare-dynamic-dns:1.0.0
  environment:
    CLOUDFLARE_TOKEN: "YOUR_CLOUDFLARE_TOKEN"
    DOMAINS: "COMMA_SEPERATED_DOMAINS" //like foo.bar.com, baz,bar.com
    INTERVAL: "TIME_INTERVAL_IN_MINUTES" //like 5
```
