import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { login } from '@/lib/auth';

// 后台只有一个管理员账号，用户名固定，登录时只输入密码。
const ADMIN_NAME = 'admin';

export function LoginPage() {
  const navigate = useNavigate();
  const [password, setPassword] = useState('');
  const [error, setError] = useState<string | null>(null);
  const [pending, setPending] = useState(false);

  const submit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);
    setPending(true);
    const result = await login(ADMIN_NAME, password);
    setPending(false);
    if (result.ok) {
      navigate('/', { replace: true });
    } else {
      setError(result.message ?? '登录失败');
    }
  };

  return (
    <div className="flex min-h-screen items-center justify-center bg-ink-900 p-4">
      <form onSubmit={submit} className="card w-full max-w-md space-y-5">
        <div>
          <h1 className="text-2xl font-semibold text-white">Pledge 运营后台</h1>
          <p className="mt-2 text-sm text-slate-400">输入管理员密码登录</p>
        </div>

        <div>
          <label className="label">密码</label>
          <input
            className="field"
            type="password"
            autoComplete="current-password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
            autoFocus
          />
        </div>

        {error && <p className="text-sm text-negative">{error}</p>}

        <button type="submit" className="btn-primary w-full" disabled={pending}>
          {pending ? '登录中…' : '登录'}
        </button>

        <p className="text-xs text-slate-500">
          登录后可浏览全部页面；执行创建池、结算、签名等链上操作时才需要连接管理员钱包。
        </p>
      </form>
    </div>
  );
}