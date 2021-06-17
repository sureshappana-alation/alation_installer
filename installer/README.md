# Alation Installer

## build

Build installer project to generate binary for linux platform using docker
```bash
docker build -o out .
```


## install
1. Move the installer bundle to your Linux box and decompress it
```bash
tar xvzf installer.tar.gz 
```


2. Download the kURL Kubernetes bootstrapper package using the temporarily script in res folder - build pipeline will replace this
```bash
sudo res && bash kurl_installer.sh
cd ..
```


3. Run the installer binary 
```bash
sudo ./installer
``` 
   
