## Hushnote

Share secrets that get deleted once read.

- ['realize' for task runner](https://github.com/oxequa/realize) It must be installed
`go get github.com/oxequa/realize`

#### Commands:

- **make install**: install dependencies
- **make dev**: start development
- **make build**: build go binary

#### Docker

`docker build . -t container-name`
`docker run -d -p 8080:8080 container-name`

#### TODO
- [ ] Add GET / for form to enter secret
- [ ] Add POST / to receive and save secret
- [ ] Deploy for use with instructions (Docker based deploy)
- [ ] Add to danielhollcraft.com portfolio
- [ ] Do this? https://github.com/gin-gonic/gin#build-a-single-binary-with-templates
