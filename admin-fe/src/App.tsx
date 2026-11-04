import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import { useAccount } from 'wagmi';
import { DEPLOYMENT } from '@shared/index';
import { LoginPage } from './pages/LoginPage';
import { DashboardPage } from './pages/DashboardPage';
import { PoolListPage } from './pages/PoolListPage';
import { CreatePoolPage } from './pages/CreatePoolPage';
import { ProposalListPage } from './pages/ProposalListPage';
import { Layout } from './components/Layout';

const ADMINS = DEPLOYMENT.admins.map((a) => a.toLowerCase());

function ProtectedRoute({ children }: { children: React.ReactNode }) {
  const { address, isConnected } = useAccount();

  if (!isConnected || !address) {
    return <Navigate to="/login" replace />;
  }

  if (!ADMINS.includes(address.toLowerCase())) {
    return (
      <div className="flex min-h-screen items-center justify-center bg-ink-900 p-4">
        <div className="card max-w-md text-center">
          <h1 className="text-xl font-semibold text-negative">无权限访问</h1>
          <p className="mt-3 text-sm text-slate-400">
            当前地址 <span className="font-mono text-xs">{address}</span> 不在管理员白名单中
          </p>
          <p className="mt-4 text-xs text-slate-500">请使用管理员账户连接钱包</p>
        </div>
      </div>
    );
  }

  return <Layout>{children}</Layout>;
}

export function App() {
  return (
    <BrowserRouter basename="/admin">
      <Routes>
        <Route path="/login" element={<LoginPage />} />
        <Route
          path="/"
          element={
            <ProtectedRoute>
              <DashboardPage />
            </ProtectedRoute>
          }
        />
        <Route
          path="/pools"
          element={
            <ProtectedRoute>
              <PoolListPage />
            </ProtectedRoute>
          }
        />
        <Route
          path="/pools/create"
          element={
            <ProtectedRoute>
              <CreatePoolPage />
            </ProtectedRoute>
          }
        />
        <Route
          path="/proposals"
          element={
            <ProtectedRoute>
              <ProposalListPage />
            </ProtectedRoute>
          }
        />
      </Routes>
    </BrowserRouter>
  );
}
