### Ex1.10
```shell
$ go run ch1/ex1.10-11/fetchall1_10.go http://alexa.com

2.96s    25380  http://alexa.com

2.96s elapsed
```

```shell
$ go run ch1/ex1.10-11/fetchall1_10.go http://alexa.com

1.20s    25380  http://alexa.com

1.20s elapsed
```

Apparentlly, the two requests fetch the same content.

### Ex1.11
```
$ ./bin/fetchall http://google.com http://facebook.com http://youtube.com http://baidu.com http://yahoo.com http://amazon.com http://wikipedia.org http://qq.com http://twitter.com http://google.co.in http://taobao.com http://xxxxyyssshhhhhsssss.com http://live.com http://yahoo.co.jp http://linkedn.com
Get http://linkedn.com: dial tcp: lookup linkedn.com: no such host
Get http://xxxxyyssshhhhhsssss.com: dial tcp: lookup xxxxyyssshhhhhsssss.com: no such host
0.03s       81  http://baidu.com
0.08s   617243  http://qq.com
0.23s   148655  http://taobao.com
0.59s    19191  http://yahoo.co.jp
1.13s    19047  http://google.com
1.19s    21353  http://google.co.in
2.22s   247160  http://twitter.com
2.92s   438140  http://youtube.com
2.92s    54537  http://wikipedia.org
3.21s    69574  http://facebook.com
3.40s   458776  http://yahoo.com
3.91s   383495  http://amazon.com
4.07s     9672  http://live.com
4.07s elapsed
```

When a website just doesn't respond, it'll timeout after 30s.