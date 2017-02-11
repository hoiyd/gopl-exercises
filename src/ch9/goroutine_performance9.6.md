#### Number of CPU Cores on my computer
```
~ ‹ruby-2.1.5›  $ sysctl -n hw.ncpu
8
```

#### Performance and number of workers(goroutines)

```
~/Projects/gopl-exercises ‹ruby-2.1.5›  ‹master*› $ go run src/ch8/parallell_computing_tasks.8.5.go

1 workers initialized.
Paralell verison rendered in: 3.656222324s

2 workers initialized.
Paralell verison rendered in: 1.932233143s

3 workers initialized.
Paralell verison rendered in: 1.394129903s

4 workers initialized.
Paralell verison rendered in: 1.143849746s

5 workers initialized.
Paralell verison rendered in: 1.096122239s

6 workers initialized.
Paralell verison rendered in: 989.590388ms

7 workers initialized.
Paralell verison rendered in: 895.958012ms

8 workers initialized.
Paralell verison rendered in: 826.907382ms

9 workers initialized.
Paralell verison rendered in: 839.575649ms

10 workers initialized.
Paralell verison rendered in: 859.062398ms

11 workers initialized.
Paralell verison rendered in: 894.139143ms

12 workers initialized.
Paralell verison rendered in: 843.492512ms

13 workers initialized.
Paralell verison rendered in: 865.89065ms

14 workers initialized.
Paralell verison rendered in: 883.644449ms

15 workers initialized.
Paralell verison rendered in: 927.749351ms

16 workers initialized.
Paralell verison rendered in: 921.757431ms
```

#### Conclusion
When the number of goroutines exceeds the # of CPU cores(which is 8), there's more performance boost.

So the optimal value is GOMAXPROCS=8 on my 8-core mac.