# Couchbox Server \u2014 Go Port

Functional port of `codefodder/couchbox-server` to Go.

The goal is behavioral compatibility with the existing client-facing server,
while allowing the Go implementation to differ internally from the Python
implementation.

## Project Skeleton

- [x] Create Go module
- [x] Create command package
- [x] Create internal package structure
- [x] Create adjacent Go test files
- [x] Establish green `go test ./...` baseline
- [ ] Add project documentation
- [ ] Add build/run targets to `Makefile`

## Filesystem

- [x] Identify filesystem behaviors from the Python implementation
- [x] Define filesystem package API
- [ ] Implement filesystem root handling
- [ ] Implement recursive file discovery
- [ ] Define filesystem error behavior
- [ ] Add filesystem edge-case tests

## Configuration

- [x] Identify configuration schema from Couchbox
- [x] Define Go configuration types
- [ ] Implement configuration loading
- [ ] Implement configuration defaults
- [ ] Test invalid configuration
- [ ] Test missing configuration
- [ ] Test configuration compatibility

note: 

config.VideoExts
       │
       ├── filesystem
       ├── listings
       └── any other extension-sensitive behavior

## Listings

- [ ] Identify listing model and behavior
- [ ] Define Go listing types
- [ ] Implement listing discovery
- [ ] Implement listing metadata
- [ ] Implement listing filtering/selection
- [ ] Add representative listing fixtures
- [ ] Add listing edge-case tests

## Locations

- [ ] Identify location model and behavior
- [ ] Define Go location types
- [ ] Implement location discovery
- [ ] Implement location handling
- [ ] Add representative location fixtures
- [ ] Add location edge-case tests

## Posters

- [ ] Identify poster behavior required by the client
- [ ] Define poster representation
- [ ] Create PNG fixture set
- [ ] Add poster fixtures to `testdata`
- [ ] Implement poster lookup/serving behavior
- [ ] Add poster tests

> Poster generation is fixture preparation only. There is no runtime
> poster-generation subsystem.

## Bundles

- [ ] Identify bundle structure and behavior
- [ ] Define Go bundle types
- [ ] Implement bundle assembly
- [ ] Add representative bundle fixtures
- [ ] Test bundle contents
- [ ] Test bundle edge cases

## HTTP Server

- [ ] Identify all routes from the Python implementation
- [ ] Identify HTTP methods
- [ ] Identify request formats
- [ ] Identify response formats
- [ ] Identify status/error behavior
- [ ] Implement routing
- [ ] Implement handlers
- [ ] Add handler tests
- [ ] Add HTTP integration tests

## Compatibility

- [ ] Document client-facing API contract
- [ ] Create compatibility fixtures
- [ ] Compare representative Python responses
- [ ] Compare representative Go responses
- [ ] Test filesystem/configuration compatibility
- [ ] Test listing compatibility
- [ ] Test location compatibility
- [ ] Test poster compatibility
- [ ] Test bundle compatibility
- [ ] Test HTTP compatibility

## Runtime

- [ ] Implement server startup
- [ ] Implement configuration loading at startup
- [ ] Implement graceful shutdown
- [ ] Implement appropriate logging
- [ ] Verify clean startup with representative configuration

## Quality

- [ ] `go test ./...` green
- [ ] `go vet ./...` clean
- [ ] `gofmt` clean
- [ ] Add useful test fixtures
- [ ] Remove unnecessary dependencies
- [ ] Review public/internal package boundaries
- [ ] Review error handling
- [ ] Review compatibility against original server

## Final

- [ ] Build production binary
- [ ] Document configuration
- [ ] Document running the server
- [ ] Document client compatibility
- [ ] Document intentional functional differences
- [ ] Final full test suite

