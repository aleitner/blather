# Blather [![PkgGoDev](https://pkg.go.dev/badge/github.com/aleitner/blather)](https://pkg.go.dev/github.com/aleitner/blather)
Library for making group calls that have spacial awareness

## Architecture

### Diagram

![](https://github.com/aleitner/blather/blob/master/diagram.png?raw=true)

## TODO

#### Library:
- [x] Stream audio to server
- [x] Server forwards recieved audio to clients
- [x] Combine all input sources into single stream
- [x] Client resamples audio before sending out
- [x] Add rooms to server for different chat groups
- [ ] Add rate limiter
- [ ] Send coordinates to server
- [ ] Retrieve coordinates from server
- [ ] Adjust sample rate and quality by connection
- [ ] Message signing and encryption
- [ ] Switch server host when server goes down
- [ ] Write Integration Tests
- [ ] Address all NB's and Todo's

#### Application:
- [ ] Front end
- [ ] Change volume depending on distance
- [ ] Pan audio depending on location
- [x] Get Mic Input [aleitner/microphone](https://github.com/aleitner/microphone)
