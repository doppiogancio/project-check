
# Project check

This tool will help checking the project code, executing scripts based on file changed.


## Usage/Examples
It works both with any **diff** or `git status`
```bash
git diff | project-check
git status | project-check
```


## Run Locally

Golang should be already installed. You can install **project-check** as follows: 

```bash
go get -u github.com/doppiogancio/project-check
```

Go to the project directory

```bash
cd my-project
```

Create the configuration file

```bash
touch check.config.yaml
```

Paste the following text

```bash
configurations:
  - name: Catalog Client CS
    directory: catalog-client/
    regex: "catalog-client/.*"
    command:
      name: make
      arg: cs-check

  - name: WILLIAM CS
    directory: william/
    regex: "william/.*"
    command:
      name: make
      arg: cs-check
```

The **check.config.yaml** can contain any number of configurations. Each **configuration** needs:
1. a unique **name**
1. a **directory** path relative to the **project root** where the command will be executed.
1. a **regex** to determine if any file was updated/created/deleted.
1. a **command** to executed in case of updated/created/deleted files.


  