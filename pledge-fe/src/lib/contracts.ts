import PledgePoolAbi from '@shared-abi/PledgePool.json';
import DebtTokenAbi from '@shared-abi/DebtToken.json';
import MultiSignatureAbi from '@shared-abi/multiSignature.json';
import OracleAbi from '@shared-abi/BscPledgeOracle.json';
import ERC20Abi from '@shared-abi/ERC20.json';
import { DEPLOYMENT } from '@shared/config';

// ABI 与 pledgev2/abi 保持同一份来源，见 shared/abi。禁止在此手工改写。
export { PledgePoolAbi, DebtTokenAbi, MultiSignatureAbi, OracleAbi, ERC20Abi };

export const POOL_ADDRESS = DEPLOYMENT.pool;
export const ORACLE_ADDRESS = DEPLOYMENT.oracle;
export const MULTISIG_ADDRESS = DEPLOYMENT.multisig;
export const STABLE_ADDRESS = DEPLOYMENT.stable;
export const SP_ADDRESS = DEPLOYMENT.sp;
export const JP_ADDRESS = DEPLOYMENT.jp;
export const ROUTER_ADDRESS = DEPLOYMENT.router;
export const FACTORY_ADDRESS = DEPLOYMENT.factory;

/** 池状态枚举，与 PledgePool.PoolState 一致 */
export const POOL_STATE = {
  MATCH: 0,
  EXECUTION: 1,
  FINISH: 2,
  LIQUIDATION: 3,
  UNDONE: 4,
} as const;