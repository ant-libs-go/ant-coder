# ant-coder
code generator

## Build
- go-bindata -pkg templates -o templates/bindata.go templates/...
- go build

## Use
- ant-coder -s [scene]
- support scene: 
  - go\_model: mysql model based on xorm.
  - go\_ui: api scaffold.
  - go\_loop\_worker: loop worker.
  - go\_crontab\_worker: crontab worker.
  - go\_rpcx\_server: rpcx server, see: https://github.com/smallnest/rpcx.
