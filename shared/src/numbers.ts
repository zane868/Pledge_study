/**
 * 数值层：链上算法的忠实镜像
 *
 * 依据 Pledge_PRD.md 17.2.3（已确认的精度规则）：
 *   - USDT 6 位，ETH 18 位，SP/JP 18 位
 *   - 利率、抵押率、费用比例基数为 1e8（baseDecimal）
 *   - 估值先统一到 18 位计算口径（calDecimal）
 *   - 可借额度向下取整，所需匹配抵押物向上取整
 *   - SP 固定分母 = settleAmountLend(USDT raw) × 10^12
 *   - JP 固定分母 = floor(SP 分母 × martgageRate / 1e8)
 *   - 赎回使用 Math.mulDiv（向下取整），不使用浮点
 *
 * 这些函数必须与 PledgePool.sol 的计算逐位一致，禁止用 Number 参与金额运算。
 */

export const CAL_DECIMAL = 10n ** 18n; // calDecimal
export const BASE_DECIMAL = 10n ** 8n; // baseDecimal，比例/利率/费用的基数
export const YEAR_SECONDS = 365n * 24n * 60n * 60n;

/** 稳定币精度 → 归一化到 18 位的缩放因子（lendScale） */
export function lendScale(stableDecimals: number): bigint {
  if (stableDecimals < 0 || stableDecimals > 18) {
    throw new Error(`unexpected stable decimals: ${stableDecimals}`);
  }
  return 10n ** BigInt(18 - stableDecimals);
}

/**
 * 0.8.24 内建的 Math.mulDiv —— 与 OpenZeppelin 实现一致：
 * floor(x * y / denominator)，中间乘积用全精度计算，不回退到浮点。
 */
export function mulDiv(x: bigint, y: bigint, denominator: bigint): bigint {
  if (denominator === 0n) throw new Error('Math: division by zero');
  return (x * y) / denominator;
}

/** Math.mulDiv 向上取整变体 */
export function mulDivUp(x: bigint, y: bigint, denominator: bigint): bigint {
  if (denominator === 0n) throw new Error('Math: division by zero');
  const p = x * y;
  const q = p / denominator;
  return p % denominator === 0n ? q : q + 1n;
}

/** 按资产精度格式化，返回原始整数（不做四舍五入） */
export function formatUnits(value: bigint, decimals: number): string {
  if (decimals === 0) return value.toString();
  const negative = value < 0n;
  const abs = negative ? -value : value;
  const base = 10n ** BigInt(decimals);
  const whole = abs / base;
  const frac = abs % base;
  if (frac === 0n) return (negative ? '-' : '') + whole.toString();
  const fracStr = frac.toString().padStart(decimals, '0').replace(/0+$/, '');
  return (negative ? '-' : '') + `${whole}.${fracStr}`;
}

/** 解析用户输入为原始整数，拒绝非法与超精度输入（不静默截断） */
export function parseUnits(input: string, decimals: number): bigint {
  const trimmed = input.trim();
  if (trimmed === '') throw new Error('空输入');
  if (!/^\d*\.?\d*$/.test(trimmed)) throw new Error('只允许数字和小数点');
  const [whole, frac = ''] = trimmed.split('.');
  if (frac.length > decimals) {
    throw new Error(`最多 ${decimals} 位小数`);
  }
  const padded = frac.padEnd(decimals, '0');
  return BigInt((whole || '0') + padded);
}

/** 展示用：截断到指定小数位，仅用于 UI 显示，不参与计算 */
export function displayUnits(value: bigint, decimals: number, maxFrac = 6): string {
  const raw = formatUnits(value, decimals);
  const [w, f] = raw.split('.');
  if (!f) return w;
  const cut = f.slice(0, maxFrac).replace(/0+$/, '');
  return cut ? `${w}.${cut}` : w;
}

/**
 * 小额非零余额：PRD 11.5.3 要求显示"小于最小显示单位"而不是显示为零。
 */
export function displayOrLessThan(value: bigint, decimals: number, maxFrac = 6): string {
  if (value === 0n) return formatUnits(0n, decimals);
  const shown = displayUnits(value, decimals, maxFrac);
  if (shown === '0') return '< 最小显示单位';
  return shown;
}

// ─── 匹配与估值（对应 settle / checkoutLiquidate）─────────────────────

