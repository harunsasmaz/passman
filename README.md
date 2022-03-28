<div align="center">
<h1>passman</h1>

a simple local password manager for daily usage

</div>


---

<div align="center">

![build](https://github.com/harunsasmaz/passman/actions/workflows/build.yaml/badge.svg)
![test](https://github.com/harunsasmaz/passman/actions/workflows/test.yaml/badge.svg)
![lint](https://github.com/harunsasmaz/passman/actions/workflows/gocilint.yaml/badge.svg)
![version](https://img.shields.io/badge/version-1.0.0-blue.svg)
<img alt="GitHub" src="https://img.shields.io/github/license/harunsasmaz/passman?color=blue">
</div>

--- 

## Download

### Homebrew

```
$ brew install harunsasmaz/tap/passman
```

### Go

```
$ go install github.com/ycd/dstp/cmd/dstp@latest
```

### Install from binary

You can see [releases section](https://github.com/harunsasmaz/passman/releases) for binaries and supported platforms. They contain the compiled executable.

### Install from source

0. Verify that you have go 1.17.x installed.

```
$ go version
```

If `go` is not installed, follow directions on [Go Install Page](https://go.dev/doc/install).

1. Clone repository

```
$ git clone https://github.com/harunsasmaz/passman.git
$ cd passman
```

2. Build and Install

On MacOS/BSD

```
# may require you to use sudo
$ make
$ cp passman /usr/local/bin/passman
```

On Unix/Linux

```
# may require you to use sudo
$ go build -o passman cmd/passman/main.go
$ cp passman /usr/local/bin/passman
```

### Verify Installation

```
$ passman -h

NAME:
   passman - generate and manage your passwords on your computer

USAGE:
   passman [COMMANDS] [FLAGS] [ARGS]

VERSION:
   1.0.0

AUTHOR:
   Harun Sasmaz <me@harunsasmaz.com>
...
```

## Usage

### Generate a password

```
USAGE:
   passman generate [FLAGS] [ARGS]
   If you set --level, other options will be discarded

OPTIONS:
   --level value, -l value   choose a strength to use built-in options. Easy: 1, Mid: 2, Hard: 3. Example: -l 1
   --length value, -n value  set the length of the password (default: 16)
   --digit value, -d value   set the number of digits included in the password. (default: 4)
   --symbol value, -s value  set the number of symbols included in the password. (default: 4)
   --upper, -u               set if password can contain uppercase letters. (default: false)
   --no-repeat, -r           set if password should not contain repeated characters. (default: false)
```

**Example**

```
$ passman generate -n 64 -d 8 -s 8 -u -r
```

Output:

```
> Copied to clipboard!
> "VLJpTB%P3O+ZvRm/W9C8czFEXqoKIwhl7uiU`&dgGnNkYM4rb6H?eyax1fjs@2S|"
```

### Save a password and associated account

```
USAGE:
   passman create [FLAGS] [ARGS]
   
OPTIONS:
   --alias value, -a value     set an alias for the new credentials.
   --account value, -u value   set account or host that password will be used for.
   --password value, -p value  set password.
   --generate, -g              generates a new secure password to save. (default: false)
```

**Example**

```
$ passman create -a website -u harunsasmaz.com -g 
```

Output:

```
> Successfully created credentials for alias: website
```

### Get a password and associated account

```
USAGE:
   passman get <alias>
```

**Example**

```
$ passman get website
$ Password: <YOUR_PASSMAN_PASSWORD>
```

Output:

```
> Successfully retrieved password!
> Password is used for account: harunsasmaz.com
> Copied password to clipboard!
```

### Update an existing account or password for an alias

```
USAGE:
   passman update [FLAGS] [ARGS]
   
OPTIONS:
   --alias value, -a value     set alias that you want to update.
   --account value, -u value  set account or host if you want to update.
   --password value, -p value  set password if you want to update.
   --generate, -g              generates a new secure password to update. (default: false)
```

**Example**

```
$ passman update -a website -u harunsasmaz.me
$ Password: <YOUR_PASSMAN_PASSWORD>
```

Output:

```
> Account name renewed!
> Successfully saved changes!
```

### Delete credentials for an alias

```
USAGE:
   passman delete [FLAGS] [ARGS]
   
OPTIONS:
   --alias value, -a value  delete credentials for provided alias
   --all                    delete all stored credentials (default: false)
```

**Example**

```
$ passman delete -a website
$ Password: <YOUR_PASSMAN_PASSWORD>
```

Output:

```
> Successfully deleted credentials for alias: website
```

## Contributing

There are still tons of work to do. So, any kind of help to improve passman is welcomed.

## License

passman's source code is licensed under [GNU GPL-3.0 License](https://choosealicense.com/licenses/gpl-3.0/)
