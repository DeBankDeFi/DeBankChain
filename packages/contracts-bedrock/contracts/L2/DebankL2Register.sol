// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Semver } from "../universal/Semver.sol";

/**
 * @title DebankL2Register
 */

contract DebankL2Register is Semver {
    /// @notice Address of the admin of this contract.
    address public owner;

    /// @notice Mapping of l1 addresses to registered l2 times.
    mapping(address => uint256) public nonces;

    /// @notice l1->l2, Mapping of l1 addresses to registered l2 accounts.
    mapping(address => address) public l1Tol2Accounts;

    /// @notice l2->l1, Mapping of l2 addresses to registered l1 accounts.
    mapping(address => address) public l2Tol1Accounts;

    /// @notice Emitted when a user registers an l2 account.
    /// @param user l1 Address of the user.
    /// @param l2Account l2 Address of the user.
    event Register(address user, address l2Account, uint256 registerCnt);

    /// @notice Emitted when the owner of this contract changes.
    /// @param oldOwner Address of the previous owner.
    /// @param newOwner Address of the new owner.
    event OwnerChanged(address oldOwner, address newOwner);

    /// @custom:semver 1.0.1
    /// @notice Constructs the DebankL2Register contract.
    constructor() Semver(1, 0, 1) {}

    /// @notice Registers an l2 account for the caller.
    /// @param l2Account l2 Address of the user.
    function register(address l2Account) public {
        require(l2Tol1Accounts[l2Account] == address(0), "DebankL2Register: This l2 account is already registered.");

        // Check if l1 account has been registered
        if(l1Tol2Accounts[msg.sender] != address(0)){
            // Clear the old l2 account associated with the l1 account
            delete l2Tol1Accounts[l1Tol2Accounts[msg.sender]];
        }

        l1Tol2Accounts[msg.sender] = l2Account;
        l2Tol1Accounts[l2Account] = msg.sender;
        nonces[msg.sender] += 1;
        emit Register(msg.sender, l2Account, nonces[msg.sender]);
    }

    /// @notice Blocks functions to anyone except the contract owner.
    modifier onlyOwner() {
        require(
            msg.sender == owner,
            "DebankL2Register: function can only be called by the owner of this contract"
        );
        _;
    }

    /// @notice Updates the owner of this contract.
    /// @param _owner Address of the new owner.
    function setOwner(address _owner) external onlyOwner {
        // Prevent users from setting the whitelist owner to address(0) except via
        // enableArbitraryContractDeployment. If you want to burn the whitelist owner, send it to
        // any other address that doesn't have a corresponding knowable private key.
        require(
            _owner != address(0),
            "DebankL2Register: can only be disabled via enableArbitraryContractDeployment"
        );

        emit OwnerChanged(owner, _owner);
        owner = _owner;
    }
}