/**
 * 抵押能力：给定抵押量与相对价格，按抵押率能支撑的借款额度（USDT raw）。
 * 对应 PledgePool.settle：
 *   valueWad  = borrowSupply × relativePrice / calDecimal
 *   capacity  = valueWad × baseDecimal / martgageRate / lendScale
 *
 * @param borrowSupply 抵押量（ETH 18 位原始量）
 * @param priceLend    借款资产价格（1e8）
 * @param priceBorrow  抵押资产价格（1e8）
 * @param martgageRate 抵押率（1e8，200% = 2e8）
 * @param scale        lendScale
 */
export function collateralCapacity(
  borrowSupply: bigint,
  priceLend: bigint,
  priceBorrow: bigint,
  martgageRate: bigint,
  scale: bigint,
): bigint {
  if (priceLend === 0n || martgageRate === 0n) return 0n;
  const relativePrice = mulDiv(priceBorrow, CAL_DECIMAL, priceLend);
  if (relativePrice === 0n) return 0n;
  const valueWad = mulDiv(borrowSupply, relativePrice, CAL_DECIMAL);
  return mulDiv(valueWad, BASE_DECIMAL, martgageRate) / scale;
}

/**
 * 反算所需抵押物（ETH raw），向上取整。
 * 对应 settle 的 else 分支：
 *   requiredValue      = lendSupply × lendScale × martgageRate / baseDecimal（Up）
 *   settleAmountBorrow = requiredValue × calDecimal / relativePrice（Up）
 */
export function requiredCollateral(
  lendSupply: bigint,
  priceLend: bigint,
  priceBorrow: bigint,
  martgageRate: bigint,
  scale: bigint,
): bigint {
  if (priceBorrow === 0n) return 0n;
  const relativePrice = mulDiv(priceBorrow, CAL_DECIMAL, priceLend);
  if (relativePrice === 0n) return 0n;
  const requiredValue = mulDivUp(lendSupply * scale, martgageRate, BASE_DECIMAL);
  return mulDivUp(requiredValue, CAL_DECIMAL, relativePrice);
}

/**
 * 预估匹配结果（前端展示用，不写入任何链上状态）。
 * 与 settle 分支逻辑一致：取「出借量」与「抵押能力」的较小值。
 */
export interface MatchEstimate {
  settleAmountLend: bigint;
  settleAmountBorrow: bigint;
  /** 出借侧超额，将退回 */
  lendRefundBase: bigint;
  /** 抵押侧超额，将退回 */
  borrowRefundBase: bigint;
  /** 无法形成有效匹配 */
  zero: boolean;
}

export function estimateMatch(params: {
  lendSupply: bigint;
  borrowSupply: bigint;
  priceLend: bigint;
  priceBorrow: bigint;
  martgageRate: bigint;
  scale: bigint;
}): MatchEstimate {
  const { lendSupply, borrowSupply, priceLend, priceBorrow, martgageRate, scale } = params;
  if (lendSupply === 0n || borrowSupply === 0n) {
    return {
      settleAmountLend: lendSupply,
      settleAmountBorrow: borrowSupply,
      lendRefundBase: 0n,
      borrowRefundBase: 0n,
      zero: true,
    };
  }
  const capacity = collateralCapacity(borrowSupply, priceLend, priceBorrow, martgageRate, scale);
  if (capacity === 0n) {
    return {
      settleAmountLend: 0n,
      settleAmountBorrow: 0n,
      lendRefundBase: lendSupply,
      borrowRefundBase: borrowSupply,
      zero: true,
    };
  }
  if (lendSupply > capacity) {
    return {
      settleAmountLend: capacity,
      settleAmountBorrow: borrowSupply,
      lendRefundBase: lendSupply - capacity,
      borrowRefundBase: 0n,
      zero: false,
    };
  }
  const needBorrow = requiredCollateral(lendSupply, priceLend, priceBorrow, martgageRate, scale);
  if (needBorrow > borrowSupply) {
    // 链上此处 require 会回滚；前端只做提示，不当作可执行
    return {
      settleAmountLend: lendSupply,
      settleAmountBorrow: needBorrow,
      lendRefundBase: 0n,
      borrowRefundBase: 0n,
      zero: false,
    };
  }
  return {
    settleAmountLend: lendSupply,
    settleAmountBorrow: needBorrow,
    lendRefundBase: 0n,
    borrowRefundBase: borrowSupply - needBorrow,
    zero: false,
  };
}

// ─── 凭证固定分母（对应 claimLend / claimBorrow / withdrawLend / withdrawBorrow）

