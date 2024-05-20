
# Restful Api full parameters usage

use more complex condition see [advanced api](advanced-api.md)

### Query

parameters

| parameter | type | default |
|-----------|------|---------|
| page      | int  | 1       |
| psize     | int  | 20      |
| orderBy     | int  | -       |
| orderDirect     | int  | -       |


```bash
GET {{host}}/restful/:tableName?page=2&psize=10&orderBy="id"&orderDirect="asc"
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