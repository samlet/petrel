pragma solidity ^0.4.25;

import "./BasicAuth.sol";

/**
 * 经典的 ACL(Access Control List) 设计
 */
contract AclGuard is BasicAuth {
    bytes4 constant public ANY_SIG = bytes4(uint(- 1));
    address constant public ANY_ADDRESS = address(bytes20(uint(- 1)));
    mapping(address => mapping(address => mapping(bytes4 => bool))) _acl;

    function canCall(
        address src, address dst, bytes4 sig
    ) public view returns (bool) {
        return _acl[src][dst][sig]
        || _acl[src][dst][ANY_SIG]
        || _acl[src][ANY_ADDRESS][sig]
        || _acl[src][ANY_ADDRESS][ANY_SIG]
        || _acl[ANY_ADDRESS][dst][sig]
        || _acl[ANY_ADDRESS][dst][ANY_SIG]
        || _acl[ANY_ADDRESS][ANY_ADDRESS][sig]
        || _acl[ANY_ADDRESS][ANY_ADDRESS][ANY_SIG];
    }

    function permit(address src, address dst, bytes4 sig) public onlyOwner {
        _acl[src][dst][sig] = true;
        // emit LogPermit(src, dst, sig);
    }

    function forbid(address src, address dst, bytes4 sig) public onlyOwner {
        _acl[src][dst][sig] = false;
        // emit LogForbid(src, dst, sig);
    }

    function permit(address src, address dst, string sig) external {
        permit(src, dst, bytes4(keccak256(sig)));
    }

    function forbid(address src, address dst, string sig) external {
        forbid(src, dst, bytes4(keccak256(sig)));
    }

    function permitAny(address src, address dst) external {
        permit(src, dst, ANY_SIG);
    }

    function forbidAny(address src, address dst) external {
        forbid(src, dst, ANY_SIG);
    }
}
