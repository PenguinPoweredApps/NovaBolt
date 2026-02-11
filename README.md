# ⚡ NovaBolt

**NovaBolt** is a high-security, minimalist CLI password generator written in Go. It generates cryptographically secure strings and injects them directly into your clipboard with a built-in "self-destruct" timer.

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org)

## ✨ Features

- **CSPRNG Security**: Uses `crypto/rand` for industrial-grade randomness (not `math/rand`).
- **Zero-Trace Clipboard**: Copies your password automatically and wipes the clipboard after 30 seconds (customizable).
- **Fully Configurable**: Toggle uppercase, numbers, and symbols via simple CLI flags.
- **Single Binary**: No dependencies required after compilation.

## 🚀 Installation

### From Source

1. Ensure you have Go installed.
2. Clone the repo and install dependencies:
   ```bash
   git clone [https://github.com/PenguinPowered/novabolt.git](https://github.com/PenguinPowered/novabolt.git)
   cd novabolt
   go get [github.com/atotto/clipboard](https://github.com/atotto/clipboard)
   go build -o novabolt
   ```

After you build the application with go build -o passgen, you can use it in several ways:

- \*\*Default (16 chars, all types): ./passgen
- \*\*Custom Length (32 chars): ./passgen -l 32
- \*\*No Special Characters: ./passgen -s=false
- \*\*A "Number Only" Pin (length 6): ./passgen -l 6 -u=false -s=false
- \*\*Change the timer to 10 seconds: ./passgen -t 10

Run the tool using flags to customize your output:

### Generate a default 16-character secure password

novabolt

### Generate a 32-character password with no special characters

novabolt -l 32 -s=false

### Generate an 8-digit numeric PIN that clears in 10 seconds

novabolt -l 8 -u=false -s=false -t 10

Available Flags

Flag Description Default

-l Length of the password 16

-u Include uppercase letters true

-n Include numbers true

-s Include special characters true

-t Seconds before clipboard clears 30
