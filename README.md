# example-twitter
kubeless example function that use to send tweet

## deploy on kubeless

```shell
kubeless function deploy sendtweet --runtime go1.10 --handler kubeless.SendTweet --from-file example-twitter/function.go --dependencies example-twitter/Gopkg.toml --env AKEY:'',ASECRET:'',ATOKEN:'',ASECRET:''
```

| params | describe |
| --- | --- |
| CKEY | API key |
| CSECRET | API secret key |
| ATOKEN | Access token |
| ASECRET | Access token secret |