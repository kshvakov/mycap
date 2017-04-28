# mycap
The MySql network traffic analyzer


## Prepare to install

### Ubuntu
```
sudo apt-get install libpcap-dev
```

## Build project
```
cd $GOPATH/src/

git clone https://github.com/kshvakov/mycap.git
cd $GOPATH/src/mycap/

chmod +x ./build.sh && ./build.sh
```

## Run project

### Run agent

Agent parse trafic on network device and collect mysql queries.
You can change settings in file agent.sh

```
cd $GOPATH/src/mycap/bin/ && chmod +x ./agent.sh && sudo ./agent.sh
```

### Run server

Server collect queries from agents by json-rpc protocol.
Now it's possible to collect queries from one agent.
You can change settings in file server.sh
```
cd $GOPATH/src/mycap/bin/ && chmod +x ./server.sh && ./server.sh
```

### Run web interface
Web app gets queries from server and draw them.
You can change settings in file web.sh

```
cd $GOPATH/src/mycap/bin/ && chmod +x ./web.sh && ./web.sh
```

By default web it's possible to open web interface at http://localhost:9700/
