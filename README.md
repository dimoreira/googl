# googl 
***

A library (and CLI) for Google Url Shortener service.

## Installation

To install the library:

```
go get github.com/dimoreira/googl
```

To install the library with the CLI:

```
go get github.com/dimoreira/googl/cli/googl
```

## Usage

```
// Shorten an url
googl shorten -k <your-api-key> <url-to-be-shortened>

// Expand an url
googl expand -k <your-api-key> <url-to-be-expanded>
```
