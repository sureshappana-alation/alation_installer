# Alation Installer

## build
Build installer project to generate binary for linux platform using docker
```bash
docker build -o out .
```
This will create a packege containing the binary and the needed resources under installer/out path

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
   
