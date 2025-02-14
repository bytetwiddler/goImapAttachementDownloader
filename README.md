
# goImapAttachementDownloader

## Description

Example of how to download attachments from an IMAP server using golang

## Getting Started

### Dependencies

* Go Version 1.23 or greater

### Set environment variables

```

IMAP_USER=<string:   set to your userid, probably your email addr>
IMAP_PASS=<string:   set to your email or app password>
IMAP_PORT=<Integer:  set to your imap servers port (probably 993)>
IMAP_SERVER=<string: set to the dns or ip of your IMAP server>

export IMAP_USER IMAP_PASS IMAP_SERVER IMAP_PORT
```

### Executing program

```
go run main.go

or

go build -o ImapDownloader main.go
./ImapDownloader
```

If all goes well, a ***files*** directory will be created in the cwd. Then
for each email that has attachements a sub directory will be created using the
***from address***.  Individual attachments from that address will be placed 
inside this sub directory. NOTE: Multiple emails from the same 
***from address*** will use the same directory. Emails from the same
***from address*** that have attachments that share the same name as a previous
attachment will overwrite each other, this meets my needs but should be trivial 
to change if you so desire.


```
└──▶ tree files/
files/
├── "nobody" <nobody@gmail.com>
│   ├── ATT00001.txt
│   ├── Easter Bunny.JPG
│   ├── Tired Boys.JPG
│   ├── image001.jpg
│   ├── image001.png
│   ├── image002.jpg
│   ├── image002.png
│   ├── image003.jpg
│   ├── image003.png
│   ├── image004.jpg
│   ├── image004.png
│   ├── image005.jpg
│   ├── image1.jpg
│   ├── image2.jpg
│   ├── image3.jpg
│   └── photo.JPG
...
```


## Help

Gmail will require;
1) imap turned on
2) an application password if 2MFA enabled.


## Authors

Contributors names and contact info

ex. Twiddlebits


## License

This project is licensed under the MIT License - see the LICENSE.md file for
details

## Acknowledgments
```
/*********************************************************
 * using Brian Leishman's go-imap to download            *
 * attachments from an imap server and place them        *
 * in cwd/files/{from address}/                          *
 * Thank you Brian                                       *
 * Brian's pkg: http://github.com/BrianLeishman/go-imap/ *
 *********************************************************/
 ```
