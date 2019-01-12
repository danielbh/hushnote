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

### MVP
- [ ] Deploy for use with instructions (Docker based deploy)

### Promote
- [ ] Add to danielhollcraft.com portfolio

### TODO
- [ ] remove dev modules from deployed image
- [ ] make read and write safe?
- [ ] Make note unretrievable server side
- [ ] Do this: https://github.com/gin-gonic/gin#build-a-single-binary-with-templates
- [ ] Make purty?
- [ ] Expire secrets after 10 minutes
- [ ] make password protected?
- [ ] email a notification when destructed

