#!/bin/bash
echo "Installing Zsh terminal..."

# Check if user is root
if [ "$EUID" -ne 0 ]; then
  echo "Please run as root"
  exit
fi

brew install zsh

echo "Installation complete. You can run Zsh terminal using."