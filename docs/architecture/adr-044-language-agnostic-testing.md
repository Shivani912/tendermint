# ADR 044: Language Agnostic Testing

## Context

There is no existing system that can check behaviour of multiple implementations of code other than manual code comparison. This is a very risky situation as we dont want the same functionality to have slightest difference in behaviour on different code implementations.

As we are planning to move forward with building Tendermint in Rust and eventually having it production ready, it is one of the most crucial tasks to ensure the Rust version is 100% functionally compliant with the Go version.
 
## Decision

Language Agnostic Test Suites using JSON files and verifying those against different versions of code will help test Tendermint's functionality and compatibility across different implementations.

![Test flow diagram](img/Language-Agnostic-Testing.png)

Generators in Go and Rust will produce json test files. These files can then be passed through tests to check whether different implementations result into same result given the same condition.

Why we need generators and tests in both languages?
- generators show the ability to produce expected behaviour
- tests show the ability to detect faulty behaviour and accept expected behaviour

## Status

Accepted

### Positive

Behaviour control across multiple implementations

