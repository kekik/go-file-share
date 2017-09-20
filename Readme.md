# Go File Sharing Server

A library that implement file sharing between multiple devices through a central server.


## Building

```shell
go build main.go
./main
```
## Configuration

The library requires the following params to be set either by modifying the config file or by setting an environment variable with the corresponding name:

* BaseURI: Sets the server's endpoint URL that the library will communicate with
* BasePath: Sets the local location of the downloaded files to
* Format: Sets the pattern to be used to organized the uploaded files

The format propperty can include one of the following attributes:
* **{TOKEN}** : the token used to upload the file
* **{EXT}** : the file's extension
* **{DAYNAME}** : The current day's name at the upload time
* **{DAY}** : The current day number at the upload time
* **{MONTH}** : The current month at the upload time
* **{YEAR}** : The current year at the upload time
* **{FILENAME}**: the original file's name

Here is a sample config file written using `yaml`
```yml
BaseURI: "https://240.0.0.1"
BasePath: "/home/storage/"
Format: "uploaded/{YEAR}/{MONTH}/{EXT}/{FILENAME}"
```

## Usage


## Licensing
