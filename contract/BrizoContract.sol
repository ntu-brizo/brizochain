pragma solidity 0.4.26;

contract Brizo {
    mapping(address => string[]) storedData;
    mapping(string => string) hashDict;
    uint256 dataCounter = 0;

    function getCounter() public view returns (uint256) {
        return dataCounter;
    }

    function writeData(string memory content) public {
        require(bytes(content).length > 0, "The content of data is empty!");
        storedData[msg.sender].push(content);
        dataCounter += 1;
    }

    function readData(address addr, uint index) public view returns (string memory) {
        require(storedData[addr].length > 0, "No data was written by this address!");
        require((index >= 0) && (index < storedData[addr].length), "Index out of range!");
        return (storedData[addr][index]);
    }

    function writeDataToHashDict(string memory hashKey, string memory content) public {
        require(bytes(hashKey).length > 0, "No hash key specified!");
        require(bytes(content).length > 0, "Content is empty!");
        hashDict[hashKey] = content;
        dataCounter += 1;
    }

    function readDataFromHashDict(string hashKey) public view returns (string memory) {
        require(bytes(hashDict[hashKey]).length > 0, "No data was written using this hash key!");
        return (hashDict[hashKey]);
    }
}