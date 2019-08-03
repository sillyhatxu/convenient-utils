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
dep ensure -update github.com/sillyhatxu/convenient-utils

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
docker build -t xushikuan/temp-backend .
docker tag xushikuan/temp-backend:latest xushikuan/temp-backend:1.0
docker push xushikuan/temp-backend:1.0
```

# Release Template

### Feature

* [NEW] Support for Go Modules [#17](https://github.com/sillyhatxu/convenient-utils/issues/17)

---

### Bug fix

* [FIX] Truncate Latency precision in long running request [#17](https://github.com/sillyhatxu/convenient-utils/issues/17)

###

```
git tag v1.0.2
git push origin v1.0.2
```
