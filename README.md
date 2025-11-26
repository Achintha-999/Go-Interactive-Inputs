# 🚀 Go Interactive Inputs
![License](https://img.shields.io/github/license/Achintha-999/Go-Interactive-Inputs?style=flat-square)
![Stars](https://img.shields.io/github/stars/Achintha-999/Go-Interactive-Inputs?style=flat-square)
![Forks](https://img.shields.io/github/forks/Achintha-999/Go-Interactive-Inputs?style=flat-square)
![Issues](https://img.shields.io/github/issues/Achintha-999/Go-Interactive-Inputs?style=flat-square)
![Last Commit](https://img.shields.io/github/last-commit/Achintha-999/Go-Interactive-Inputs?style=flat-square)
![Repo Size](https://img.shields.io/github/repo-size/Achintha-999/Go-Interactive-Inputs?style=flat-square)
![Contributors](https://img.shields.io/github/contributors/Achintha-999/Go-Interactive-Inputs?style=flat-square)
![Go Version](https://img.shields.io/badge/go-%3E%3D1.18-blue?style=flat-square)
![Lines of Code](https://img.shields.io/tokei/lines/github/Achintha-999/Go-Interactive-Inputs?style=flat-square)

---

✨ Welcome to **Go Interactive Inputs** — compact, copy-paste friendly examples showing how to read interactive input in Go: from simple fmt scans to buffered readers, scanners, flags, and secure password entry.

🎯 Purpose
- 🧰 Provide quick, idiomatic examples for handling user input in CLI apps.
- 📚 Be a learning resource for beginners and a snippet library for devs.
- 🔎 Show common gotchas (newline handling, scanner limits, secure input).

Contents
- [Features](#features) 🧾
- [Getting Started](#getting-started) ⚙️
- [Examples](#examples) 💡
- [Project Structure](#project-structure) 📁
- [Contributing](#contributing) 🤝
- [License](#license) 📜
- [Maintainer](#maintainer) 👤
- [Acknowledgements](#acknowledgements) 🙏

---

Features
--------
- ✅ fmt.Scan / fmt.Scanln examples
- ✅ bufio.Reader full-line reading (with spaces)
- ✅ bufio.Scanner tokenized input example
- ✅ flag-based CLI options
- ✅ Secure (no-echo) password input via golang.org/x/term
- ✅ Tips and warnings (scanner token size, trimming)

Getting Started
---------------
Requirements:
- Go 1.18+ (recommended) 🟦
- Git 🔧

Clone:
```bash
git clone https://github.com/Achintha-999/Go-Interactive-Inputs.git
cd Go-Interactive-Inputs
```

Run an example:
```bash
go run examples/simple_scan.go
```

If secure password example requires term:
```bash
go get golang.org/x/term
```

Examples
--------
1) Simple fmt.Scanln — quick token read
```go
package main

import "fmt"

func main() {
  var name string
  fmt.Print("Enter your name: ")
  fmt.Scanln(&name)
  fmt.Printf("Hello, %s!\n", name)
}
```
Icon: 🙋‍♂️

2) bufio.Reader — read a whole line (including spaces)
```go
package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter a sentence: ")
  line, _ := reader.ReadString('\n')
  line = strings.TrimSpace(line)
  fmt.Println("You typed:", line)
}
```
Icon: 📝

3) bufio.Scanner — tokenized scanning (good for line-by-line input)
```go
package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  fmt.Println("Enter lines (CTRL+D to end):")
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    fmt.Println("Got:", scanner.Text())
  }
}
```
Icon: 🔁  
Note: bufio.Scanner default max token ~64K — use bufio.Reader for huge lines.

4) Command-line flags — standard flag parsing
```go
package main

import (
  "flag"
  "fmt"
)

func main() {
  name := flag.String("name", "World", "name to greet")
  age  := flag.Int("age", 0, "your age")
  flag.Parse()
  fmt.Printf("Hello, %s! (age %d)\n", *name, *age)
}
```
Icon: 🚩

5) Secure password input (no echo)
```go
package main

import (
  "fmt"
  "golang.org/x/term"
  "syscall"
)

func main() {
  fmt.Print("Enter password: ")
  bytePassword, _ := term.ReadPassword(int(syscall.Stdin))
  fmt.Println()
  fmt.Printf("Password length: %d\n", len(bytePassword))
}
```
Icon: 🔒

Project Structure
-----------------
- examples/        - runnable example programs (simple_scan.go, reader.go, scanner.go, flags.go, password.go, ...)
- README.md        - this file
- LICENSE          - add a LICENSE (MIT recommended)
- CONTRIBUTING.md  - (optional) contribution guidelines

Decorations, icons and badges used
---------------------------------
- Top-of-file badges: repo stats, license, last commit, contributors, repo size, Go version, lines-of-code.
- Emojis sprinkled per-section for quick visual scanning (🚀, ✨, 🧰, 🔒).
- If you want extra visual flair I can add:
  - A custom SVG banner (project title + Go gopher)
  - Topic badges (example: golang, cli, input)
  - CI badges if you add a workflow for tests

Contributing
------------
Contributions welcome! Suggested flow:
1. Fork the repo 🍴
2. Add an example or improve docs ✍️
3. Run `go fmt` and `go vet` ✅
4. Open a PR with a clear description 🧾












