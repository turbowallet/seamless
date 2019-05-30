pragma solidity ^0.5.7;

contract Incrementor {

    uint64 counter;

    /* This runs when the contract is executed */
    constructor() public {
        counter = 0;
    }

    function increment() public returns (uint64) {
        counter++;
        return counter;
    }

    function count() public view returns (uint64) {
        return counter;
    }

}
