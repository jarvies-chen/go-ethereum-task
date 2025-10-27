// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Counter {
    uint256 private count;
    address public owner;

    event CountIncremented(address indexed caller, uint256 newValue);
    event CountReset(address indexed caller);

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }

    constructor() {
        owner = msg.sender;
        count = 0;
    }

    function increment() public {
        count++;
        emit CountIncremented(msg.sender, count);
    }

    function getCount() public view returns (uint256) {
        return count;
    }

    function reset() public onlyOwner {
        count = 0;
        emit CountReset(msg.sender);
    }

    function setCount(uint256 _count) public onlyOwner {
        count = _count;
    }
}