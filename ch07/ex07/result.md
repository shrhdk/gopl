# Result

https://github.com/golang/go/blob/5fea2ccc77eb50a9704fa04b7c61755fe34e1d95/src/flag/flag.go#L765-L766

> // Remember the default value as a string; it won't change.
> flag := &Flag{name, usage, value, value.String()}
