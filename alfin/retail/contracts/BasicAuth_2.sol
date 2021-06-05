pragma solidity ^0.4.25;

contract BasicAuth_2 {
    address public _owner;
    address public _bakOwner;

    constructor(address bakOwner) public {
        _owner = msg.sender;
        _bakOwner = bakOwner;
    }

    function setOwner(address owner)
    public
    isAuthorized
    {
        _owner = owner;
    }

    function setBakOwner(address owner)
    public
    isAuthorized
    {
        _bakOwner = owner;
    }

    modifier isAuthorized() {
        require(msg.sender == _owner || msg.sender == _bakOwner, "BasicAuth: only owner or back owner is authorized.");
        _;
    }

    modifier onlyOwner() {
        require(msg.sender == _owner, "BasicAuth: only owner is authorized.");
        _;
    }
}
