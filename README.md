# kunigu

Just need to write
```
{{kunigu @file.html}}
```
inside the file, and run:
```
./kunigu
```

All the files including the {{kunigu file.html}} will be processed, and will save a copy named fileOUT.html


Also, can include Markdown files as HTML content:
```
{{kunigu @file.html @--md-to-html}}
```
