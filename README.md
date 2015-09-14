# mycap
The MySql network traffic analyzer

```
sudo apt-get install libpcap-dev

git https://github.com/kshvakov/mycap.git && cd mycap

go build --ldflags '-extldflags "-static" -s' 
```