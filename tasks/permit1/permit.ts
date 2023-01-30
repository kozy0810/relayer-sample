import { task } from "hardhat/config";
import { Wallet } from "@ethersproject/wallet";
import { splitSignature } from "ethers/lib/utils";
import { ERC2612PermitMessage, signERC2612PermitMessage } from "../../utils/utils"

task("erc20Permit", "Excute permit myself")
  // .addParam("contract", "Contract address")
  .addParam("owner", "Owner address")
  .addParam("spender", "Spender address")
  .addParam("value", "Sending amount")
  .setAction(async (taskArgs, hre) => {
    const wallet = new Wallet(process.env.PRIVATE_KEY as string);
    const token = await hre.ethers.getContract("ERC20Token");
    const nonce = await token.nonces(taskArgs.owner);
    const value = hre.ethers.utils.parseUnits(taskArgs.value, 18);

    const date = new Date();
    const deadline = Math.floor(date.getTime() + (1000 * 60 * 60 * 1) / 1000);

    const message: ERC2612PermitMessage = {
      owner: taskArgs.owner,
      spender: taskArgs.spender,
      value: value.toString(),
      nonce: nonce.toString(),
      deadline: deadline,
    };

    const chainId = await hre.getChainId();
    const sig = signERC2612PermitMessage(wallet.privateKey, token.address, Number(chainId), message);

    const { v, r, s } = splitSignature(sig);

    const resp = await token.permit(wallet.address, taskArgs.spender, value.toString(), deadline, v, r, s);
    console.log(resp);
  });