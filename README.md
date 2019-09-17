# brizochain.go
### This package provides methods to interact with Brizo Chain
- NewBrizoChain
```go
  // You have to have a brizochain instance to interact with smart contract.
  func NewBrizoChain() (*brizoChain, error)
``` 
- Write
```go
  // Pass string data and store it to Brizo Chain
  func (brizo *brizoChain) Write(msg string) error
```
- Read
```go
  // Read returns data on Brizo Chain according to the given index (start at 0)
  func (brizo *brizoChain) Read(index int64) (string, error)
```
- ReadByAddress
```go
  // ReadByAddress returns data on Brizo Chain by passing address (common.Address) and index
  func (brizo *brizoChain) ReadByAddress(address common.Address, index int64) (string, error)
```
- ReadByHexAddress
```go
  // ReadByHexAddress returns data on Brizo Chain by passing hex address (string) and index
  func (brizo *brizoChain) ReadByHexAddress(hexAddress string, index int64) (string, error)
```

# brizochain_test.go
### Test `Write` and `Read` methods of brizochain.go
- `TestBrizoChain_Write` tests `Write` function of brizochain

- `TestBrizoChain_Read` tests `Read` function of brizochain

# contract /...
- BrizoContract.sol: BrizoContract written in solidity

- BrizoContract.go: BrizoContract Golang version