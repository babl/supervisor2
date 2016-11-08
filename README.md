# Supervisor II

Its the service responsible for starting and supervising babl modules on the platform, being the man in the middle between babl client and babl server.


- Starts A gprc server
- Connects to kafka, as client and producer
- consumes from kafka topic:modules the modules available
- for each module available creates a mew bablmodule
- opens a channel for each module, and listens to modules responses, on topic supervisor.hostname
- call Gprc.serve to start listening for RPC on a tcp port

- IO and Ping are the Process functions for RPC calls from bablmodule, each of them make a request to , that publish the message on module topic, and if async, returns, if not waits for a response on the channel, timing out after ModuleExecutionWaitTimeout


## BABL SERVER

Running on every docker instance its the common interface to interact with your module

