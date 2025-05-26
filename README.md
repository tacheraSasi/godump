# godump

<p align="center">
  <img src="./assets/godump.png" width="400" alt="godump logo">
</p>

> Pretty-print and debug Go structs with a Laravel-inspired developer experience.

[![Go Reference](https://pkg.go.dev/badge/github.com/goforj/godump.svg)](https://pkg.go.dev/github.com/goforj/godump)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

---

`godump` is a developer-friendly, zero-dependency debug dumper for Go. It provides pretty, colorized terminal output of your structs, slices, maps, and more — complete with cyclic reference detection and control character escaping.

Inspired by tools like Laravel's `dump()` and Symfony's VarDumper.

---

## ✨ Features

- 🧠 Struct field inspection with visibility markers (`+`, `-`)
- 🔄 Cycle-safe reference tracking
- 🎨 ANSI color or HTML output
- 🧪 Handles slices, maps, nested structs, pointers, time, etc.
- 🪄 Control character escaping (`\n`, `\t`, etc.)

---

## 📦 Installation

```bash
go get github.com/goforj/godump
````

---

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

---

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