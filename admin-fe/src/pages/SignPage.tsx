import { useState } from 'react';
import { useSearchParams } from 'react-router-dom';
import { useAccount } from 'wagmi';
import { DEPLOYMENT, shortAddress } from '@shared/index';
import { useMultisigApproval } from '@/hooks/useMultisigApproval';

const ADMINS = DEPLOYMENT.admins.map((a) => a.toLowerCase());

/**
 * 独立签名页：通过链接直达某个凭证的多签授权签名。
 * 用法：/admin/sign?to=0x凭证地址&label=SP
 *
 * 申请人固定为 operator（msghash = keccak256(operator, to)），
 * 打开链接的管理员用自己的钱包连入后点「签名批准」即可，
 * 无需进入创建向导、也无需理解申请创建流程。
 *
 * 注意：signApplication 是链上交易（onlyOwner），必须由管理员本人钱包签名，
 * 无法代签；此页面只是把「找到入口 + 连钱包 + 签名」收敛成一条链接。
 */
export function SignPage() {
  const [params] = useSearchParams();
  const to = params.get('to') ?? '';
  const label = params.get('label') ?? '凭证';

  const validAddr = /^0x[0-9a-fA-F]{40}$/.test(to);

  return (
    <div className="mx-auto max-w-2xl space-y-6">
      <div>
        <h2 className="text-xl font-semibold text-white">多签授权签名</h2>
        <p className="mt-1 text-sm text-slate-400">
          对下方凭证的授权申请签名。签名由你当前连接的钱包直接发起链上交易。
        </p>
      </div>

      {!validAddr ? (
        <div className="card text-sm text-red-400">❌ 链接无效：缺少或非法的凭证地址参数（to）。</div>
      ) : (
        <SignCard to={to} label={label} />
      )}
    </div>
  );
}

function SignCard({ to, label }: { to: string; label: string }) {
  const { address } = useAccount();
  const approval = useMultisigApproval(DEPLOYMENT.operator, to);

  // 签名交易确认后强制重新挂载，刷新链上签名进度
  const [epoch, setEpoch] = useState(0);
  const [justSigned, setJustSigned] = useState(false);
  const wasPending = approval.isPending;

  const isAdmin = !!address && ADMINS.includes(address.toLowerCase());
  const isApplicant = !!address && address.toLowerCase() === DEPLOYMENT.operator.toLowerCase();

  const doSign = async () => {
    setJustSigned(false);
    const h = await approval.signApplication();
    if (h) setJustSigned(true);
  };

  // 签名交易从 pending 转为完成时，刷新读取
  if (wasPending && !approval.isPending && justSigned) {
    setJustSigned(false);
    setEpoch((e) => e + 1);
  }

  return (
    <div className="card space-y-4" key={epoch}>
      <div className="flex items-center justify-between">
        <h3 className="font-medium text-white">{label} 凭证授权</h3>
        <span
          className={`rounded px-2 py-0.5 text-xs font-medium ${
            approval.approved ? 'bg-green-500/20 text-green-400' : 'bg-amber-500/20 text-amber-400'
          }`}
        >
          {approval.approved ? '已批准' : `待签名 ${approval.signCount}/${approval.required}`}
        </span>
      </div>

      <div className="space-y-2 text-sm">
        <div>
          <div className="text-xs text-slate-500">凭证地址</div>
          <a
            className="mt-0.5 block break-all font-mono text-xs text-accent-400 hover:underline"
            href={`${DEPLOYMENT.explorer}/address/${to}`}
            target="_blank"
            rel="noreferrer"
          >
            {to}
          </a>
        </div>
        <div>
          <div className="text-xs text-slate-500">申请人（operator）</div>
          <div className="mt-0.5 font-mono text-xs text-slate-300">{DEPLOYMENT.operator}</div>
        </div>
        <div>
          <div className="text-xs text-slate-500">msghash</div>
          <div className="mt-0.5 break-all font-mono text-xs text-slate-400">
            {approval.hash ?? '计算中…'}
          </div>
        </div>
        <div>
          <div className="text-xs text-slate-500">已签名管理员</div>
          <div className="mt-1 flex flex-wrap gap-1.5">
            {approval.signers.length === 0 ? (
              <span className="text-xs text-slate-500">暂无签名</span>
            ) : (
              approval.signers.map((s) => (
                <span key={s} className="rounded bg-green-500/10 px-2 py-0.5 font-mono text-xs text-green-400">
                  {shortAddress(s)}
                </span>
              ))
            )}
          </div>
        </div>
      </div>

      {approval.error && <p className="text-sm text-red-400">❌ {approval.error}</p>}

      <div className="border-t border-slate-700 pt-4">
        {!address ? (
          <p className="text-sm text-amber-400">请先连接钱包（右上角），再签名。</p>
        ) : !isAdmin ? (
          <p className="text-sm text-red-400">
            当前钱包 <span className="font-mono">{shortAddress(address)}</span> 不是管理员，无法签名。
            请切换到管理员账户。
          </p>
        ) : isApplicant ? (
          <p className="text-sm text-slate-400">
            当前钱包是申请人 operator。签名需由其他管理员完成，请用另外两个管理员账户打开本链接。
          </p>
        ) : approval.approved ? (
          <p className="text-sm text-green-400">✅ 该授权已达标（{approval.signCount}/{approval.required}），无需再签。</p>
        ) : approval.hasSigned ? (
          <div className="flex items-center gap-3">
            <span className="text-sm text-green-400">✅ 你已签名</span>
            <button className="btn-ghost text-sm" onClick={approval.revokeSign} disabled={approval.isPending}>
              撤销签名
            </button>
          </div>
        ) : (
          <button className="btn-primary w-full" onClick={doSign} disabled={approval.isPending}>
            {approval.isPending ? '签名中…' : `签名批准 ${label}`}
          </button>
        )}
      </div>

      {approval.txHash && (
        <a
          className="block text-xs text-accent-400 hover:underline"
          href={`${DEPLOYMENT.explorer}/tx/${approval.txHash}`}
          target="_blank"
          rel="noreferrer"
        >
          查看签名交易 →
        </a>
      )}
    </div>
  );
}
