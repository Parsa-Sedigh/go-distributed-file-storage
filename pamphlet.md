In golang, put the mutex above the thing that it's gonna protect from race conditions(using lock)

Organize public funcs at the top of the file and the private funcs at the bottom.

**Note:** The transport shouldn't maintain the map of Peers. Because we could have peers that are connected via
tcp or grpc or websocket or ... . Instead, the server should maintain this map. So the **server** should hold any kind
of peers with different transports that they are using.

- Storage: Can store files on disk based on the transformFunc.
- server: will be running as a daemon and it will receive commands

For security reasons, we shouldn't directly store the key names(shouldn't be stored in plain text) that the clients give 
to us. We should hash them, because the key could have a sensitive word in it!