import { useMemo, useState } from 'react';
import { useAccount, useReadContract, useWaitForTransactionReceipt, useWriteContract } from 'wagmi';
import { DEPLOYMENT, MULTISIG_ABI } from '@shared/index';

const MULTISIG = DEPLOYMENT.multisig as `0x${string}`;

/**
 * 本项目的多签合约不是 Safe，而是一个自定义的 multiSignature 合约，流程如下：
 *
 * 1. 申请人（applicant）调用 `createApplication(pool)` 创建一个授权申请；
 *    申请的唯一标识 msghash = keccak256(applicant, pool)
 * 2. 其他管理员调用 `signApplication(msghash)` 签名（只能签 index 0 的申请）
 * 3. 签名数达到 threshold 后，`getValidSignature(msghash, 0) > 0`
 * 4. 申请人再调用 pool 的受限方法（settle / finish），
 *    pool 的 validCall 修饰符会校验 keccak256(msg.sender, pool) 是否已被多签批准
 *
 * 注意：合约没有提供 msghash 的枚举接口，所以无法列出所有申请，
 * 只能按「申请人地址 + 目标合约」计算 msghash 后查询。
 */
export function useMultisigApproval(applicant?: string, to?: string) {
  const { address } = useAccount();
  const [error, setError] = useState<string | null>(null);
  const [isPending, setIsPending] = useState(false);
  const [txHash, setTxHash] = useState<string | null>(null);

  const { writeContractAsync } = useWriteContract();
  const { isLoading: isMining } = useWaitForTransactionReceipt({ hash: txHash as `0x${string}` });

  // msghash = keccak256(abi.encodePacked(from, to))
  const { data: msgHash } = useReadContract({
    address: MULTISIG,
    abi: MULTISIG_ABI,
    functionName: 'getApplicationHash',
    args: [applicant as `0x${string}`, to as `0x${string}`],
    query: { enabled: !!applicant && !!to },
  });

  const hash = msgHash as `0x${string}` | undefined;

  // 申请总数（一般只有 1 条）
  const { data: appCount } = useReadContract({
    address: MULTISIG,
    abi: MULTISIG_ABI,
    functionName: 'getApplicationCount',
    args: [hash!],
    query: { enabled: !!hash },
  });

  // index 0 的申请信息：申请人 + 已签名地址列表
  const { data: appInfo, refetch: refetchAppInfo } = useReadContract({
    address: MULTISIG,
    abi: MULTISIG_ABI,
    functionName: 'getApplicationInfo',
    args: [hash!, 0n],
    query: { enabled: !!hash && Number(appCount ?? 0n) > 0 },
  });

  // 是否已达阈值（>0 表示已批准）
  const { data: validSignature } = useReadContract({
    address: MULTISIG,
    abi: MULTISIG_ABI,
    functionName: 'getValidSignature',
    args: [hash!, 0n],
    query: { enabled: !!hash },
  });

  const { data: threshold } = useReadContract({
    address: MULTISIG,
    abi: MULTISIG_ABI,
    functionName: 'threshold',
  });

  const applicantAddr = (appInfo as any[])?.[0] as string | undefined;
  const signers = ((appInfo as any[])?.[1] as string[]) ?? [];

  const required = Number(threshold ?? DEPLOYMENT.threshold ?? 2);
  const approved = Number(validSignature ?? 0) > 0;

  const isApplicant = !!address && !!applicant && address.toLowerCase() === applicant.toLowerCase();
  const hasSigned = !!address && signers.some((s) => s.toLowerCase() === address.toLowerCase());

  const createApplication = async () => {
    if (!address) {
      setError('请先连接钱包');
      return null;
    }
    setIsPending(true);
    setError(null);
    try {
      const h = await writeContractAsync({
        address: MULTISIG,
        abi: MULTISIG_ABI,
        functionName: 'createApplication',
        args: [to as `0x${string}`],
      });
      setTxHash(h);
      setIsPending(false);
      return h;
    } catch (e: any) {
      console.error('创建授权申请失败:', e);
      setError(e.shortMessage || e.message || '创建授权申请失败');
      setIsPending(false);
      return null;
    }
  };

  const signApplication = async () => {
    if (!hash) {
      setError('无法计算 msghash');
      return null;
    }
    setIsPending(true);
    setError(null);
    try {
      const h = await writeContractAsync({
        address: MULTISIG,
        abi: MULTISIG_ABI,
        functionName: 'signApplication',
        args: [hash],
      });
      setTxHash(h);
      setIsPending(false);
      return h;
    } catch (e: any) {
      console.error('签名失败:', e);
      setError(e.shortMessage || e.message || '签名失败');
      setIsPending(false);
      return null;
    }
  };

  const revokeSign = async () => {
    if (!hash) return null;
    setIsPending(true);
    setError(null);
    try {
      const h = await writeContractAsync({
        address: MULTISIG,
        abi: MULTISIG_ABI,
        functionName: 'revokeSignApplication',
        args: [hash],
      });
      setTxHash(h);
      setIsPending(false);
      return h;
    } catch (e: any) {
      console.error('撤销签名失败:', e);
      setError(e.shortMessage || e.message || '撤销签名失败');
      setIsPending(false);
      return null;
    }
  };

  const shortHash = useMemo(() => (hash ? `${hash.slice(0, 10)}…${hash.slice(-8)}` : ''), [hash]);

  return {
    hash,
    shortHash,
    appCount: Number(appCount ?? 0),
    applicant: applicantAddr,
    signers,
    signCount: signers.length,
    required,
    approved,
    isApplicant,
    hasSigned,
    isPending: isPending || isMining,
    txHash,
    error,
    createApplication,
    signApplication,
    revokeSign,
    refetch: async () => {
      await refetchAppInfo();
    },
  };
}
