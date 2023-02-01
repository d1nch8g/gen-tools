package templates

const BufGenYaml = `version: v1
plugins:
  - plugin: go
    out: gen/pb
    opt: paths=source_relative
  - plugin: go-grpc
    out: gen/pb
    opt:
      - paths=source_relative
      - require_unimplemented_servers=false
`

const BufYaml = `version: v1
breaking:
  use:
    - FILE
lint:
  use:
    - DEFAULT`

const GrpcProto = `syntax = "proto3";
package proto.v1;

option go_package = "gen/go/pb";

// Service description example.
service ExampleService {
  // Method description example.
  rpc Add(AddRequest) returns (AddResponse);
}

message AddRequest {
  string example = 1;
}

message AddResponse {}

`