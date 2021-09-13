# Alation Installer

## build
Build installer project to generate binary for linux platform using docker
```bash
docker build -f Dockerfile-local -o out .
```
This will create a package containing the binary and the needed resources under installer/out path

## install
1. Upload the installer bundle to your Linux box and decompress it
```bash
tar xvzf installer.tar.gz 
cd installer
```


2. Download the kURL Kubernetes bootstrapper package using the temporarily script in res folder - build pipeline will replace this
```bash
bash kurl_downloader.sh
```


3. Run the installer binary 
```bash
./installer
``` 
   
## Local registry
The local registry is available as a service inside kubernetes cluster and should be referred to as "registry.kurl"
This url is on http protocol and no ssl is needed as it is only accessible inside the cluster.
### example
registry.kurl/<IMAGE_NAME>:<IMAGE_VERSION>

The local registry storage defined to have 15Gi of space which is going to be hosting Alation images. 
