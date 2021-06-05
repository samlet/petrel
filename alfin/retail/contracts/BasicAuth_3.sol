pragma solidity ^0.4.25;

import "./Authority.sol";

contract BasicAuth_3 {
    event LogSetAuthority (Authority indexed authority, address indexed from);

    Authority  public  _authority;
    address public _owner;

    constructor() public {
        _owner = msg.sender;
    }

    function setAuthority(Authority authority) public {
        _authority = authority;
        // 当 setAuthority 函数被调用时，会同时触发 LogSetAuthority，
        // 将事件中定义的 Authority 合约地址以及调用者地址记录到区块链交易回执中。
        emit LogSetAuthority(authority, msg.sender);
    }

    modifier isAuthorized() {
        require(auth(msg.sender, msg.sig), "BasicAuth: only owner or back owner is authorized.");
        _;
    }

    function auth(address src, bytes4 sig) public view returns (bool) {
        if (src == address(this)) {
            return true;
        } else if (src == _owner) {
            return true;
        } else if (_authority == Authority(0)) {
            return false;
        } else {
            return _authority.canCall(src, this, sig);
        }
    }
}
