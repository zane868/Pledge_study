/**
 * 后端 API 客户端
 *
 * 依据 Pledge_PRD.md 13.4：响应可能 HTTP 200 但业务 code != 0，
 * 客户端必须同时检查 code。同步滞后时市场接口返回 503。
 */

const API_BASE = (globalThis as any).__PLEDGE_API_BASE__ ?? '/pledge/api/v1';

export interface ApiEnvelope<T> {
  code: number;
  message: string;
  data: T;
}

export class ApiError extends Error {
  constructor(
    message: string,
    readonly status: number,
    readonly code?: number,
  ) {
    super(message);
    this.name = 'ApiError';
  }
}

async function request<T>(path: string, init?: RequestInit): Promise<T> {
  let res: Response;
  try {
    res = await fetch(`${API_BASE}${path}`, {
      ...init,
      headers: { 'Content-Type': 'application/json', ...(init?.headers ?? {}) },
    });
  } catch (err) {
    throw new ApiError(`网络请求失败：${(err as Error).message}`, 0);
  }

  // 503 = 同步未就绪或快照过期，属于"数据不可用"，不是空市场
  if (res.status === 503) {
    const body = await res.json().catch(() => ({}));
    throw new ApiError(body?.message ?? '数据同步中，请稍后重试', 503, body?.code);
  }
  if (!res.ok) {
    throw new ApiError(`请求失败（HTTP ${res.status}）`, res.status);
  }

  const body = (await res.json()) as ApiEnvelope<T>;
  if (body && typeof body.code === 'number' && body.code !== 0) {
    throw new ApiError(body.message || `业务错误 ${body.code}`, res.status, body.code);
  }
  return body as T;
}

// ─── 类型（与 pledge-backend/api/models 对应）─────────────────────────

export interface LendTokenInfo {
  lendFee: string;
  tokenLogo: string;
  tokenName: string;
  tokenPrice: string;
}

export interface BorrowTokenInfo {
  borrowFee: string;
  tokenLogo: string;
  tokenName: string;
  tokenPrice: string;
}

export interface PoolBaseInfo {
  poolContract: string;
  syncBlockNumber: string;
  syncBlockHash: string;
  pool_id: number;
  autoLiquidateThreshold: string;
  borrowSupply: string;
  borrowToken: string;
  borrowTokenInfo: BorrowTokenInfo;
  endTime: string;
  interestRate: string;
  jpCoin: string;
  lendSupply: string;
  lendToken: string;
  lendTokenInfo: LendTokenInfo;
  martgageRate: string;
  maxSupply: string;
  settleTime: string;
  spCoin: string;
  state: string;
}

export interface PoolBaseInfoRes {
  poolIndex: number;
  displayPoolId: number;
  index: number;
  pool_data: PoolBaseInfo;
}

export interface PoolData {
  poolContract: string;
  syncBlockNumber: string;
  syncBlockHash: string;
  pool_id: number;
  chain_id: string;
  finish_amount_borrow: string;
  finish_amount_lend: string;
  liquidation_amoun_borrow: string;
  liquidation_amoun_lend: string;
  settle_amount_borrow: string;
  settle_amount_lend: string;
  created_at: string;
  updated_at: string;
}

export interface SyncStatus {
  chainId: string;
  poolContract: string;
  blockNumber: string;
  blockHash: string;
  poolCount: number;
  complete: boolean;
  lastSuccess: string;
  lastError?: string;
}

export interface SyncStatusEnvelope {
  code: number;
  message: string;
  data: SyncStatus;
  stale: boolean;
}

export interface DeploymentInfo {
  chainId: number;
  pool: string;
  oracle: string;
  multisig: string;
  stable: string;
  sp: string;
  jp: string;
  operator: string;
  admins: string[];
  threshold: number;
  stableDecimals: number;
  deploymentStartBlock: number;
}

export interface TokenEntry {
  name: string;
  decimals: number;
  symbol: string;
  address: string;
  chainId: number;
  logoURI?: string;
}

export interface TokenListResponse {
  name: string;
  logoURI: string;
  tokens: TokenEntry[];
  timestamp: string;
}

// ─── 公开接口 ─────────────────────────────────────────────────────────

export function fetchSyncStatus(): Promise<SyncStatusEnvelope> {
  return request<SyncStatusEnvelope>('/sync/status');
}

export function fetchDeployment(): Promise<ApiEnvelope<DeploymentInfo>> {
  return request<ApiEnvelope<DeploymentInfo>>('/deployment');
}

export function fetchPoolBaseInfo(chainId: number): Promise<ApiEnvelope<PoolBaseInfoRes[]>> {
  return request<ApiEnvelope<PoolBaseInfoRes[]>>(`/poolBaseInfo?chainId=${chainId}`);
}

export function fetchPoolDataInfo(chainId: number): Promise<ApiEnvelope<{ pool_data: PoolData }[]>> {
  return request<ApiEnvelope<{ pool_data: PoolData }[]>>(`/poolDataInfo?chainId=${chainId}`);
}

export function fetchTokenList(chainId: number): Promise<TokenListResponse> {
  return fetch(`${API_BASE}/token?chainId=${chainId}`).then((r) => {
    if (!r.ok) throw new ApiError(`代币列表读取失败（HTTP ${r.status}）`, r.status);
    return r.json();
  });
}

// ─── 后台接口（需要 authCode）─────────────────────────────────────────

let authToken: string | null = null;

export function setAuthToken(token: string | null) {
  authToken = token;
}

export function getAuthToken(): string | null {
  return authToken;
}

function authHeaders(): Record<string, string> {
  return authToken ? { authCode: authToken } : {};
}

export async function login(name: string, password: string): Promise<string> {
  const res = await fetch(`${API_BASE}/user/login`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ name, password }),
  });
  const body = await res.json().catch(() => ({}));
  if (body?.code !== 0) {
    throw new ApiError(body?.message ?? '登录失败', res.status, body?.code);
  }
  const token = body?.data?.token_id ?? body?.data?.token;
  if (!token) throw new ApiError('登录响应缺少 token', res.status);
  authToken = token;
  return token as string;
}

export async function logout(): Promise<void> {
  if (!authToken) return;
  await fetch(`${API_BASE}/user/logout`, { method: 'POST', headers: authHeaders() }).catch(() => {});
  authToken = null;
}

export interface SearchParams {
  chainID: number;
  lend_token_symbol?: string;
  state?: string;
  page?: number;
  pageSize?: number;
}

export function searchPools(params: SearchParams): Promise<ApiEnvelope<{ count: number; rows: unknown[] }>> {
  return request<ApiEnvelope<{ count: number; rows: unknown[] }>>('/pool/search', {
    method: 'POST',
    headers: authHeaders(),
    body: JSON.stringify(params),
  });
}