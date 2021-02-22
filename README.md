# gitconvex GoLang project

This is the back-end go source repo for the [gitconvex](https://github.com/neel1996/gitconvex) project.

<p align="center">
    <img src="https://user-images.githubusercontent.com/47709856/99139411-503eb400-265e-11eb-9b61-05562dd89b8c.png" width="280">
</p>

## Dependencies

The depedency packages used by this project can be
found [here](https://github.com/neel1996/gitconvex-server/network/dependencies)

- **📜 Git Library** - The project uses [git2go](https://github.com/libgit2/git2go) library for performing majority of
  the git operations without relying on the native git client
- **📈 GraphQL** - [gqlgen](https://github.com/99designs/gqlgen) is used for generating boiler plate GraphQL code which
  is the backbone of the whole project
- **📡 HTTP Router** - [mux](https://github.com/gorilla/mux) is used as the HTTP router for graphql playground and sets
  a handler for the API enttry point
- **⌚ Go time Library** - [goment](https://github.com/nleeper/goment) is used for formatting commit times and also for
  calculating the difference between the commit time and the current clock time

### Libgit2 usage

As the project used git2go - A libgit2 based binding for go, libgit2 must be setup properly to run the project. Follow [these](https://libgit2.org/docs/guides/build-and-link/) instructions to build libgit2 from [soruce](https://github.com/libgit2/libgit2)

If you stumble upon any challenger, then use [this](https://github.com/neel1996/gitconvex-server/discussions/7) discussion forum for assistance

### Guidelines

Fork the repo and raise a new Pull Request to merge your branch with the `development` branch of this repo. Once the
review is complete, the PR will be approved and merged with `main`

## Project directory tree

**📂 api** - All the common api modules which does not modify the git repo in anyway resides in this directory

**📂 git** - The files in this directory will handle all the git related operations behind the scenes using `go-git` and
the native `git` client

**📂 graph** - The GQL schema and other files which are used for enabling GQL based communication are included in this
directory.

**📂 tests** - As the name suggests, all the test scripts are stored here

**📂 utils** - The common utility modules which are required by other functions to execute common tasks are stored in
this directory

```
├── api
│   ├── add_repo.go
│   ├── code_file_view.go
│   ├── fetch_repo.go
│   ├── health_check.go
│   ├── repo_status.go
│   └── settings_api.go
├── git
│   ├── git_branch_add.go
│   ├── git_branch_checkout.go
│   ├── git_branch_compare.go
│   ├── git_branch_delete.go
│   ├── git_branch_list.go
│   ├── git_changed_files.go
│   ├── git_clone.go
│   ├── git_commit_changes.go
│   ├── git_commit_compare.go
│   ├── git_commit_files.go
│   ├── git_commit_log_search.go
│   ├── git_commit_logs.go
│   ├── git_fetch.go
│   ├── git_fileline_diff.go
│   ├── git_init.go
│   ├── git_ls_files.go
│   ├── git_pull.go
│   ├── git_push.go
│   ├── git_remote_add.go
│   ├── git_remote_data.go
│   ├── git_remote_name.go
│   ├── git_repo.go
│   ├── git_repo_validate.go
│   ├── git_resetall_items.go
│   ├── git_reset_item.go
│   ├── git_stageall_items.go
│   ├── git_stage_item.go
│   ├── git_total_commits.go
│   └── git_unpushed_commits.go
├── gitclient
├── global
│   └── GlobalLogger.go
├── go.mod
├── go.sum
├── gqlgen.yml
├── graph
│   ├── generated
│   │   └── generated.go
│   ├── model
│   │   ├── aux_models.go
│   │   └── models_gen.go
│   ├── resolver.go
│   ├── schema.graphqls
│   └── schema.resolvers.go
├── make.bat
├── Makefile
├── README.md
├── server.go
├── tests
│   ├── git_branch_add_test.go
│   ├── git_branch_delete_test.go
│   ├── git_branch_test.go
│   ├── git_changed_files_test.go
│   ├── git_clone_test.go
│   ├── git_commit_files_test.go
│   ├── git_commit_logs_test.go
│   ├── git_ls_files_test.go
│   ├── git_remote_add_test.go
│   ├── git_remote_data_test.go
│   ├── git_total_commits_test.go
│   └── health_check_test.go
└── utils
    ├── db_file_reader.go
    ├── env_file_handler.go
    └── git_standalone_client.go

8 directories, 65 files
    
```

