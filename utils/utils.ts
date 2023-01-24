import { ethers } from "hardhat";
import { signTypedData, SignTypedDataVersion } from "@metamask/eth-sig-util";

export interface ERC2612PermitMessage {
  owner: string;
  spender: string;
  value: number | string;
  nonce: number | string;
  deadline: number | string;
}

// interface EIP712Domain {
//   name: string;
//   version: string;
//   chainId: number;
//   verifyingContract: string;
// }

const DomainName = "MyToken Permit Message";
const DomainVersion = "1";

export function getRelayerURL() {
  if (process.env.RELAYER_BASE_URL === "") {
    return "http://localhost:8080"
  } else {
    return process.env.RELAYER_BASE_URL
  }
}

export function signERC2612PermitMessage(privateKey: string, tokenAddress: string, chainId: number, message: any) {
  if (privateKey.indexOf('0x') === 0) {
    privateKey = privateKey.slice(2);
  }

  const { domain, types } = getERC2612PermitDomainDefinition(chainId, tokenAddress);
  return signTypedData({
    privateKey: Buffer.from(privateKey, "hex"),
    data: {
      types,
      primaryType: "Permit",
      domain,
      message,
    },
    version: SignTypedDataVersion.V4,
  });
}

function getERC2612PermitDomainDefinition(chainId: number, tokenAddress: string) {
  const domain = getEIP712DomainTypedData(chainId, tokenAddress);

  const types = {
    EIP712Domain: getEIP712DomainType(),
    Permit: getERC2612PermitType(),
  }

  return { domain, types }
}

function getERC2612PermitType() {
  return [
    { name: "owner", type: "address" },
    { name: "spender", type: "address" },
    { name: "value", type: "uint256" },
    { name: "nonce", type: "uint256" },
    { name: "deadline", type: "uint256" },
  ];
}

function getEIP712DomainType() {
  return [
    { name: "name", type: "string" },
    { name: "version", type: "string" },
    { name: "chainId", type: "uint256" },
    { name: "verifyingContract", type: "address" },
  ];
}

function getEIP712DomainTypedData(chainId: number, tokenAddress: string) {
  return {
    name: DomainName,
    version: DomainVersion,
    chainId,
    verifyingContract: tokenAddress
  };
}