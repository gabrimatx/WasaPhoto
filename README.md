# WasaPhoto

Web Application for sharing photos and interact with other users.

## How to build container images

### Backend

```sh
docker build -t wasa-photo-backend:latest -f Dockerfile.backend .
```

### Frontend

```sh
docker build -t wasa-photo-frontend:latest -f Dockerfile.frontend .
```

## How to run container images

### Backend

```sh
docker run -it --rm -p 3000:3000 wasa-photo-backend:latest
```

### Frontend

```sh
docker run -it --rm -p 8080:80 wasa-photo-frontend:latest
```

## How to build for development

If you're not using the WebUI, or if you don't want to embed the WebUI into the final executable, then:

```shell
go build ./cmd/webapi/
```

If you're using the WebUI and you want to embed it into the final executable:

```shell
./open-npm.sh
# (here you're inside the NPM container)
npm run build-embed
exit
# (outside the NPM container)
go build -tags webui ./cmd/webapi/
```

## License

See [LICENSE](LICENSE).
