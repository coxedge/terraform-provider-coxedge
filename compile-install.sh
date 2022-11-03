#!/bin/sh
echo "Compiling and installing terraform-provider-coxedge"

echo "Detecting OS and Architecture"

# Detect OS
if [ "$(uname)" == "${OS}" ]; then
    OS="${OS}"
elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
    OS="linux"
elif [ "$(expr substr $(uname -s) 1 10)" == "MINGW32_NT" ]; then
    OS="windows"
elif [ "$(expr substr $(uname -s) 1 10)" == "MINGW64_NT" ]; then
    OS="windows"
else
    echo "Your platform ($(uname -a)) is not supported."
    exit 1
fi

# Detect Architecture
if [ "$(uname -m)" == "x86_64" ]; then
    ARCH="amd64"
elif [ "$(uname -m)" == "i386" ]; then
    ARCH="386"
elif [ "$(uname -m)" == "i686" ]; then
    ARCH="386"
elif [ "$(uname -m)" == "armv6l" ]; then
    ARCH="arm"
elif [ "$(uname -m)" == "armv7l" ]; then
    ARCH="arm"
elif [ "$(uname -m)" == "aarch64" ]; then
    ARCH="${ARCH}"
else
    echo "Your platform ($(uname -a)) is not supported."
    exit 1
fi

go build -o build/terraform-provider-coxedge
mkdir -p ~/.terraform.d/plugins/coxedge.com/cox/coxedge/0.1/${OS}_${ARCH}/
rm ~/.terraform.d/plugins/coxedge.com/cox/coxedge/0.1/${OS}_${ARCH}/* || true
mv build/terraform-provider-coxedge ~/.terraform.d/plugins/coxedge.com/cox/coxedge/0.1/${OS}_${ARCH}/
