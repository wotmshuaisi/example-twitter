# example-twitter
openfaas example function that use to send tweet

## deploy on kubeless

```shell
faas-cli build -f sendtweet.yml
faas-cli deploy -f sendtweet.yml
```

| params | describe |
| --- | --- |
| CKEY | API key |
| CSECRET | API secret key |
| ATOKEN | Access token |
| ASECRET | Access token secret |