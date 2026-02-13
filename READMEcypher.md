# Cypher Tool

## Description
The Cypher Tool is a command-line application written in Go that allows users to encrypt and decrypt messages using simple encryption techniques.  
The program interacts with the user through the terminal, validates user input, and ensures that non-alphabet characters (such as spaces, numbers, and symbols) remain unchanged during encryption and decryption.

This tool was created to practice problem-solving, input validation, and implementing basic encryption algorithms.

---

## Features
- Encrypt or decrypt messages
- Interactive command-line interface
- Input validation with trimmed user input
- Preserves non-alphabet characters
- Supports multiple encryption methods

---

## Usage

Run the program from the terminal:

```bash
$ ./cyphertool

Welcome to the Cypher Tool!

Select operation (1/2):
1. Encrypt.
2. Decrypt.
2

Select cypher (1/2):
1. ROT13.
2. Reverse.
2

Enter the message:
zb

Decrypted message using reverse:
ay


