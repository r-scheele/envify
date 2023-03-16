# ENVIFY

## Introduction

This is a command line application designed to convert between two configuration file formats: e.g `.env` to `.json` or `.env` to `.yaml` and vice versa.

## Installation and usage

To use the application, follow these steps:

1. Clone the repository
2. Run `go mod tidy` to install dependencies
3. Run `go build -o envify` to build the application
4. Move the built executable to the `/usr/local/bin` directory by running the command `sudo mv envify /usr/local/bin`
5. Update your bash profile to include the new executable by running the command `source ~/.bash_profile` or `source ~/.zshrc` if you are using zsh as your shell or just restart your terminal
6. Run the command `envify` to see the help menu

## Features

1. `.env` to `.json`
2. `.json` to `.env`

The envify command supports the following flags to perform the desired conversion:

* `-p` or `--path`: Specifies the input file path to convert
* `-o` or `--output`: Specifies the output file to write the converted data to.

For example, to convert an .env file to .json, run the command:

```bash
envify convert -p=app.env -o=app.json 
```

## Contributing

We'll be working between the following configuration files:

* `.env`
* `.json`
* `.yaml`
* `.toml`
* `.xml`
* `.ini`
* `.conf`
* `.cfg`
