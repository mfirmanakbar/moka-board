# Moka Board
Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. 

## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## Prerequisites

Install GVM (optional: you can use another way)

```shell
$ zsh < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
```

Install Go v1.14.1 using GVM
```shell
$ gvm install go1.14.1 -B
```

Update Bash/Zsh
```shell
$ open -e ~/.zshrc
```

Update it with:
```shell
export GOROOT_BOOTSTRAP=$GOROOT
export GO111MODULE=auto
```

## Installing
Default GO verision
```shell
$ gvm use go1.14.1 --default
```

Running with VsCode, just put the launch.json. Then tun by Click `Run` on VsCode.
```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Debug Staging",
      "type": "go",
      "request": "launch",
      "mode": "debug",
      "program": "${workspaceRoot}",
      "envFile": "${workspaceFolder}/.env",
      "args": []
    }
  ]
}
```

Running with Terminal
```shell
$ 
  export JURNAL_MOKA_DB_USERNAME=FILL_THIS_VALUE;
  export JURNAL_MOKA_DB_PASSWORD=FILL_THIS_VALUE;
  export JURNAL_MOKA_DB_HOST=FILL_THIS_VALUE;
  export JURNAL_MOKA_DB_PORT=FILL_THIS_VALUE;
  export JURNAL_MOKA_DB_NAME=FILL_THIS_VALUE;
  export BASE_URL_OLD_FE=FILL_THIS_VALUE;
  export BASE_URL_NEW_FE=FILL_THIS_VALUE;

$ go run main.go
```

## Running the Linter 
This way os optinal. On this project, we use `golangci-lint`
```shell
$ golangci-lint run
```

## Running the Test
Lorem Ipsum is simply dummy text of the printing and typesetting industry.

## Deployment
Lorem Ipsum is simply dummy text of the printing and typesetting industry.

## Built With
- [Go]() - programming language `v1.14.1`
- [Beego]() - web framework golang
- [GoMod]() - dependency management golang
- [Gorm]() - object relational mapping golang
- [MySQL]() - Relational database management system

## Contributing
Please read [CONTRIBUTING.md](https://github.com/mfirmanakbar/moka-board) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning
We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags](https://github.com/mfirmanakbar/moka-board/tags) on this repository.

## Authors
- Muhammad Firman Akbar 

See also the list of [contributors](https://github.com/mfirmanakbar/moka-board/contributors) who participated in this project.

## License
This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/mfirmanakbar/moka-board/blob/master/LICENSE) file for details

## Acknowledgments
- Hat tip to anyone whose code was used
- A cup of coffee, of course 😎