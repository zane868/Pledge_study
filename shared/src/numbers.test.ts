/**
 * 取整一致性测试：把链上测试 pledgev2/test/PledgeRounding.t.sol 里用过的
 * 同一组数字，喂给前端的镜像函数，断言两侧结果完全一致。
 *
 * 为什么要锁死：前端预览「预计可领 / 预计退款 / 预计赎回」用的是自己复刻的
 * 乘除顺序。只要有一处把两步除法合并成一步、或把向下取整改成四舍五入，
 * 用户看到的金额就会和链上实际到账对不上 —— 这属于会让用户误判收益的缺陷。
 *
 * 对应的链上基准测试：pledgev2/test/PledgeRounding.t.sol
 * 两边的向量必须同步修改。
 */
import { describe, expect, it } from 'vitest';
import {
  CAL_DECIMAL,
  BASE_DECIMAL,
  lendScale,
  mulDiv,
  expectedSp,
  expectedJp,
  expectedLoan,
  refundAmount,
  previewSpRedeem,
  previewJpRedeem,
  spDenominator,
  jpDenominator,
  formatRate,
  mortgageToLtv,
} from './numbers';

// ── 与线上部署一致的参数 ──────────────────────────────────────────────
const USDT = 6; // 6 位小数 USDT
const SCALE = lendScale(USDT); // 1e12
const MORTGAGE_RATE = 200_000_000n; // 200%
const ETH = 10n ** 18n;
const USDT_UNIT = 10n ** 6n;

