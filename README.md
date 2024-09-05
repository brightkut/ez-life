# EZ Life

This repository provides a CLI tool designed to streamline the setup process of a development environment. It automates the installation of essential programs and tools, including:

- IntelliJ IDEA: A powerful integrated development environment (IDE) for Java and other programming languages.
- Homebrew: A popular package manager for macOS that simplifies the installation of software.
- Zsh Shell: An advanced shell designed to enhance the terminal experience.
- Hyper: A highly customizable terminal emulator built on web technologies.

By using this tool, users can quickly configure their systems with these key applications, ensuring a consistent and efficient development environment setup.

### Installation
1. Clone this repo
2. Run build command `go build -o ez`
3. Run `sudo mv ez /usr/local/bin/` to move binary file to system path
4. Verify the installation by this command `which ez`