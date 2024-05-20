# Arguments

## arguments list

```text
	host = flag.String("host", "0.0.0.0", "listen host.")
	port = flag.String("port", "80", "Port to run the server.")

	root = flag.String("root", "", "specify the root.")

	runmode = flag.String("runmode", "", "runmode default is empty")

	logfile  = flag.String("logfile", "", "logfile default os stdout")
	loglevel = flag.String("loglevel", "error", "log level")
```

## Demo

```bash

./gobly -port=8888

```