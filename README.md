# godump

<p align="center">
  <img src="./assets/godump.png" width="600" alt="godump logo">
</p>

<blockquote align="center">
    Pretty-print and debug Go structs with a Laravel-inspired developer experience.
</blockquote>

<p align="center">
    <a href="https://pkg.go.dev/github.com/goforj/godump"><img src="https://pkg.go.dev/badge/github.com/goforj/godump.svg" alt="Go Reference"></a>
    <a href="LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="License: MIT"></a>
</p>

<p align="center">
  <code>godump</code> is a developer-friendly, zero-dependency debug dumper for Go. It provides pretty, colorized terminal output of your structs, slices, maps, and more — complete with cyclic reference detection and control character escaping.
    Inspired by tools like Laravel's <code>dump()</code> and Symfony's VarDumper.
</p>

<p align="center">
  <img src="./assets/demo-terminal.png" width="600">
  Terminal output example
</p>

<p align="center">
  <img src="./assets/demo-html.png" width="600">
  HTML output example
</p>

## ✨ Features

- 🧠 Struct field inspection with visibility markers (`+`, `-`)
- 🔄 Cycle-safe reference tracking
- 🎨 ANSI color or HTML output
- 🧪 Handles slices, maps, nested structs, pointers, time, etc.
- 🪄 Control character escaping (`\n`, `\t`, etc.)

## 📦 Installation

```bash
go get github.com/goforj/godump
````

## 🚀 Usage

```go
import "github.com/goforj/godump"

type Profile struct {
    Age   int
    Email string
}

type User struct {
    Name    string
    Profile Profile
}

user := User{
    Name: "Alice",
    Profile: Profile{
        Age:   30,
        Email: "alice@example.com",
    },
}

// Pretty-print to stdout
godump.Dump(user)

// Dump and exit
godump.Dd(user)

// Get dump as string
output := godump.DumpStr(user)

// HTML for web UI output
html := godump.DumpHTML(user)
```

## 🧪 Example Output

```text
<#dump // main.go:26
#main.User
  +Name    => "Alice"
  +Profile => #main.Profile
    +Age   => 30
    +Email => "alice@example.com"
  }
}
```

---

## 🧩 License

MIT © [goforj](https://github.com/goforj)