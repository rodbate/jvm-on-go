## Jvm On Go

使用golang基于[JVM8 Spec](https://docs.oracle.com/javase/specs/jvms/se8/html/index.html) 构建的简易JVM，仅供学习使用。

### Roadmap
+ [x] 类加载
+ [x] 字节码解析
+ [x] 类解析及初始化
+ [x] 字节码解释器
+ [x] 单线程
+ [x] 部分Native方法(使用golang实现)
+ [ ] invokedynamic
+ [ ] Lambda

### Getting Started

#### 1. 环境准备
+ 本地需要Java8运行环境并且设置`JAVA_HOME`环境变量
+ 若是通过源码构建，本地需要安装golang(>= 1.15)

#### 2. 源码构建
```shell
git clone https://github.com/rodbate/jvm-on-go.git
cd jvm-on-go && go build -o jvm
```

#### 3. 下载二进制包
```shell
#windows
wget https://github.com/rodbate/jvm-on-go/releases/download/1.0.0/jvm-windows

#linux
wget https://github.com/rodbate/jvm-on-go/releases/download/1.0.0/jvm-linux

#mac os
wget https://github.com/rodbate/jvm-on-go/releases/download/1.0.0/jvm-darwin
```
> 以上都是基于64位系统


#### 4. 使用
在`testdata`目录下有几个简单的java示例

比如：

```java
//ChineseSupportSample.java
//ChineseSupportSample.class
public class ChineseSupportSample {
    public static void main(String[] args) {
        for (int i = 0; i < 1000; i++) {
            System.out.println("中文 - Hello world: " + i);
        }
    }
}
```

输入以下命令即可运行
```shell
./jvm -cp testdata ChineseSupportSample
```

