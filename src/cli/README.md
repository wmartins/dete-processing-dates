# CLI

This entrypoint contains a tiny CLI to interact with the scraper.

## Executing

First, make sure you have all dependencies installed:

```bash
go get ./...
```

Second, build the project, which will produce an executable file called `cli`:

```bash
go build -o cli
```

Then, grab the correct DETE URL and execute it like this:

```bash
DETE_PROCESSING_DATES_URL="<DETE URL>" ./cli
```
