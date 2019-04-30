# ant-coder
code generator

## Build
- go-bindata -pkg templates -o templates/bindata.go templates/...
- go build

## Use
- ant-coder -s <scene>
- support scene: 
 - go\_model: A mysql model based on xorm.
 - go\_ui: An api scaffold.
