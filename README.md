# aftersales-backend


## Initialize your project

```
dep init
```

## Adding a dependency

```
dep ensure -add github.com/foo/bar github.com/foo/baz...

dep ensure -add github.com/foo/bar@1.0.0 github.com/foo/baz@master
```

## Updating dependencies

```
dep ensure -update github.com/sillyhatxu/go-utils

dep ensure -update
```

## build

```
docker build -f application-api/Dockerfile .
```

## alpine-build Dockerfile

```
FROM alpine

RUN apk add --no-cache tzdata
RUN apk --update --no-cache add curl
RUN apk add --no-cache ca-certificates
```

## alpine-build build

```
docker build -t xushikuan/alpine-build .
docker tag xushikuan/alpine-build:latest xushikuan/alpine-build:1.0
docker push xushikuan/alpine-build:1.0
```

# Release Template

### Feature

* [NEW] Support for Go Modules [#17](https://github.com/sillyhatxu/convenient-utils/issues/17)

---

### Bug fix

* [FIX] Truncate Latency precision in long running request [#17](https://github.com/sillyhatxu/convenient-utils/issues/17)
