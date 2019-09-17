pragma solidity ^0.5.11;

contract Brizo {
    mapping(address => string[]) storedData;

    function dataCounter(address addr) public view returns (uint256) {
        return storedData[addr].length;
    }
    
    function writeData(string memory data) public {
        require(bytes(data).length > 0, "The content of data is empty!");
        storedData[msg.sender].push(data);
    }
    
    function readData(address addr, uint index) public view returns (string memory) {
        require(storedData[addr].length > 0, "No data was written by this address");
        require((index >= 0) && (index < storedData[addr].length), "Index out of range!");
        return (storedData[addr][index]);
    }
}
