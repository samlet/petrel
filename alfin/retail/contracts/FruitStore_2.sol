pragma solidity ^0.4.25;

import "./BasicAuth.sol";

contract FruitStore_2 is BasicAuth {
    mapping(bytes => uint) _fruitStock;

    modifier validFruitName(bytes fruitName) {
        require(fruitName.length > 0, "fruite name is invalid!");
        _;
    }
    function setFruitStock(bytes fruitName, uint stock)
    onlyOwner validFruitName(fruitName) external {
        _fruitStock[fruitName] = stock;
    }
}
