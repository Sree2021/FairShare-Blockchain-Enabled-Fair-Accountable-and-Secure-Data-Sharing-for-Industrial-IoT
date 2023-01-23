pragma solidity ^0.5.0;

contract Store {

	bytes storedParams;
	bytes stored_g;
	bytes storedPK_C;
	bytes storedPK_F;
	string storedDId_M1;
	string stored_HashofF;
	bytes stored_HashofC;
	string storedDId_M2;
	string stored_HashofRC;
	string accessP;
	bytes storedNonce;

	function setParams(bytes memory input1, bytes memory input2) public {
		storedParams = input1;
		stored_g = input2;
	}

	function getParams() public view returns (bytes memory, bytes memory) {
		return (storedParams, stored_g);
	}

	function setPKC(bytes memory input) public {
		storedPK_C = input;
	}

	function getPKC() public view returns (bytes memory) {
		return storedPK_C;
	}

	function setPKF(bytes memory input) public {
		storedPK_F = input;
	}

	function getPKF() public view returns (bytes memory) {
		return storedPK_F;
	}

	function setMeta1(string memory input1, string memory input2, bytes memory input3) public {
		storedDId_M1 = input1;
		stored_HashofF = input2;
		stored_HashofC = input3;
	}

	function compareMeta1(string memory req) public view returns (bool) {
		return (keccak256(abi.encodePacked((stored_HashofF))) == keccak256(abi.encodePacked((req))) );
	}

	function setMeta2(string memory input1, string memory input2) public {
		storedDId_M2 = input1;
		stored_HashofRC = input2;
	}

	function compareMeta2(string memory req) public view returns (bool) {
		return (keccak256(abi.encodePacked((stored_HashofRC))) == keccak256(abi.encodePacked((req))) );
	}

	function setAccessP(string memory input1) public {
		accessP = input1;
	}
	function compareAccessPolicy(string memory req) public view returns (bool) {
		return (keccak256(abi.encodePacked((accessP))) == keccak256(abi.encodePacked((req))) );
	}

	function setNonce(bytes memory input) public {
		storedNonce = input;
	}

	function getNonce() public view returns (bytes memory) {
		return storedNonce;
	}
}
