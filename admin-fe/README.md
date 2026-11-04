# 运营后台

后台管理界面，用于管理 Pledge 协议的池子和多签提案。

## 功能

- 池列表与创建
- 多签提案管理
- 系统概览

## 关键操作接口

### 完成募集（settle）

**调用时机**：池子状态为 MATCH (0)，募集期结束，需要转入执行期。

**合约接口**：`pool.settle(uint256 pid)` 或 `pool.forceSettle(uint256 pid)`（强制结算，跳过时间校验）

**多签流程**：
1. 创建授权申请：`multisig.createApplication(pool_address)`，生成 msghash = `keccak256(申请人地址, pool地址)`
2. 其他管理员签名：`multisig.signApplication(msghash)`
3. 达到阈值（2/3）后，申请人调用 `pool.settle(pid)` 或 `pool.forceSettle(pid)`

**状态转换**：MATCH → EXECUTION，锁定资金，记录 settleAmountLend / settleAmountBorrow，开始计息。

### 结束池子（finish）

**调用时机**：池子状态为 EXECUTION (1)，借贷期到期，需要结算利息并清算。

**合约接口**：`pool.finish(uint256 pid)` 或 `pool.forceFinish(uint256 pid)`（强制结算，跳过时间校验）

**多签流程**：同 settle（需创建授权申请并获得签名）

**状态转换**：EXECUTION → FINISH，计算利息，执行 DEX 交换，分配资金给出借人和借款人。

**强制接口说明**：
- `forceSettle` / `forceFinish` / `forceLiquidate` 为运营紧急操作接口，跳过时间校验，但仍需多签授权。
- 计息、交换、状态机逻辑与常规接口完全一致，仅时间门槛不同。

## 运行

```bash
npm install
npm run dev
```

访问 http://localhost:5175
