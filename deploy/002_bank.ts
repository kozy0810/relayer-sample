import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { DeployFunction } from 'hardhat-deploy/types'

const func: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts, ethers } = hre;
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  const token = await ethers.getContract('ERC20Token')

  await deploy('Bank', {
    from: deployer,
    args: [token.address],
    log: true,
    deterministicDeployment: true,
  });
}

export default func;
func.id = 'deploy_bank';
func.tags = ['Bank'];