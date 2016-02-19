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