<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/koroutine/rusty">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">Rusty</h3>

  <p align="center">
    Bringing the best bits of Rust to the worst bits of Go
    <br />
    <a href="https://github.com/koroutine/rusty"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/koroutine/rusty">View Examples</a>
    ·
    <a href="https://github.com/koroutine/rusty/issues">Report Bug</a>
    ·
    <a href="https://github.com/koroutine/rusty/issues">Request Feature</a>
  </p>
</div>


### Built With

Generics!



<!-- GETTING STARTED -->
## Getting Started

Install like any other Go library, nice and easy. ❤️ Go

```
go install -u github.com/koroutine/rusty
```


<!-- USAGE EXAMPLES -->
## Usage

Here's the top secret part, so please don't tell anyone...

<sup>Usage is basically just as seen here: https://doc.rust-lang.org/std/</sup>

We've done our best to keep things the way Gophers expect, so it's not 1 to 1 with Rust. But the idea is there. We've been slowly adding support for features as we found them useful in our own work, so the current release in no way covers everything.

```go

import (
  "strconv"

  "github.com/koroutine/rusty"
  "github.com/koroutine/rusty/json"
)

myValue := rusty.ToResult(strconv.ParseInt("20", 10, 64)).Unwrap() // Will panic if error
myRune := rusty.ToEither(utf8.DecodeRuneInString(" ")).Left() // Get the left side of the output
jsonString := json.ToString(&data).Unwrap() // Returns JSON string or panics!
myFloat := rusty.ToString("ABC: 1.0").Replace("ABC: ", "").ParseFloat().Unwrap()
```

<!-- CONTRIBUTING -->
## Contributing

We welcome any Rust or Rust-ish features. The hope is to make our Go life a little easier, while keeping the language we love! Any contributions you make are **greatly appreciated**.

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.
