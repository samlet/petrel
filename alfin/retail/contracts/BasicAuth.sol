pragma solidity ^0.4.25;

contract BasicAuth {
    address public _owner;

    constructor() public {
        _owner = msg.sender;
    }

    function setOwner(address owner)
    public
    onlyOwner
    {
        _owner = owner;
    }

    modifier onlyOwner() {
        require(msg.sender == _owner, "BasicAuth: only owner is authorized.");
        _;
    }
}
