Instructions to get Docker commands in jenkinsfile running-

using Docker as build agent-
sudo groupadd docker
sudo usermod -aG docker jenkins
sudo systemctl restart jenkins

using Podman as build agent (RHEL/CoreOS)-
sudo su root
echo jenkins:10000:65536 >> /etc/subuid
echo jenkins:10000:65536 >> /etc/subgid

Add docker-pipeline plugin

Add docker registry credentials

Create a pipeline and reference this repository and dockerfile
