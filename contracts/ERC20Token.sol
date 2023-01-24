// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/draft-ERC20Permit.sol";

import "hardhat/console.sol";

contract ERC20Token is ERC20, ERC20Permit {
    constructor() ERC20("MyToken", "MTK") ERC20Permit("MyToken Permit Message") {
        _mint(msg.sender, 10000 ether);
        transfer(0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266, 10000 ether);
    }
}