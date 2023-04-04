# Simple hash

Ported/based on Apache Shiro [default hash service](https://shiro.apache.org/static/1.11.0/apidocs/org/apache/shiro/crypto/hash/DefaultHashService.html).

## Features

- `SHA-512` algorithm only
- reading data to be hashed from pipe or stdin (prompt)
- configurable iterations, defaults to `1`
- generating random salt, if none provided
- providing a private salt
- salt & private salt need to be base64 encoded, when provided
- computed hash is base64 encoded

## Build

```shell
go build -o hasher ./main.go
```

## Usage

Prompt:
```shell
; ./hasher
Data to hash: hello
m3HSJL1i83hdltRq0+o9czGb+8KJDKra4t/3JRlnPKcjI8PZm6XBHXx6zG4UuMXaDEZjR1wuXDre9G9zvN7AQw==
```

Pipe:
```shell
; echo -n "hello" | ./hasher
m3HSJL1i83hdltRq0+o9czGb+8KJDKra4t/3JRlnPKcjI8PZm6XBHXx6zG4UuMXaDEZjR1wuXDre9G9zvN7AQw==
```

More iterations:
```shell
; echo -n "hello" | ./hasher --iterations 1000
Lw0Vk8nYLMBpNeaeTosjDbEkD2Olrnl30ahRc5eJX3JyVExUEGGN7qtvZyDDOzs6c+oERcchJ76IFLtNB/9hbg==
```

Provide a salt, no random generation:
```shell
; echo -n "hello" | ./hasher --salt $(echo -n "my salt" | base64)
Llq3/XC0PWKCZ75IR5gc/wJ2Yxhc+7ijWWdWoBc46lHeMpyimQnWgegmQT9T//uSEmfoCPXBCWziPfbxlO7d8w==
```

Provide a private salt, extra random salt generated:
```shell
; echo -n "hello" | ./hasher --private-salt $(echo -n "my private salt" | base64)
Z1/zSYHgQSRDFLGJHPdFcyGFom9X2DSSvjJYp4w9TSfGa+h6jgbFlnEkaNHu847BlYhnv7jXTAH+Qf02cIh0Yw==
```

Provide a private salt, no random salt generated:
```shell
; echo -n "hello" | ./hasher --private-salt $(echo -n "my private salt" | base64) --salt $(echo -n "my salt" | base64)
YaB1xuISPQpzOVm1PKmx8ixb176/11YsPgX2PrHcqeHgXHG55+DKFGi5mJU3ZrLwbLVLUxLdvvb+jfJbiUKWZw==
```
