# BitFlux

BitFlux is a lightweight library designed to simply custom packet generation and binary serialization.
With a focus on flexibility and ease of use, BitFlux provides low-level control over binary data layouts while maintaining focusing on a straightforward API.

## Goals

- **Intuitive API**: Read and Write data with simple methods like `WriteUint32`, `ReadFloat32`
- **Endianness Support**: Seamlessly handle both little-endian and big-endian formats for all supported types. Separate functions are provided for each format, allowing packets containing fields of both formats.
- **Custom Serialization**: Tailored for custom binary serialization purposes.
- **Lightweight Design**: No external dependencies.

## Installation

```
go get github.com/jon-ski/bitflux
```

## Quick Start

```go
package main

import (
  "fmt"
  "log"
  
  "github.com/jon-ski/bitflux"
)

func main() {
  var buf bitflux.Buffer
  buf.WriteLUint16(1_234) // Little-Endian 16-bit 1234 (d2 04)
  buf.WriteBUint32(78_910) // Little-Endian 32-bit 78910 (00 01 34 3e)
  if buf.Err() != nil {
    // buf will keep track of the last error and will
    // nop on new calls if an error has occurred
    // 
    // Alternatively, functions that can error return an error.
    log.Fatalf("unexpected error: %w", buf.Err())
  }
  fmt.Printf("% 2x\n", buf.Bytes())
  // d2 04 00 01 34 3e
}
```
