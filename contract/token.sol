pragma solidity ^0.8.0;

import "./core/token/ERC20/ERC20.sol";

contract Token is ERC20 {
    constructor(uint256 initialSupply) ERC20("LI LU YANG TOKEN", "LLY") {
        _mint(msg.sender, initialSupply);
    }
}