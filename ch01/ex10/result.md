報告される時間が大きく変化することはなかった。

```
kirarin% ./fetchall http://www.yahoo.co.jp/ http://www.google.com/ http://www.amazon.co.jp/
0.78s    19412  http://www.yahoo.co.jp/
1.01s    19389  http://www.google.com/
3.00s   515592  http://www.amazon.co.jp/
3.00s elapsed
```

```
kirarin% ./fetchall http://www.yahoo.co.jp/ http://www.google.com/ http://www.amazon.co.jp/
1.45s    19413  http://www.yahoo.co.jp/
4.73s   524339  http://www.amazon.co.jp/
5.05s    19339  http://www.google.com/
5.05s elapsed
```

```
kirarin% ./fetchall http://www.yahoo.co.jp/ http://www.google.com/ http://www.amazon.co.jp/
0.23s    19413  http://www.yahoo.co.jp/
0.46s    19315  http://www.google.com/
4.67s   494782  http://www.amazon.co.jp/
4.67s elapsed
```
