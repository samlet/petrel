pragma solidity ^0.4.25;

contract Authority {
    function canCall(
        address src, address dst, bytes4 sig
    ) public view returns (bool);
}
