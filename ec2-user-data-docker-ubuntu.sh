#!/bin/bash
set -euxo pipefail

# Redirect output to both console and logfile for easier troubleshooting
exec > >(tee /var/log/user-data-docker-setup.log | logger -t user-data -s 2>/dev/console) 2>&1

export DEBIAN_FRONTEND=noninteractive

# Update apt index and install prerequisites
apt-get update -y
apt-get install -y ca-certificates curl gnupg lsb-release

# Add Docker's official GPG key
install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /etc/apt/keyrings/docker.gpg
chmod a+r /etc/apt/keyrings/docker.gpg

# Add Docker apt repository
ARCH="$(dpkg --print-architecture)"
CODENAME="$(. /etc/os-release && echo "$VERSION_CODENAME")"
echo "deb [arch=${ARCH} signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu ${CODENAME} stable" \
  > /etc/apt/sources.list.d/docker.list

# Install Docker Engine components
apt-get update -y
apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

# Ensure Docker starts now and on reboot
systemctl enable --now docker

# Add common default Ubuntu users to the docker group (if they exist)
for user in ubuntu ec2-user admin; do
  if id -u "$user" >/dev/null 2>&1; then
    usermod -aG docker "$user"
  fi
done

# Basic post-install verification
systemctl is-active docker
/usr/bin/docker --version
/usr/bin/docker compose version
