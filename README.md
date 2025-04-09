# htmlLinkParser

Build:
```
git clone https://github.com/taran1sdev/htmlLinkParser.git
cd htmlLinkParser
go install golang.org/x/net/html
go mod init link
go build main.go
```

Usage: 'htmlLinkParser -htmlFile=yourfile.html'

Example: 
```bash
htmlLinkParser -htmlFile=./test/ex1.html
href: /other-page | Text: A link to another page
```

The application takes a html file as input and uses the x/net/html packacke to parse the html tree.

It performs a recursive depth-first-search to find all the ElementNodes that are anchor `<a>` tags.

It creates a struct to store the value of the 'href' Attribute and the Data stored within the FirstChild of each ElementNode.

It prints these values to the command-line.
