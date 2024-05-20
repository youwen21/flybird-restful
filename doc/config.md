# config file

you can create a toml config file in work dir named config, below demo:

```toml

[mysql_default]
host = "127.0.0.1"
username = "root"
password = "123456"
port = "3306"
database = ""
charset = "utf8"

[redis_default]
host = ""
port = "3306"
password = "123456"
db = 1

[smtp]
host = ""
port = 465
user = ""
password = ""

```

Flybird Restful load the config.yaml by default.

if you specify runmode argument, it will load corresponding config file

```cgo
./gofly --runmod="prod"
```

it attempt load config_prod.toml file in the work dir.



