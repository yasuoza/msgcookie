# Conversation server example
To write conversation server.

## How to run

Open two terminal and run respectively.

**Ruby**

http://localhost:4567

```
$ gem install sinatra msgpack
$ ruby session_count_server.rb
```

**Go**

http://localhost:8080

```
$ go get github.com/yasuoza/msgcookie
$ go run conversation_server.go
```

## Common configuration

Defined in `config.json`

```js
{
  "cookie_name":"rack.session",
  "cookie_secret":"special_secret"
}
```
