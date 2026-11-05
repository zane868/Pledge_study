import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import { LoginPage } from './pages/LoginPage';
import { DashboardPage } from './pages/DashboardPage';
import { PoolListPage } from './pages/PoolListPage';
import { CreatePoolPage } from './pages/CreatePoolPage';
import { ProposalListPage } from './pages/ProposalListPage';
import { SignPage } from './pages/SignPage';
import { Layout } from './components/Layout';
import { isLoggedIn } from './lib/auth';

/**
 * 页面访问只要求「账号密码已登录」。
 * 涉及链上写操作时，各操作组件会自行检测钱包连接与管理员身份。
 */
function ProtectedRoute({ children }: { children: React.ReactNode }) {
  if (!isLoggedIn()) {
    return <Navigate to="/login" replace />;
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
        <Route
          path="/sign"
          element={
            <ProtectedRoute>
              <SignPage />
            </ProtectedRoute>
          }
        />
      </Routes>
    </BrowserRouter>
  );
}