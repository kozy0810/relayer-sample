// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { ERC20 } from "solmate/src/tokens/ERC20.sol";

import "hardhat/console.sol";

contract ERC20Token is ERC20 {
    constructor() ERC20("MyToken", "MTK", 18) {
        _mint(msg.sender, 1000e18);
        console.log(msg.sender, "balance is: ", balanceOf[msg.sender]);
        // transfer(0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266, 10000 ether);
    }
}