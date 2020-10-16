# Jenkins/docker config

Install Jenkins using a linux package manager following docs from their website. Make sure You have docker or podman installed, and git installed

## Configure Docker/Podman

Whether using docker or podman- you'll have to add permissions so that the jenkins user can build containers using the tool. If on a Red Hat associated Distro such as RHEL or CentOS- you'll want to use podman

### using Docker as build agent-

```bash
sudo groupadd docker
sudo usermod -aG docker jenkins
sudo systemctl restart jenkins
```

Add insecure-registry to /etc/docker/daemon.json (make it if it doesn't exist)

```json
{
  "insecure-registries" : ["myregistrydomain.com:5000"]
}
```

### using Podman as build agent (RHEL/CentOS/Fedora)-

```bash
sudo su root
echo jenkins:10000:65536 >> /etc/subuid
echo jenkins:10000:65536 >> /etc/subgid
```

Add insecure registry to /etc/containers/registries.conf

```
[registries.insecure]
registries = ["default-route-openshift-image-registry.apps.cpat-tf.ocp.csplab.local"]
```

## Add plugin

Install docker-pipeline and git plugins on jenkins

Add container registry credentials

Create a pipeline and reference this repository and jenkinsfile
