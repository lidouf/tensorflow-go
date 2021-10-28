# Tensorflow Go Bindings
This is a fork of the official tensorflow's repo(currently at version 2.6) but with working go bindings only. 
I suppose you've already heard of they'll no longer maintain the tensorflow go bindings.
I don't know whether I should feel happy or not, since the official tf-go bindings hasn't been working for me for long time, 
especially if you are working in go module. All in all, it works like ****.
I add `go.mod` in the repo and auto-generate pb file so that you can use it directly(in go module).

## Install
If you're using go module, just use the `replace` instruction like this:
```
replace github.com/tensorflow/tensorflow v2.6 => github.com/lidouf/tensorflow-go v2.6
```

## License

[Apache License 2.0](LICENSE)
