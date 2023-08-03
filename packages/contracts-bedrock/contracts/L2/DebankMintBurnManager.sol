// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Semver } from "../universal/Semver.sol";

/**
 * @title DebankMintBurnManager
 */

contract DebankMintBurnManager is Semver {
    /// @notice Emitted when the owner of this contract changes.
    /// @param oldOwner Address of the previous owner.
    /// @param newOwner Address of the new owner.
    event OwnerChanged(address oldOwner, address newOwner);

    /// @notice Emitted when a minter is added.
    /// @param account Address of the minter.
    event MinterAdded(address indexed account);

    /// @notice Emitted when a minter is removed.
    /// @param account Address of the minter.
    event MinterRemoved(address indexed account);

    /// @notice Emitted when a mint is executed.
    /// @param to Address of the receiver.
    /// @param value Amount of tokens minted.
    /// @param l1TxId L1 tx id to debank's recharge address.
    event Mint(address indexed to, uint256 value, bytes32 l1TxId);

    /// @notice Emitted when a withdraw is executed.
    /// @param from Address of the sender.
    /// @param l1ChainId L1 chain id.
    /// @param l1TokenId L1 token id.
    event Withdraw(address indexed from, uint256 l1ChainId, address l1TokenId);

    /// @notice owner of this contract.
    address owner;

    /// @notice Blocks functions to anyone except the contract owner.
    modifier onlyOwner() {
        require(
            msg.sender == owner,
            "DebankMintBurnManager: function can only be called by the owner"
        );
        _;
    }

    /// @notice set owner of this contract.
    function setOwner(address _owner) external onlyOwner {
        require(
            _owner != address(0),
            "DebankMintBurnManager: owner can not be zero address"
        );

        emit OwnerChanged(owner, _owner);
        owner = _owner;
    }

    /// @custom:semver 1.0.1
    /// @notice Constructs the DebankL2Register contract.
    constructor() Semver(1, 0, 1) {}

    /// @notice minter address -> isMinter
    mapping (address => bool) _minters;

    /// @notice add minter
    function addMinter(address account) public onlyOwner {
        require(account != address(0) , "DebankMintBurnManager: minter can not be zero address");
        _minters[account] = true;
        emit MinterAdded(account);
    }

    /// @notice remove minter
    function removeMinter(address account) public onlyOwner {
        require(account != address(0) , "DebankMintBurnManager: minter can not be zero address");
        _minters[account] = false;
        emit MinterRemoved(account);
    }

    /// @notice Blocks functions to anyone except the minter.
    modifier onlyMinters() {
        require(
            _minters[msg.sender],
            "DebankMintBurnManager: caller does not have the Minter role"
        );
        _;
    }

    /// @notice mint token to ${to} with ${value} and ${l1TxId}
    /// @param to Address of the receiver.
    /// @param value Amount of tokens minted.
    /// @param l1TxId L1 tx id to debank's recharge address.
    function mint(address to, uint256 value, bytes32 l1TxId) onlyMinters external {
        // from地址为合约地址，该合约地址被设置为转账不扣余额
        (bool success, ) = to.call{value: value}("");
        require(success, "transfer in mint function failed.");
        emit Mint(to, value, l1TxId);
    }

    /// @notice withdraw token to ${l1ChainId} with ${l1TokenId}
    /// @param l1ChainId L1 chain id.
    /// @param l1TokenId L1 token id.
    function withdraw(uint256 l1ChainId, address l1TokenId) public payable {
        require(msg.value > 0, "no token withdraw with transaction");
        address burnAddress = address(0);
        (bool success, ) = burnAddress.call{value: msg.value}("");
        require(success, "transfer in withdraw function failed.");
        emit Withdraw(msg.sender, l1ChainId, l1TokenId);
    }
}
