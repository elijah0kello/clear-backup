# CLEAR OLD  BACK UP TOOL
This tool reads a directory's contents and deletes any file that is older than 1 week 

## Usage

To build use this tool, you would need to build an executable from it, and then you can execute it on your 
preferred platform.

Run the following commands

To build

`cd` into this directory

Then run 
```bash 
go build -o ~/clear-old-backup ./main.go
```

This will produce a file `clear-old-backup` in your home directory.

You can execute it by calling like this

```bash 
cd ~
```
```bash
./clear-old-backup
```
