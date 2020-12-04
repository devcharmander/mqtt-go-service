# mqtt-go-service

Install Emitter https://github.com/emitter-io/emitter

```
go build && ./mqtt-pub-sub-service 
```

To run the code as a publisher 

```
./mqtt-pub-sub-service -p=true
```

This server also exposes a rest endpoint that can be used to publish messages into emitter broker 
sample post 
```
curl --location --request POST 'http://localhost:8080/mqtt' \
--header 'Content-Type: text/plain' \
--data-raw 'hello world'
```
