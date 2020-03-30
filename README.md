# microservice-template
Get started developing microservices on our service mesh with our new protobuf requirements.

## Motivation
Within TCorp, we work hard to develop a maintainable service mesh.

One of our core principles is to offload most middleware to Istio.

Generally speaking, your service should be defined in grpc (and not in json). This repository will help you get started with that.

Transcoding to json [can still be done by Istio](https://github.com/tetratelabs/istio-tools/tree/master/grpc-transcoder) in our middleware layer. The manifests for this should be located at the consuming services, not the service itself. This means that when you are writing a service within TCorp, you are not responsible for maintaining a json API, only the grpc interface.

## Repository
The server is defined in [main.go](main.go), an example client is defined under [client/main.go](client/main.go)

## Starting from scratch
This tutorial requires protoc to be present on your machine ([Installation instructions](https://developers.google.com/protocol-buffers/docs/gotutorial))
1. Create your project in your Go path
2. Setup your go modules: `go mod init`
3. Get grpc: `go get google.golang.org/grpc`
4. Setup your proto interface: eg. `greeter.proto`
5. Generate your client code (Add this to a makefile):  
```bash
  git clone https://github.com/googleapis/googleapis ${GOOGLEAPIS} # in production you want to fetch a specific commit
  proto -I $GOPATH/src  -I ${GOOGLEAPIS} github.com/tcorp-bv/microservice-template/pb/greeter.proto --go_out=plugins=grpc:$GOPATH/src --descriptor_set_out=$GOPATH/src/github.com/tcorp-bv/microservice-template/pb/greeter.desc
   ```
and write the actual service similar to main.go!

## How to use googleapis?
Some examples of how to setup the http bindings for your protocol buffer.

Please see [google's documentation](https://cloud.google.com/apis/design/standard_methods#get) for more.

### Url identifier
```proto
option (google.api.http) = {
    get: "/v1/resource/{id}"
};
```
### Resource in body
```proto
option (google.api.http) = {
    post: "/v1/resource"
    body: "*"
};
```

### Query parameters
If the fields are not defined in the url or body, they are automatically defined as query parameters.
```proto
option (google.api.http) = {
    get: "/v1/resource"
};
```
## How to create a json endpoint?
Json endpoints should be generated by the [service mesh](https://github.com/tetratelabs/istio-tools/tree/master/grpc-transcoder) if necessary.

If for some reason you still want to try out how your service would work with json, take a look at [json_endpoint/main.go](json_endpoint/main.go). This should be used solely for testing, do not deploy it to production or even source control.
