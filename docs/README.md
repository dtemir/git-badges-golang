# git-badges-golang

Flex your GitHub stats with badges

Inspired by [git-badges](https://github.com/puf17640/git-badges) and [serverless-github-badges](https://github.com/STRRL/serverless-github-badges) but implemented in golang

## Features

### Organizations

[![Organizations Badge](http://129.80.135.121:8080/organizations?username=dtemir&style=for-the-badge&logo=github&color=yellow)](http://129.80.135.121:8080/organizations?username=dtemir&style=for-the-badge&logo=github&color=yellow)

Number of organizations the user is a part of

#### Endpoint

`http://129.80.135.121:8080/organizations?username={username}&style=for-the-badge&logo=github&color=yellow`

#### Markdown

`[![Organizations Badge](http://129.80.135.121:8080/organizations?username={username}&style=for-the-badge&logo=github&color=yellow)](http://129.80.135.121:8080/organizations?username=dtemir&style=for-the-badge&logo=github&color=yellow)`

---

### Years

[![Years Badge](http://129.80.135.121:8080/years?username=dtemir&style=for-the-badge&logo=github&color=yellow)](http://129.80.135.121:8080/years?username=dtemir&style=for-the-badge&logo=github&color=yellow)

Number of years the user has been registered at GitHub

#### Endpoint

`http://129.80.135.121:8080/years?username={username}&style=for-the-badge&logo=github&color=yellow`

#### Markdown

`[![Years Badge](http://129.80.135.121:8080/years?username={username}&style=for-the-badge&logo=github&color=yellow)](http://129.80.135.121:8080/years?username=dtemir&style=for-the-badge&logo=github&color=yellow)`

---

### Repos

[![Repos Badge](http://129.80.135.121:8080/repos?username=dtemir&style=for-the-badge&logo=github&color=yellow)](http://129.80.135.121:8080/repos?username=dtemir&style=for-the-badge&logo=github&color=yellow)

Number of public repos the user owns

#### Endpoint

`http://129.80.135.121:8080/repos?username={username}&style=for-the-badge&logo=github&color=yellow`

#### Markdown

`[![Repos Badge](http://129.80.135.121:8080/repos?username={username}&style=for-the-badge&logo=github&color=yellow)](http://129.80.135.121:8080/repos?username=dtemir&style=for-the-badge&logo=github&color=yellow)`

---

### Visits

[![Visits Badge](http://129.80.135.121:8080/visits?username=dtemir&repo=dtemir&style=for-the-badge&logo=github&color=yellow)](http://129.80.135.121:8080/visits?username=dtemir&style=for-the-badge&logo=github&color=yellow)

Number of visitors the user had

#### Endpoint

`http://129.80.135.121:8080/visits?username={username}&repo={repo}&style=for-the-badge&logo=github&color=yellow`

#### Markdown

`[![Visits Badge](http://129.80.135.121:8080/visits?username={username}&repo={repo}&style=for-the-badge&logo=github&color=yellow)](http://129.80.135.121:8080/visits?username={username}&repo={repo}&style=for-the-badge&logo=github&color=yellow)`

## Deploy

If you would like to deploy it yourself, please follow these steps:

### Manually

1. Install [Go](https://go.dev/doc/install)
2. Download dependencies with `go mod download`
3. Create a `.env` file with a [GitHub token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token) (look at `.env.example`)
4. Run with `go run main.go`

### Docker

1. Install [Docker Engine](https://docs.docker.com/engine/install/)
2. Create a `.env` file with a [GitHub token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token) (look at `.env.example`)
3. Build an image with `docker build --tag git-badges-golang .`
4. Run the image with `docker run -p 8080:8080/tcp git-badges-golang:latest`