describe('取整一致性：与 pledgev2/test/PledgeRounding.t.sol 对齐', () => {
  it('lendScale 由出借资产精度推导，而非前端写死的常量', () => {
    expect(SCALE).toBe(10n ** 12n);
    expect(lendScale(18)).toBe(1n);
    expect(lendScale(6)).toBe(10n ** 12n);
  });

  it('claimLend：SP 铸造量与链上一致', () => {
    // 出借 1500 USDT / 抵押 1 ETH → 匹配 1000 USDT
    const settleAmountLend = 1000n * USDT_UNIT;
    const lendSupply = 1500n * USDT_UNIT;
    const stakeAmount = 1500n * USDT_UNIT;

    // 链上：totalSpAmount = settleAmountLend * lendScale
    //       spAmount = mulDiv(totalSpAmount, stakeAmount, lendSupply)
    const totalSpAmount = settleAmountLend * SCALE;
    expect(totalSpAmount).toBe(1000n * ETH);

    expect(expectedSp(settleAmountLend, SCALE, stakeAmount, lendSupply)).toBe(totalSpAmount);
  });

  it('claimLend：向下取整，多人分摊时总额小于总 SP（差额留在合约）', () => {
    const settleAmountLend = 1000n * USDT_UNIT;
    // 三个出借人各出 100.000001 / 500.000001 / 900.000001 USDT
    const stakes = [100_000_001n, 500_000_001n, 900_000_001n];
    const lendSupply = 1500_000_003n;

    const totalSpAmount = settleAmountLend * SCALE;
    const sum = stakes.reduce((acc, s) => acc + expectedSp(settleAmountLend, SCALE, s, lendSupply), 0n);

    expect(sum).toBeLessThan(totalSpAmount);
    // 差额恰好 1 wei（18 位）
    expect(totalSpAmount - sum).toBe(1n);
  });

  it('claimBorrow：JP 数量与借款额分别取整，不可互相推导', () => {
    const settleAmountLend = 1000n * USDT_UNIT;
    const borrowSupply = 1n * ETH;

    // 链上：totalJpAmount = mulDiv(settleAmountLend * lendScale, martgageRate, baseDecimal)
    const totalJpAmount = mulDiv(settleAmountLend * SCALE, MORTGAGE_RATE, BASE_DECIMAL);
    expect(totalJpAmount).toBe(2000n * ETH);
    expect(jpDenominator(settleAmountLend, SCALE, MORTGAGE_RATE)).toBe(totalJpAmount);
    expect(spDenominator(settleAmountLend, SCALE)).toBe(1000n * ETH);

    expect(expectedJp(settleAmountLend, SCALE, MORTGAGE_RATE, 1n * ETH, borrowSupply)).toBe(
      2000n * ETH,
    );
    // 借款是 6 位 USDT 原始单位
    expect(expectedLoan(settleAmountLend, 1n * ETH, borrowSupply)).toBe(1000n * USDT_UNIT);
  });

  it('claimBorrow：多借款人各自向下取整', () => {
    const settleAmountLend = 1000n * USDT_UNIT;
    const borrowSupply = 15n * 10n ** 17n; // 1.5 ETH
    const totalJpAmount = mulDiv(settleAmountLend * SCALE, MORTGAGE_RATE, BASE_DECIMAL);

    const jpA = expectedJp(settleAmountLend, SCALE, MORTGAGE_RATE, 1n * ETH, borrowSupply);
    const jpB = expectedJp(settleAmountLend, SCALE, MORTGAGE_RATE, 5n * 10n ** 17n, borrowSupply);
    const loanA = expectedLoan(settleAmountLend, 1n * ETH, borrowSupply);
    const loanB = expectedLoan(settleAmountLend, 5n * 10n ** 17n, borrowSupply);

    expect(jpA + jpB).toBeLessThanOrEqual(totalJpAmount);
    expect(loanA + loanB).toBeLessThanOrEqual(settleAmountLend);
  });

  it('refundLend：未成交部分按本金占比退回', () => {
    // 出借 1500 / 匹配 1000 → 未成交 500
    const settleAmountLend = 1000n * USDT_UNIT;
    const lendSupply = 1500n * USDT_UNIT;
    const stakeAmount = 1500n * USDT_UNIT;

    expect(refundAmount(stakeAmount, lendSupply, settleAmountLend)).toBe(500n * USDT_UNIT);
  });

  it('refundLend：实现必须按链上的两步取整，而不是合并成一次除法', () => {
    // 链上：userShare = floor(stake * calDecimal / lendSupply)
    //       refund    = floor((lendSupply - settle) * userShare / calDecimal)
    //
    // 用 18 位口径的数值，两步与合并取整的结果相差 1 —— 合并会让用户
    // 看到比实际可领金额多 1 个最小单位的退款。
    const settleAmountLend = 1000n * ETH;
    const lendSupply = 1500n * ETH;
    const stakeAmount = 333_000_000_000_000_000_003n;

    const twoStep = refundAmount(stakeAmount, lendSupply, settleAmountLend);
    const merged = ((lendSupply - settleAmountLend) * stakeAmount) / lendSupply;

    expect(twoStep).not.toBe(merged);
    expect(twoStep).toBe(
      ((lendSupply - settleAmountLend) * ((stakeAmount * CAL_DECIMAL) / lendSupply)) / CAL_DECIMAL,
    );
    // 链上取整向下，前端不能比链上多给
    expect(twoStep).toBeLessThan(merged);
  });

  it('refundBorrow：分母是 borrowSupply，未成交的 ETH 按同样两步退回', () => {
    const settleAmountBorrow = 6n * 10n ** 17n;
    const borrowSupply = 10n ** 18n + 1n;
    const stakeAmount = 10n ** 17n;

    const twoStep = refundAmount(stakeAmount, borrowSupply, settleAmountBorrow);
    const merged = ((borrowSupply - settleAmountBorrow) * stakeAmount) / borrowSupply;

    expect(twoStep).not.toBe(merged);
    expect(twoStep).toBe(
      ((borrowSupply - settleAmountBorrow) * ((stakeAmount * CAL_DECIMAL) / borrowSupply)) / CAL_DECIMAL,
    );
  });

  it('refundLend：多人退款之和不超过未成交总额', () => {
    const settleAmountLend = 1000n * USDT_UNIT;
    const lendSupply = 2200n * USDT_UNIT;

    const refundA = refundAmount(1500n * USDT_UNIT, lendSupply, settleAmountLend);
    const refundB = refundAmount(700n * USDT_UNIT, lendSupply, settleAmountLend);

    expect(refundA + refundB).toBeLessThanOrEqual(lendSupply - settleAmountLend);
  });

  it('refundBorrow：分母是 borrowSupply，不是凭证总供应量', () => {
    // 匹配 0.6 ETH / 抵押 1 ETH → 未成交 0.4 ETH
    const settleAmountBorrow = 6n * 10n ** 17n;
    const borrowSupply = 1n * ETH;

    expect(refundAmount(1n * ETH, borrowSupply, settleAmountBorrow)).toBe(4n * 10n ** 17n);
  });

  it('赎回：SP 与 JP 分母不同，不可混用', () => {
    const settleAmountLend = 1000n * USDT_UNIT;
    const totalSpAmount = settleAmountLend * SCALE;
    const totalJpAmount = mulDiv(settleAmountLend * SCALE, MORTGAGE_RATE, BASE_DECIMAL);

    expect(totalSpAmount).not.toBe(totalJpAmount);
    expect(totalJpAmount).toBe(totalSpAmount * 2n);
  });

  it('赎回：SP 侧零兑付必须回滚，前端据此禁用按钮', () => {
    const settleAmountLend = 1000n * USDT_UNIT;
    const totalSpAmount = settleAmountLend * SCALE;

    // 链上 finishAmountLend 假设为 1200 USDT
    const finishAmountLend = 1200n * USDT_UNIT;
    const { amount, willRevert } = previewSpRedeem(2, 1n, totalSpAmount, finishAmountLend, 0n);

    expect(amount).toBe(0n);
    expect(willRevert).toBe(true);
  });

  it('赎回：JP 侧零兑付不报错 —— DEC-14 需向前端披露', () => {
    const settleAmountLend = 1000n * USDT_UNIT;
    const totalJpAmount = mulDiv(settleAmountLend * SCALE, MORTGAGE_RATE, BASE_DECIMAL);
    const finishAmountBorrow = 1n * ETH;

    const { amount, willRevert } = previewJpRedeem(2, 1n, totalJpAmount, finishAmountBorrow, 0n);

    expect(amount).toBe(0n);
    // JP 侧没有 require(> 0)，交易会成功执行
    expect(willRevert).toBe(false);
  });

  it('赎回：正常路径与链上一致', () => {
    const settleAmountLend = 1000n * USDT_UNIT;
    const totalSpAmount = settleAmountLend * SCALE;
    const totalJpAmount = mulDiv(settleAmountLend * SCALE, MORTGAGE_RATE, BASE_DECIMAL);
    const finishAmountLend = 1200n * USDT_UNIT;
    const finishAmountBorrow = 1n * ETH;

    expect(previewSpRedeem(2, 1000n * ETH, totalSpAmount, finishAmountLend, 0n)).toEqual({
      amount: finishAmountLend,
      willRevert: false,
    });
    expect(previewJpRedeem(2, 2000n * ETH, totalJpAmount, finishAmountBorrow, 0n)).toEqual({
      amount: finishAmountBorrow,
      willRevert: false,
    });
  });

  it('退款金额不会为负：匹配额不超过出借总额', () => {
    expect(refundAmount(1000n * USDT_UNIT, 1000n * USDT_UNIT, 1000n * USDT_UNIT)).toBe(0n);
  });
});

