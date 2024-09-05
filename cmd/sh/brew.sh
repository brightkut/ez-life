#!/bin/bash
echo "Installing Homebrew..."

# Check if user is root
if [ "$EUID" -ne 0 ]; then
  echo "Please run as root"
  exit
fi

/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

echo "Installation complete. You can run Homebrew using the command 'brew --version'"