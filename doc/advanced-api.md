# Advanced Api usage

Use advanced Api when you need complex query

### Query

```bash
POST {{host}}/access_db/query

{
  json body request conditions below
}

```


#### request raw sql condition
```bash

{
  	"OrderBy":    "id",
  	"OrderDirect":    "desc",
  	"Page":    1,
  	"Psize":    20,
  	"Table":"my_table",
  	"condition":{
  	  "raw":true,
  	  "rawCondition":"name='bob' and status=2"
  	}
	
}
```

#### request description items condition
```bash

{
  	"OrderBy":    "id",
  	"OrderDirect":    "desc",
  	"Page":    1,
  	"Psize":    20,
  	"Table":"my_table",
  	"condition":{
  	  "raw":false,
  	  "items": [
  	    "filed_name":"name",
  	    "compare":"is equal to",
  	    "value":"alice",
  	  ]
  	}
	
}
```

##### compare list
```cgo

case "is equal to":
case "is nor equal to":
case "is less then":
case "is less then or equal to":
case "is 
case "is greater then or equal to":
case "contains":
case "does not contain":
case "begin with":
case "does not begin with":
case "end with":
case "does not end with":
case "is null":
case "is not null":
case "is empty":
case "is not empty":
case "is between":
case "is not between":
case "is in list":
case "is not in list":
			
```


