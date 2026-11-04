import { useAccount } from 'wagmi';
import { DEPLOYMENT } from '@shared/index';
import { useMultisigApproval } from '@/hooks/useMultisigApproval';

/**
 * 本项目的多签为自定义 multiSignature 合约：
 * 申请以 (申请人地址, 目标合约) 为唯一键，msghash = keccak256(from, to)。
 * 合约没有提供枚举全部申请的接口，因此这里展示「当前登录管理员对 Pool 合约」的授权申请。
 */
export function ProposalListPage() {
  const { address } = useAccount();
  const approval = useMultisigApproval(address, DEPLOYMENT.pool);

  return (
    <div className="space-y-6">
      <div>
        <h2 className="text-xl font-semibold text-white">多签授权申请</h2>
        <p className="mt-1 text-sm text-slate-400">
          签名阈值 {approval.required} · 管理员 {DEPLOYMENT.admins.length} 位
        </p>
      </div>

      {approval.error && (
        <div className="rounded bg-red-500/10 p-3 text-sm text-red-400">❌ {approval.error}</div>
      )}

      <div className="card space-y-3">
        <div className="flex items-center justify-between">
          <h3 className="font-medium text-white">Pool 合约操作授权</h3>
          <span
            className={`rounded px-2 py-0.5 text-xs font-medium ${
              approval.approved
                ? 'bg-green-500/20 text-green-400'
                : 'bg-amber-500/20 text-amber-400'
            }`}
          >
            {approval.approved ? '已批准' : '待签名'}
          </span>
        </div>

        <div className="space-y-2 text-sm">
          <div>
            <div className="text-xs text-slate-500">申请人</div>
            <div className="mt-0.5 font-mono text-xs text-slate-300">
              {approval.applicant ?? (address ? address : '未连接钱包')}
            </div>
          </div>
          <div>
            <div className="text-xs text-slate-500">目标合约</div>
            <div className="mt-0.5 font-mono text-xs text-slate-300">{DEPLOYMENT.pool}</div>
          </div>
          <div>
            <div className="text-xs text-slate-500">msghash</div>
            <div className="mt-0.5 break-all font-mono text-xs text-slate-300">
              {approval.hash ?? '计算中…'}
            </div>
          </div>
          <div>
            <div className="text-xs text-slate-500">
              签名进度（{approval.signCount}/{approval.required}）
            </div>
            <div className="mt-1 flex flex-wrap gap-1.5">
              {approval.signers.length === 0 ? (
                <span className="text-xs text-slate-500">暂无签名</span>
              ) : (
                approval.signers.map((s) => (
                  <span
                    key={s}
                    className="rounded bg-green-500/10 px-2 py-0.5 font-mono text-xs text-green-400"
                  >
                    {s}
                  </span>
                ))
              )}
            </div>
          </div>
        </div>

        <div className="flex gap-2 pt-2">
          {approval.appCount === 0 ? (
            <button
              className="btn-primary text-sm"
              onClick={approval.createApplication}
              disabled={approval.isPending}
            >
              {approval.isPending ? '提交中…' : '创建授权申请'}
            </button>
          ) : !approval.hasSigned ? (
            <button
              className="btn-primary text-sm"
              onClick={approval.signApplication}
              disabled={approval.isPending}
            >
              {approval.isPending ? '签名中…' : '签名批准'}
            </button>
          ) : (
            <button
              className="btn-ghost text-sm"
              onClick={approval.revokeSign}
              disabled={approval.isPending}
            >
              撤销签名
            </button>
          )}
        </div>
      </div>
    </div>
  );
}
