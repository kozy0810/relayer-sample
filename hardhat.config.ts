import "@nomiclabs/hardhat-waffle";
import "@typechain/hardhat";
import "@nomiclabs/hardhat-ethers";
import "hardhat-gas-reporter";
import "hardhat-deploy";
import "@nomiclabs/hardhat-etherscan";
import "hardhat-contract-sizer";
import { HardhatUserConfig } from "hardhat/config";
import "hardhat-storage-layout";
import dotenv from "dotenv";
dotenv.config();

import "./tasks";

// const PRIVATE_KEY = process.env.PRIVATE_KEY || null;

const config: HardhatUserConfig = {
  solidity: {
    version: "0.8.17",
    settings: {
      optimizer: {
        enabled: true,
        runs: 10000,
      },
    },
  },
  defaultNetwork: "hardhat",
  networks: {
    hardhat: {
      // allowUnlimitedContractSize: true,
      // forking: {
      //   url: process.env.ASTAR_RPC_URL as string
      // }
    },
  },
  paths: {
    sources: "src",
  },
  namedAccounts: {
    deployer: 0
  },
  typechain: {
    outDir: "typechain",
    target: "ethers-v5",
  },
  mocha: {
    timeout: 0,
  },
  etherscan: {
    apiKey: process.env.ETHERSCAN_API_KEY,
  },
  gasReporter: {
    currency: "USD",
    gasPrice: 100,
    enabled: process.env.REPORT_GAS ? true : false,
    coinmarketcap: process.env.COINMARKETCAP_API_KEY,
    maxMethodDiff: 10,
  },
  contractSizer: {
    alphaSort: true,
    disambiguatePaths: false,
    runOnCompile: true,
    strict: true,
  },
};

export default config;