/** SP 固定分母 = settleAmountLend × lendScale */
export function spDenominator(settleAmountLend: bigint, scale: bigint): bigint {
  return settleAmountLend * scale;
}

/**
 * JP 固定分母 = floor(settleAmountLend × lendScale × martgageRate / baseDecimal)
 * 对应 claimBorrow: Math.mulDiv(data.settleAmountLend * lendScale, pool.martgageRate, baseDecimal)
 */
export function jpDenominator(settleAmountLend: bigint, scale: bigint, martgageRate: bigint): bigint {
  return mulDiv(settleAmountLend * scale, martgageRate, BASE_DECIMAL);
}

/** 用户应得 SP：mulDiv(总 SP, stakeAmount, lendSupply) */
export function expectedSp(
  settleAmountLend: bigint,
  scale: bigint,
  stakeAmount: bigint,
  lendSupply: bigint,
): bigint {
  if (lendSupply === 0n) return 0n;
  return mulDiv(settleAmountLend * scale, stakeAmount, lendSupply);
}

/** 用户应得 JP：mulDiv(总 JP, stakeAmount, borrowSupply) */
export function expectedJp(
  settleAmountLend: bigint,
  scale: bigint,
  martgageRate: bigint,
  stakeAmount: bigint,
  borrowSupply: bigint,
): bigint {
  if (borrowSupply === 0n) return 0n;
  const total = jpDenominator(settleAmountLend, scale, martgageRate);
  return mulDiv(total, stakeAmount, borrowSupply);
}

/** 借款人可领借款：mulDiv(settleAmountLend, stakeAmount, borrowSupply) */
export function expectedLoan(
  settleAmountLend: bigint,
  stakeAmount: bigint,
  borrowSupply: bigint,
): bigint {
  if (borrowSupply === 0n) return 0n;
  return mulDiv(settleAmountLend, stakeAmount, borrowSupply);
}

// ─── 退款（对应 refundLend / refundBorrow）───────────────────────────

/**
 * 未匹配退款份额。
 * 链上 refundLend：
 *   userShare    = floor(stakeAmount × calDecimal / lendSupply)
 *   refundAmount = floor((lendSupply - settleAmountLend) × userShare / calDecimal)
 * 两步都向下取整，此处按同样顺序计算，避免合并成一次除法产生偏差。
 *
 * 注意分母是 lendSupply（出借总额），不是凭证 totalSupply。
 */
export function refundAmount(
  stakeAmount: bigint,
  lendSupply: bigint,
  settleAmountLend: bigint,
): bigint {
  if (lendSupply === 0n) return 0n;
  const userShare = mulDiv(stakeAmount, CAL_DECIMAL, lendSupply);
  if (userShare === 0n) return 0n;
  return mulDiv(lendSupply - settleAmountLend, userShare, CAL_DECIMAL);
}

// ─── 赎回（对应 withdrawLend / withdrawBorrow）───────────────────────

/**
 * SP 赎回预览：FINISH 用 finishAmountLend，LIQUIDATION 用 liquidationAmounLend。
 * 链上：mulDiv(分配总额, spAmount, totalSpAmount)，且结果必须 > 0，否则回滚。
 */
export function previewSpRedeem(
  state: number,
  spAmount: bigint,
  totalSpAmount: bigint,
  finishAmountLend: bigint,
  liquidationAmountLend: bigint,
): { amount: bigint; willRevert: boolean } {
  if (totalSpAmount === 0n) return { amount: 0n, willRevert: true };
  const poolAmount = state === 2 ? finishAmountLend : liquidationAmountLend;
  const amount = mulDiv(poolAmount, spAmount, totalSpAmount);
  // 合约 require(redeemAmount > 0, "Zero USDT redemption")
  return { amount, willRevert: amount === 0n };
}

/**
 * JP 赎回预览，语义与 SP 对称。
 * 链上 withdrawBorrow 的分母是 jp 总量（= settleAmountLend × lendScale × martgageRate / baseDecimal），
 * 与 SP 赎回的分母（settleAmountLend × lendScale）不同，不要混用。
 * JP 侧合约没有 require(>0)，零兑付会成功执行 —— 属 DEC-14 需披露的行为。
 */
