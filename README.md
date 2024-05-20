# FlyBird  Restful

FlyBird  Restful provide restful method access MySQL database with json data.

## Quickstart

download executable and run it , it's start working!

### download executable file
[Download from Github](https://github.com/youwen21/flybird-restful/releases)

### start server

```bash
MYSQL_HOST=127.0.0.1 \
MYSQL_PORT=3306 \
MYSQL_USERNAME="valueN" \
MYSQL_PASSWORD="valueN" \
MYSQL_CHARSET=utf8mb4 \
MYSQL_DATABASE="valueN" \
./gofly 

```

## Access Data Apis

### Query
```bash
GET {{host}}/restful/:tableName
```

### Get
```bash
GET {{host}}/restful/:tableName/:id
```

### Insert
```bash
PUT {{host}}/restful/:tableName/

{
  "params":{
    "key1":"value1",
    "key2":"value2"
  }
}
```

### Update
```bash
POST {{host}}/restful/:tableName/:id

{
  "params":{
    "key1":"value1",
    "key2":"value2"
  }
}
```

### Delete
```bash
DELETE {{host}}/restful/:tableName/:id
```

## Advanced Usage
[Arguments](doc/arguments.md)

[Config](doc/config.md)

[Restful API](doc/restful-api)

[Advanced API](doc/advanced-api)

[Docker](doc/docker.md)

HAVE A GOOD DAY.