describe('比例格式化：1e8 基数 → 百分比', () => {
  it('formatRate 把 1e8 基数换算成百分数', () => {
    // 1e8 = 100%，不是 1%
    expect(formatRate(20_000_000n)).toBe('20%');
    expect(formatRate(200_000_000n)).toBe('200%');
    expect(formatRate(100_000_000n)).toBe('100%');
    expect(formatRate(12_500_000n)).toBe('12.5%');
    expect(formatRate(0n)).toBe('0%');
    // 极小的非零比例不会被抹成 0
    expect(formatRate(1n, 2)).toBe('0%');
    expect(formatRate(1n, 6)).toBe('0.000001%');
  });

  it('mortgageToLtv 是抵押率的倒数', () => {
    expect(mortgageToLtv(200_000_000n)).toBe('50%');
    expect(mortgageToLtv(150_000_000n)).toBe('66.66%');
    expect(mortgageToLtv(400_000_000n)).toBe('25%');
    expect(mortgageToLtv(0n)).toBe('—');
  });
});

describe('formatRate 的 1e8 基数口径', () => {
  it('1e8 表示 100%，不是 1%', () => {
    expect(formatRate(20_000_000n)).toBe('20%');
    expect(formatRate(200_000_000n)).toBe('200%');
  });
});
