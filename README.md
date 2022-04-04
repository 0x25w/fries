# fries
finger-scan

## 基础用法
1.识别单条url指纹

```fries -u http://www.xxx.com```

2.识别一个url list文档的指纹

```fries -l url.txt```

3.使用fofa查询并查看结果数据的指纹

```fries -f "ip=x.x.x.x"```

4.使用zoomeye查询并查看指纹

```fries -z "ip:"8.8.8.8""```

## 通用参数

1.输出到{文件名},会自动补充后缀.txt

```-o filename```

```-o ../filename```

```-o c:/filename```

2.扫描线程,默认为5，1-30范围内

```-t 5```

