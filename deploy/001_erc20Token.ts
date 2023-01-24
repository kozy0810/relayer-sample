import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { DeployFunction } from 'hardhat-deploy/types'

const func: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts, ethers } = hre;
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  await deploy('ERC20Token', {
    from: deployer,
    args: [],
    log: true,
    deterministicDeployment: true,
  });
}

export default func;
func.id = 'deploy_erc20_token';
func.tags = ['ERC20Token'];