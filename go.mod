module hello-client-go

go 1.22

require (
	adaptivemsg v0.0.0
	hello-server-rust v0.0.0
)

require (
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
)

replace adaptivemsg => ../adaptivemsg-go
replace hello-server-rust => ../hello-server-rust
