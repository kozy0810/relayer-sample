// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/draft-IERC20Permit.sol";

contract Bank is ReentrancyGuard {
    address _token;
    mapping(address => uint256) public _balance;

    event Deposit(address indexed _account, uint256 _amount);
    event Withdraw(address indexed _account, uint256 _amount);

    constructor(address token_) ReentrancyGuard() {
        _token = token_;
    }

    function depositWithSig(
        address _owner,
        address _spender,
        uint256 _value,
        uint256 _deadline,
        uint8 _v,
        bytes32 _r,
        bytes32 _s
    ) public nonReentrant {
        IERC20Permit(_token).permit(_owner, _spender, _value, _deadline, _v, _r, _s);

        _balance[_owner] = _value;
        IERC20(_token).transfer(address(this), _value);
        emit Deposit(_owner, _value);
    }

    function withdrawWithSig(
        address _owner,
        address _spender,
        uint256 _value,
        uint256 _deadline,
        uint8 _v,
        bytes32 _r,
        bytes32 _s
    ) public nonReentrant {
        require(_balance[_owner] >= _value, "insufficient funds");

        IERC20Permit(_token).permit(_owner, _spender, _value, _deadline, _v, _r, _s);
        IERC20(_token).transfer(address(this), _value);
        emit Withdraw(_owner, _value);
    }
}