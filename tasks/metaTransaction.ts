import { task } from "hardhat/config";
import { Wallet } from "@ethersproject/wallet";
import { splitSignature } from "ethers/lib/utils";
import axios from "axios";

import { ERC2612PermitMessage, signERC2612PermitMessage, getRelayerURL } from "../utils/utils";

task("erc20Permit-via-relayer", "Excute permit via relayer")
  // .addParam("contract", "Contract address")
  .addParam("owner", "Owner address")
  .addParam("relayer", "Relayer address")
  .addParam("value", "Sending amount")
  .setAction(async (taskArgs, hre) => {
    const wallet = new Wallet(process.env.PRIVATE_KEY as string);
    const token = await hre.ethers.getContract("ERC20Token");
    console.log(token.address)
    const nonce = await token.nonces(taskArgs.owner);
    const value = hre.ethers.utils.parseUnits("1", 18)

    const date = new Date();
    const deadline = Math.floor(date.getTime() + (1000 * 60 * 60 * 1) / 1000);

    const message: ERC2612PermitMessage = {
      owner: taskArgs.owner,
      spender: taskArgs.relayer,
      value: value.toString(),
      nonce: nonce.toString(),
      deadline: deadline,
    };

    const chainId = await hre.getChainId();
    const sig = signERC2612PermitMessage(wallet.privateKey, token.address, Number(chainId), message);

    const { v, r, s } = splitSignature(sig);

    const relayerURL = getRelayerURL();
    console.log(relayerURL)

    const requestParams = {
      "owner": taskArgs.owner,
      "speder": taskArgs.spender,
      "value": Number(value),
      "deadline": deadline,
      "signature": sig,
      "v": v,
      "r": r,
      "s": s
    }

    try {
      await axios.post(relayerURL+"/permit", requestParams).then(resp => {
        if (resp.status !== 200) {
          throw new Error(resp.data)
        }
        console.log(resp.data);
      })
    } catch(e) {
      console.error(e);
    }
  });