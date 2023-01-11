# git-badges-golang

![git-badges-golang](https://socialify.git.ci/dtemir/git-badges-golang/image?description=1&descriptionEditable=Show%20your%20GitHub%20stats%20with%20Shields.io%20badges%20&font=Raleway&language=1&name=1&theme=Light)

Show your GitHub stats in badges

Inspired by [git-badges](https://github.com/puf17640/git-badges) and [serverless-github-badges](https://github.com/STRRL/serverless-github-badges) but implemented in Golang

## Features

### Visits

[![Visits Badge](https://badges.temir.dev/visits?username=dtemir&repo=git-badges-golang&style=for-the-badge&logo=github&color=yellow)](https://badges.temir.dev/visits?username=dtemir&repo=git-badges-golang&style=for-the-badge&logo=github&color=yellow)

Number of visitors the user had, recorded in a MongoDB database and updated on every GET request

#### Endpoint

`https://badges.temir.dev/visits?username={username}&repo={repo}&style=for-the-badge&logo=github&color=yellow`

#### Markdown

`[![Visits Badge](https://badges.temir.dev/visits?username={username}&repo={repo}&style=for-the-badge&logo=github&color=yellow)](https://badges.temir.dev/visits?username={username}&repo={repo}&style=for-the-badge&logo=github&color=yellow)`

### Organizations

[![Organizations Badge](https://badges.temir.dev/organizations?username=dtemir&style=for-the-badge&logo=github&color=yellow)](https://badges.temir.dev/organizations?username=dtemir&style=for-the-badge&logo=github&color=yellow)

Number of organizations the user is a part of

#### Endpoint

`https://badges.temir.dev/organizations?username={username}&style=for-the-badge&logo=github&color=yellow`

#### Markdown

`[![Organizations Badge](https://badges.temir.dev/organizations?username={username}&style=for-the-badge&logo=github&color=yellow)](https://badges.temir.dev/organizations?username=dtemir&style=for-the-badge&logo=github&color=yellow)`

#### Reference

[GitHub's API](https://docs.github.com/en/rest/orgs/orgs?apiVersion=2022-11-28#list-organizations-for-a-user)

---

### Years

[![Years Badge](https://badges.temir.dev/years?username=dtemir&style=for-the-badge&logo=github&color=yellow)](https://badges.temir.dev/years?username=dtemir&style=for-the-badge&logo=github&color=yellow)

Number of years the user has been registered at GitHub

#### Endpoint

`https://badges.temir.dev/years?username={username}&style=for-the-badge&logo=github&color=yellow`

#### Markdown

`[![Years Badge](https://badges.temir.dev/years?username={username}&style=for-the-badge&logo=github&color=yellow)](https://badges.temir.dev/years?username=dtemir&style=for-the-badge&logo=github&color=yellow)`

#### Reference

[GitHub's API](https://docs.github.com/en/rest/users/users?apiVersion=2022-11-28#get-a-user)

---

### Repos

[![Repos Badge](https://badges.temir.dev/repos?username=dtemir&style=for-the-badge&logo=github&color=yellow)](https://badges.temir.dev/repos?username=dtemir&style=for-the-badge&logo=github&color=yellow)

Number of public repositories the user owns

#### Endpoint

`https://badges.temir.dev/repos?username={username}&style=for-the-badge&logo=github&color=yellow`

#### Markdown

`[![Repos Badge](https://badges.temir.dev/repos?username={username}&style=for-the-badge&logo=github&color=yellow)](https://badges.temir.dev/repos?username=dtemir&style=for-the-badge&logo=github&color=yellow)`

#### Reference

[GitHub's API](https://docs.github.com/en/rest/users/users?apiVersion=2022-11-28#get-a-user)

---

## Deploy

If you would like to deploy it yourself, please follow these steps:

### Manually

1. Install [Go](https://go.dev/doc/install)
2. Install [MongoDB Community](https://www.mongodb.com/docs/manual/tutorial/install-mongodb-on-ubuntu/)
3. Download dependencies with `go mod download`
4. Create a `.env` file with a [GitHub token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token) (look at `.env.example`)
5. Run with `go run *.go`

### Docker (Recommended)

1. Install [Docker Engine](https://docs.docker.com/engine/install/) with the [Compose plugin](https://docs.docker.com/compose/install/linux/)
2. Create a `.env` file with a [GitHub token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token) (look at `.env.example`)
3. Build an image with `docker compose up` (add `-d` if you want to run it in the background)

### Reverse Proxy and SSL/TLS

To avoid having to send all requests to port 8080, you want to setup a reverse proxy that will forward traffic through port 80 (the default port when accessing websites)

If I didn't set it up, you would have to access the website through `badges.temir.dev:8080` instead of regular `badges.temir.dev`

To avoid using HTTP and use secure HTTPS instead, you want to receive an SSL certificate with Let's Encrypt, which is automatable with `certbot`

If this wasn't setup, you would have to make requests to `http://badges.temir.dev` instead of `https://badges.temir.dev` and it would cause problems with common browsers

Find how to accomplish both of these things in this article: [here](https://www.nginx.com/blog/using-free-ssltls-certificates-from-lets-encrypt-with-nginx/)

## CI/CD

To make sure this project is properly maintained, I used GitHub workflows to test and automatically deploy to Oracle Cloud [Micro Instance](https://docs.oracle.com/en-us/iaas/Content/FreeTier/freetier_topic-Always_Free_Resources.htm) that comes with [Always Free Tier](https://www.oracle.com/cloud/free/)

You can find workflows under [Actions](https://github.com/dtemir/git-badges-golang/actions)

1. [check_build.yml](https://github.com/dtemir/git-badges-golang/blob/main/.github/workflows/check_build.yml) to make sure Go compiles
2. [check_compose.yml](https://github.com/dtemir/git-badges-golang/blob/main/.github/workflows/check_compose.yml) to make sure [docker-compose.yml](https://github.com/dtemir/git-badges-golang/blob/main/docker-compose.yml) is up-to-date
3. [deploy.yml](https://github.com/dtemir/git-badges-golang/blob/main/.github/workflows/deploy.yml) to deploy the latest changes to Oracle Cloud
