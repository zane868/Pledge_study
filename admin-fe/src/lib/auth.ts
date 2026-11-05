/**
 * 后台登录态管理
 *
 * 账号密码登录走后端 /pledge/api/v1/user/login（JWT + Redis 会话），
 * token 存 localStorage。登录后可浏览全部页面；
 * 涉及链上写操作时再要求连接管理员钱包（链上 validCall 仍会校验地址）。
 */

const TOKEN_KEY = 'pledge_admin_token';
const USER_KEY = 'pledge_admin_user';
export const API_BASE = '/pledge/api/v1';

export interface LoginResult {
  ok: boolean;
  message?: string;
}

export function getToken(): string | null {
  return localStorage.getItem(TOKEN_KEY);
}

export function getUsername(): string | null {
  return localStorage.getItem(USER_KEY);
}

export function isLoggedIn(): boolean {
  return Boolean(getToken());
}

export function clearSession() {
  localStorage.removeItem(TOKEN_KEY);
  localStorage.removeItem(USER_KEY);
}

export async function login(name: string, password: string): Promise<LoginResult> {
  try {
    const res = await fetch(`${API_BASE}/user/login`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ name, password }),
    });
    const body = await res.json().catch(() => ({}));
    if (body?.code !== 0) {
      return { ok: false, message: body?.message ?? `登录失败（HTTP ${res.status}）` };
    }
    const token: string | undefined = body?.data?.token_id ?? body?.data?.token;
    if (!token) return { ok: false, message: '登录响应缺少 token' };
    localStorage.setItem(TOKEN_KEY, token);
    localStorage.setItem(USER_KEY, name);
    return { ok: true };
  } catch (err) {
    return { ok: false, message: `网络错误：${(err as Error).message}` };
  }
}

export async function logout(): Promise<void> {
  const token = getToken();
  if (token) {
    await fetch(`${API_BASE}/user/logout`, {
      method: 'POST',
      headers: { authCode: token },
    }).catch(() => {});
  }
  clearSession();
}