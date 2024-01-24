# Cryogen

Cryogen is a lightweight encryption/decryption tool written in Go, designed for secure and efficient data handling. Like most other file encryption/decryption tools we use the AES algorithm, to ensure robust protection of sensitive data.

## Features

- **Modern Encryption:** Utilises state-of-the-art cryptographic techniques, including AES, for securing data security.
- **Minimal Design:** A lightweightand straightforward tool, oriented around simplicity and ease of use.
- **Efficient Handling:** Provides efficient encryption/decryption processes for seamless integration into your projects.

## Usage

go run main.go

Flags: -file (specify the file)
       -encrypt (encrypt a specified file)
       -decrypt (decrypt a specified file)
       -key (set the 32-bit decryption key)

## Installation

```bash
go get -u github.com/n0desplit/cryogen
