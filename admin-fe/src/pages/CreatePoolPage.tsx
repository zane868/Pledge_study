import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { DEPLOYMENT } from '@shared/index';

export function CreatePoolPage() {
  const navigate = useNavigate();
  const [formData, setFormData] = useState({
    lendToken: DEPLOYMENT.stable,
    borrowToken: '0x0000000000000000000000000000000000000000',
    interestRate: '10',
    martgageRate: '150',
    maxSupply: '10000',
    settleTime: '',
    endTime: '',
  });

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    // TODO: 实现创建池的多签提案逻辑
    alert('创建池功能开发中...');
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData((prev) => ({
      ...prev,
      [e.target.name]: e.target.value,
    }));
  };

  // 计算时间戳
  const now = new Date();
  const tomorrow = new Date(now.getTime() + 24 * 60 * 60 * 1000);
  const nextWeek = new Date(now.getTime() + 7 * 24 * 60 * 60 * 1000);

  const defaultSettleTime = tomorrow.toISOString().slice(0, 16);
  const defaultEndTime = nextWeek.toISOString().slice(0, 16);

  return (
    <div className="mx-auto max-w-2xl space-y-6">
      <div>
        <button
          className="text-sm text-slate-400 hover:text-slate-200"
          onClick={() => navigate(-1)}
        >
          ← 返回
        </button>
        <h2 className="mt-2 text-xl font-semibold text-white">创建新池</h2>
        <p className="mt-1 text-sm text-slate-400">
          填写池参数，创建多签提案等待其他管理员签名
        </p>
      </div>

      <form onSubmit={handleSubmit} className="card space-y-5">
        <div>
          <label className="label">借款资产地址（USDT）</label>
          <input
            type="text"
            name="lendToken"
            className="field font-mono"
            value={formData.lendToken}
            onChange={handleChange}
            placeholder="0x..."
            required
          />
          <p className="mt-1 text-xs text-slate-500">
            当前使用新部署的 TestUSDT（支持 permit）
          </p>
        </div>

        <div>
          <label className="label">抵押资产地址（原生 ETH）</label>
          <input
            type="text"
            name="borrowToken"
            className="field font-mono"
            value={formData.borrowToken}
            readOnly
          />
          <p className="mt-1 text-xs text-slate-500">
            零地址表示原生 ETH
          </p>
        </div>

        <div className="grid gap-5 md:grid-cols-2">
          <div>
            <label className="label">固定年利率（%）</label>
            <input
              type="number"
              name="interestRate"
              className="field tabular"
              value={formData.interestRate}
              onChange={handleChange}
              placeholder="10"
              step="0.01"
              min="0"
              max="1000"
              required
            />
            <p className="mt-1 text-xs text-slate-500">
              如 10% 输入 10
            </p>
          </div>

          <div>
            <label className="label">募集抵押率（%）</label>
            <input
              type="number"
              name="martgageRate"
              className="field tabular"
              value={formData.martgageRate}
              onChange={handleChange}
              placeholder="150"
              step="0.01"
              min="100"
              max="500"
              required
            />
            <p className="mt-1 text-xs text-slate-500">
              如 150% 输入 150
            </p>
          </div>
        </div>

        <div>
          <label className="label">出借上限（USDT）</label>
          <input
            type="number"
            name="maxSupply"
            className="field tabular"
            value={formData.maxSupply}
            onChange={handleChange}
            placeholder="10000"
            step="0.01"
            min="0"
            required
          />
        </div>

        <div className="grid gap-5 md:grid-cols-2">
          <div>
            <label className="label">募集截止时间</label>
            <input
              type="datetime-local"
              name="settleTime"
              className="field"
              value={formData.settleTime || defaultSettleTime}
              onChange={handleChange}
              required
            />
          </div>

          <div>
            <label className="label">到期时间</label>
            <input
              type="datetime-local"
              name="endTime"
              className="field"
              value={formData.endTime || defaultEndTime}
              onChange={handleChange}
              required
            />
          </div>
        </div>

        <div className="border-t border-slate-700 pt-5">
          <h3 className="text-sm font-medium text-white">预览参数</h3>
          <div className="mt-3 space-y-2 rounded bg-slate-800/50 p-3 font-mono text-xs text-slate-300">
            <div>年利率: {formData.interestRate}%</div>
            <div>抵押率: {formData.martgageRate}%</div>
            <div>上限: {formData.maxSupply} USDT</div>
            <div>募集截止: {formData.settleTime || defaultSettleTime}</div>
            <div>到期: {formData.endTime || defaultEndTime}</div>
          </div>
        </div>

        <div className="flex gap-3">
          <button type="button" className="btn-ghost" onClick={() => navigate(-1)}>
            取消
          </button>
          <button type="submit" className="btn-primary flex-1">
            创建多签提案
          </button>
        </div>
      </form>
    </div>
  );
}
