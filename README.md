![go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![gql](https://img.shields.io/badge/GraphQl-E10098?style=for-the-badge&logo=graphql&logoColor=white)
![docker](https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white)

# gitconvex GoLang project

This is the back-end go source repo for the [gitconvex](https://github.com/neel1996/gitconvex) project.

<p align="center">
    <img src="https://user-images.githubusercontent.com/47709856/99139411-503eb400-265e-11eb-9b61-05562dd89b8c.png" width="280">
</p>

## Dependencies

- **📜 Git Library** - The project uses [git2go](https://github.com/libgit2/git2go) library for performing all the git
  operations
- **📈 GraphQL** - [gqlgen](https://github.com/99designs/gqlgen) is used for generating boiler-plate GraphQL code which
  is the backbone of the whole project
- **📡 HTTP Router** - [mux](https://github.com/gorilla/mux) is used as the HTTP router for graphql playground and sets
  a handler for the API entry point

### Libgit2 usage

The project uses **git2go** - A libgit2 based binding for go to handle all the git operations. So libgit2 must be setup
properly to run the project.

Follow [these](https://libgit2.org/docs/guides/build-and-link/) instructions to build libgit2
from [source](https://github.com/libgit2/libgit2). Follow this only if you have openssl and libssh setup in your
machine. If this is not the case then follow the detailed instructions mentioned in [LIBGIT_NOTES](LIBGIT_NOTES.md) for
a step-by-step guide

If you stumble upon any challenges, then use [this](https://github.com/neel1996/gitconvex-server/discussions/7)
discussion forum for assistance

### Contribution Guidelines

Fork the repo and raise a new Pull Request to merge your changes with the `main` branch of this repo. Once the review is
complete, the PR will be approved and merged

## Gitconvex as a Git API

Gitconvex can function as a service without the UI. The whole API is build with graphql, and the underlying
queries/mutations can be accessed by invoking `http://localhost:9001/gitconvexapi`.

GraphQL playground is enabled to experiment with the GQL queries and mutations. It can be accessed
from `http://localhost:9001/gitconvexapi/graph`

![gql-playground](https://user-images.githubusercontent.com/47709856/113421107-248f5100-93e8-11eb-8c22-9f6337f7c25f.png)
