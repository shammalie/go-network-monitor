# go-network-monitor

A scalable server that receives pcap events via [go-network-monitor-client](https://github.com/shammalie/go-network-monitor-client). Pcap events are communicated to the server via gRPC, where the server can make decisions for each client to effect the pcap collection. e.g. if an ip is deemed to be creating too much traffic or attempting brute force attacks via ports, the server can issue a temporary or permanent block for said IP. MongoDB is the primary database for this service, and Redis for caching and channels.

![Diagram showing the process for go-network-monitor, version 1.](diagrams/network_monitor_v1.svg)

## Future work

- [ ] Add support for client api keys
- [ ] Create seperate gRPC streams for clients and groups of clients
- [ ] Create decisions for ip events
- [ ] Enable decision making for group of clients
- [ ] Add Tests and benchmarks
