# go captcha

<p align="center">
  <img src="https://github.com/user-attachments/assets/5e75b2bf-d7b7-47d6-b05a-ca059c175dc9" alt="gocaptcha logo" width="170px"/>
</p>
<p align="center">
  <a href="https://pkg.go.dev/github.com/kal72/go-captcha">
    <img src="https://pkg.go.dev/badge/github.com/kal72/go-captcha.svg" alt="Go Reference">
  </a>
  <a href="https://goreportcard.com/report/github.com/kal72/go-captcha">
    <img src="https://goreportcard.com/badge/github.com/kal72/go-captcha" alt="Go Report Card">
  </a>
  <a href="LICENSE">
    <img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License: MIT">
  </a>
  <img src="https://img.shields.io/github/v/release/kal72/go-captcha?sort=semver" alt="GitHub release (latest by date)">
</p>



> 🧩 **go captcha** is a simple, lightweight library for generating and verifying image-based CAPTCHA codes written in Go.

---

## ✨ Features

- 🖼️ Generate CAPTCHA images with customizable size and font  
- 🔐 Built-in encryption for secure token handling  
- 🧠 Optional cache drivers: in-memory or Redis  
- ⚙️ Flexible configuration using options pattern  
- 💨 Zero external dependencies (except `freetype` and `go-redis/v9`)  

---

## 🚀 Installation

```bash
go get github.com/kal72/go-captcha@latest
```

---

## 🧱 Usage Example

### Basic Example
```go
package main

import (
    "fmt"
    "github.com/kal72/go-captcha"
)

func main() {
    cap := captcha.New("secret-key")

    base64Image, text, token, err := cap.Generate()
	if err != nil {
		panic(err)
	}

	fmt.Println("text: " + text)
	fmt.Println("token: " + token)
	fmt.Println("image: " + base64Image)

    // Verify
    _, err = cap.Verify(text, token)
	if err != nil {
		panic("verify failed: " + err.Error())
	} else {
        fmt.Println("verify success ")
    }
}
```

### Using Redis Driver
```go
import (
    "github.com/kal72/go-captcha"
    "github.com/kal72/go-captcha/driver/redisstore"
)

cap := captcha.New("secret-key",
    captcha.WithStore(redisstore.New(redisstore.Config{
        Addr: "localhost:6379",
        Password: "pass",
        DB: 0,
        Prefix: "captcha",
    })),
)
```

---

## 📸 Example CAPTCHA Image

Example output image generated:

[CAPTCHA Example](https://github.com/kal72/go-captcha/issues/1)


---

## ⚙️ Configuration Options

| Option | Description |
|---------|--------------|
| `WithLength(n int)` | Number of characters in the CAPTCHA code |
| `WithSize(width, height int)` | CAPTCHA image size in pixels |
| `WithFontSize(size int)` | Font size in pixels |
| `WithExpire(seconds int)` | Expiration time in seconds |
| `WithStore(store Store)` | Use custom store driver (memory, Redis, or your own) |

---

## 🧩 Custom Driver

To create a custom driver, implement the `Store` interface:
```go
type Store interface {
    Set(key, value string, ttl time.Duration) error
    Get(key string) (string, error)
    Delete(key string) error
}
```

Then register it via:
```go
captcha.WithStore(myCustomDriver)
```

---

## 🖋️ Example Directory

See full runnable examples in the [`examples/`](examples) folder:
```
examples/
 ├── basic/
 ├── redis/
```

Run an example:
```bash
go run ./examples/basic/main.go
```

---

## 📦 Project Structure

```
gocaptcha/
├── captcha.go
├── config.go
├── store.go
├── driver/
│   ├── memorystore/
│   └── redisstore/
├── internal/
│   ├── image/
│   ├── random/
│   └── assets/
│       └── fonts/
│           └── DejaVuSans-Bold.ttf
└── examples/
```
