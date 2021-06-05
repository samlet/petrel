pragma solidity ^0.4.25;

contract FruitStore {
    address public  _owner;
    // 对店内所有水果品类和库存数量的管理
    mapping(bytes => uint) _fruitStock;
    modifier validFruitName(bytes fruitName) {
        require(fruitName.length > 0, "fruite name is invalid!");
        _;
    }
    // 鉴权函数修饰器
    modifier onlyOwner() {
        require(msg.sender == _owner, "Auth: only owner is authorized.");
        _;
    }

    /**
     * @dev 查询具体水果库存数量的函数
     */
    function getStock(bytes fruit) external view returns(uint) {
        return _fruitStock[fruit];
    }

    function setFruitStock(bytes fruitName, uint stock)
        onlyOwner validFruitName(fruitName) external {
        // 只允许合约部署者来调用 setFruitStock 函数
        _fruitStock[fruitName] = stock;
    }
}

