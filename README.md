# mycap
The MySql network traffic analyzer


## Ubuntu
```
sudo apt-get install libpcap-dev

git clone https://github.com/kshvakov/mycap.git && cd mycap

go build --ldflags '-extldflags "-static" -s' 

sudo ./mycap -bpf_filter "tcp and port 3306"
```

## Mac OS
```
git clone https://github.com/kshvakov/mycap.git && cd mycap

go build --ldflags '-extldflags -s' 

sudo ./mycap -device utun1 -bpf_filter "tcp and port 3306"
```


## WebUI

After start daemon it's possible to open web interface at http://localhost:8080/