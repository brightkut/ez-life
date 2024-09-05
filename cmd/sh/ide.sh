#!/bin/bash
echo "Installing IntelliJ..."

# Variables
IDEA_VERSION="2023.2.1"  # Update this to the version you want
INSTALL_DIR="/opt/intellij-idea-ultimate"  # Directory where IntelliJ IDEA will be installed
DOWNLOAD_URL="https://download.jetbrains.com/idea/ideaIU-$IDEA_VERSION.tar.gz"
# Check if user is root
if [ "$EUID" -ne 0 ]; then
  echo "Please run as root"
  exit
fi

echo "Downloading IntelliJ IDEA Ultimate..."
wget -O ideaIU.tar.gz "$DOWNLOAD_URL"

echo "Extracting IntelliJ IDEA..."
mkdir -p "$INSTALL_DIR"
tar -xzf ideaIU.tar.gz -C "$INSTALL_DIR" --strip-components=1

# Clean up
rm ideaIU.tar.gz

echo "Creating symbolic link..."
ln -sf "$INSTALL_DIR/bin/idea.sh" /usr/local/bin/idea

# Optionally, add IntelliJ IDEA to your applications menu (Linux desktop)
echo "[Desktop Entry]
Version=1.0
Type=Application
Name=IntelliJ IDEA Ultimate
Icon=$INSTALL_DIR/bin/idea.png
Exec=\"$INSTALL_DIR/bin/idea.sh\" %f
Comment=The Ultimate Java IDE
Categories=Development;IDE;
Terminal=false
StartupWMClass=jetbrains-idea" > /usr/share/applications/jetbrains-idea.desktop

echo "Installation complete. You can run IntelliJ IDEA using the command 'idea'."
