# ant-coder
一个常用脚手架的生成器

## Build
- go-bindata -pkg templates -o templates/bindata.go templates/...
- go build

## Use
- ant-coder -s [scene]
- support scene: 
  - go\_model: 基于xorm可以根据表结构自动生成mysql model，model的结构借鉴PHP的Yii框架。核心意图是约束xorm过于灵活的使用方式，让任何实现只有一种写法
  - go\_ui: API接口脚手架
  - go\_loop\_worker: 循环任务脚手架，适合于轮训场景
  - go\_crontab\_worker: 定时任务脚手架，适合于每分钟或每小时执行一次的场景
  - go\_rpcx\_server: RPCX的服务端脚手架。RPCX是一个类Dubbo的RPC框架，可以参考：https://github.com/smallnest/rpcx.