export function previewJpRedeem(
  state: number,
  jpAmount: bigint,
  totalJpAmount: bigint,
  finishAmountBorrow: bigint,
  liquidationAmountBorrow: bigint,
): { amount: bigint; willRevert: boolean } {
  if (totalJpAmount === 0n) return { amount: 0n, willRevert: false };
  const poolAmount = state === 2 ? finishAmountBorrow : liquidationAmountBorrow;
  const amount = mulDiv(poolAmount, jpAmount, totalJpAmount);
  // 合约在 JP 侧没有 require(>0)，零兑付会成功执行
  return { amount, willRevert: false };
}

// ─── 利息与费用（对应 finish / liquidate）───────────────────────────

/**
 * 期限利息（USDT raw）。
 * 链上：timeRatio = (endTime - settleTime) × baseDecimal / baseYear
 *       interest  = timeRatio × interestRate × settleAmountLend / 1e16
 */
export function termInterest(
  settleTime: bigint,
  endTime: bigint,
  interestRate: bigint,
  settleAmountLend: bigint,
): bigint {
  if (endTime <= settleTime) return 0n;
  const timeRatio = mulDiv(endTime - settleTime, BASE_DECIMAL, YEAR_SECONDS);
  return mulDiv(timeRatio, interestRate * settleAmountLend, 10n ** 16n);
}

/** 用户期限收益预估：按份额摊分利息 */
export function userInterestShare(
  interest: bigint,
  stakeAmount: bigint,
  settleAmountLend: bigint,
): bigint {
  if (settleAmountLend === 0n) return 0n;
  return mulDiv(interest, stakeAmount, settleAmountLend);
}

// ─── 利率 / 抵押率（1e8 基数）────────────────────────────────────────

/**
 * 把 1e8 基数的比例（利率 / 抵押率 / 阈值）格式化成百分比。
 * 1e8 = 100%。金额始终用整数运算，避免 Number 的精度损失。
 *
 *   20_000_000n  → "20%"
 *   200_000_000n → "200%"
 *   12_500_000n  → "12.5%"
 */
export function formatRate(value1e8: bigint, frac = 2): string {
  // 1e8 基数 = 100%，所以先乘 100 再按 1e8 拆整数和小数
  const scaled = value1e8 * 100n;
  const whole = scaled / BASE_DECIMAL;
  const rest = scaled % BASE_DECIMAL;
  if (rest === 0n) return `${whole}%`;
  const digits = (rest * 10n ** BigInt(frac)) / BASE_DECIMAL;
  const fracStr = digits.toString().padStart(frac, '0').replace(/0+$/, '');
  return fracStr ? `${whole}.${fracStr}%` : `${whole}%`;
}

/** 抵押率 200% 同口径 LTV 50% */
export function mortgageToLtv(value1e8: bigint): string {
  if (value1e8 === 0n) return '—';
  // 抵押率 200% → LTV 50%。抵押率是 1e8 基数（1e8 = 100%），
  // LTV = 100% / 抵押率比例，用 1e4 基数保留两位小数。
  const ratePct = (value1e8 * 100n) / BASE_DECIMAL; // 200
  if (ratePct === 0n) return '—';
  const ltvBp = (100n * 10_000n) / ratePct;
  const whole = ltvBp / 100n;
  const frac = ltvBp % 100n;
  if (frac === 0n) return `${whole}%`;
  return `${whole}.${frac.toString().padStart(2, '0').replace(/0+$/, '')}%`;
}

// ─── 时间 ─────────────────────────────────────────────────────────────

const TZ = 'Asia/Shanghai';

export function formatTimestamp(unix: bigint | number): string {
  const ms = Number(unix) * 1000;
  return new Intl.DateTimeFormat('zh-CN', {
    timeZone: TZ,
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
    hour12: false,
  }).format(new Date(ms));
}

export function formatTimestampWithTz(unix: bigint | number): string {
  return `${formatTimestamp(unix)} (UTC+8)`;
}

/** 剩余时间，仅作辅助展示；可执行性必须以链上时间为准 */
export function remainingLabel(target: bigint, chainNow: bigint): string {
  if (chainNow >= target) return '已到期';
  let secs = Number(target - chainNow);
  const d = Math.floor(secs / 86400);
  secs -= d * 86400;
  const h = Math.floor(secs / 3600);
  secs -= h * 3600;
  const m = Math.floor(secs / 60);
  if (d > 0) return `${d} 天 ${h} 小时`;
  if (h > 0) return `${h} 小时 ${m} 分`;
  return `${m} 分`;
}

/** 约定期限（天） */
export function termDays(settleTime: bigint, endTime: bigint): number {
  if (endTime <= settleTime) return 0;
  return Number((endTime - settleTime) / 86400n);
}