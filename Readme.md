# Introduction

Inspired by [unduck](https://unduck.link).

A minimal search redirector implementing DuckDuckGo's bang system.

![se-bang showcase.](./public/banger.gif)

## Usage

```sh
go build
```
And,

`./sebang`

And you can access it using. `http://localhost:8080/`

## To Do:

- [x] Enhance `index.html`
- [x] Use go concurrency to make it faster and recover from error.
  By default, go does that.
- [ ] Implement user-defined bangs (potentially via cookies).
