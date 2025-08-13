import React, { useState } from 'react';
import { BarChart3, DollarSign, Users, Package, Star, Calendar, Home, ShoppingBag, Search, Shield, LogIn, Box } from 'lucide-react';
import AccountSettings from './AccountSettings';
import QuotesPanel from './QuotesPanel';
import ClientMessages from './ClientMessages';
import PasswordPanel from './PasswordPanel';

function DashboardNovo() {
  const [activeTab, setActiveTab] = useState('profile');



  // Mock data for dashboard statistics
  const stats = [
    {
      title: 'Vendas Totais',
      value: 'R$ 24.780,00',
      change: '+12%',
      icon: DollarSign,
      color: 'text-green-600',
      bgColor: 'bg-green-100'
    },
    {
      title: 'Clientes Ativos',
      value: '1,483',
      change: '+8%',
      icon: Users,
      color: 'text-blue-600',
      bgColor: 'bg-blue-100'
    },
    {
      title: 'Produtos Vendidos',
      value: '245',
      change: '+23%',
      icon: Package,
      color: 'text-purple-600',
      bgColor: 'bg-purple-100'
    },
    {
      title: 'Avaliação Média',
      value: '4.8',
      change: '+2%',
      icon: Star,
      color: 'text-yellow-600',
      bgColor: 'bg-yellow-100'
    }
  ];

  const recentSales = [
    {
      product: 'Kit Ferramentas Profissional',
      amount: 'R$ 599,90',
      date: '2024-03-15',
      status: 'Concluído'
    },
    {
      product: 'Furadeira de Impacto',
      amount: 'R$ 299,90',
      date: '2024-03-14',
      status: 'Em andamento'
    },
    {
      product: 'Serra Circular',
      amount: 'R$ 449,90',
      date: '2024-03-13',
      status: 'Concluído'
    },
    {
      product: 'Conjunto de Pincéis',
      amount: 'R$ 49,90',
      date: '2024-03-12',
      status: 'Concluído'
    }
  ];

  const handleNavigation = (page: string) => {
      window.location.href = `/${page}`;
  };

  return (
    <div className="p-8">
      <div className="flex items-center justify-between mb-8">
        <div className="flex items-center gap-3">
          <BarChart3 className="w-8 h-8 text-blue-600" />
          <h2 className="text-2xl font-bold text-gray-900">
            Dashboard
          </h2>
        </div>
        
        {/* Navigation Buttons */}
        <div className="flex gap-3">
          <button
            onClick={() => handleNavigation('home')}
            className="flex items-center gap-2 px-4 py-2 bg-white rounded-lg shadow-md hover:bg-gray-50 transition-colors text-gray-700"
          >
            <Home className="w-5 h-5" />
            <span>Página Inicial</span>
          </button>
          <button
            onClick={() => handleNavigation('marketplace')}
            className="flex items-center gap-2 px-4 py-2 bg-white rounded-lg shadow-md hover:bg-gray-50 transition-colors text-gray-700"
          >
            <ShoppingBag className="w-5 h-5" />
            <span>Marketplace</span>
          </button>
          <button
            onClick={() => handleNavigation('search')}
            className="flex items-center gap-2 px-4 py-2 bg-white rounded-lg shadow-md hover:bg-gray-50 transition-colors text-gray-700"
          >
            <Search className="w-5 h-5" />
            <span>Busca de Profissionais</span>
          </button>
          <button
            onClick={() => handleNavigation('privacy')}
            className="flex items-center gap-2 px-4 py-2 bg-white rounded-lg shadow-md hover:bg-gray-50 transition-colors text-gray-700"
          >
            <Shield className="w-5 h-5" />
            <span>Política de Privacidade</span>
          </button>
          <button
            onClick={() => handleNavigation('login')}
            className="flex items-center gap-2 px-4 py-2 bg-blue-600 text-white rounded-lg shadow-md hover:bg-blue-700 transition-colors"
          >
            <LogIn className="w-5 h-5" />
            <span>Login</span>
          </button>
        </div>
      </div>

      {/* Statistics Grid */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
        {stats.map((stat, index) => {
          const Icon = stat.icon;
          return (
            <div
              key={index}
              className="bg-white rounded-lg shadow-md p-6"
            >
              <div className="flex items-center justify-between mb-4">
                <div className={`${stat.bgColor} p-3 rounded-lg`}>
                  <Icon className={`w-6 h-6 ${stat.color}`} />
                </div>
                <span className={`text-sm font-medium ${
                  stat.change.startsWith('+') ? 'text-green-600' : 'text-red-600'
                }`}>
                  {stat.change}
                </span>
              </div>
              <h3 className="text-2xl font-bold text-gray-900 mb-1">
                {stat.value}
              </h3>
              <p className="text-sm text-gray-600">
                {stat.title}
              </p>
            </div>
          );
        })}
      </div>

      {/* Quick Access Buttons */}
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mb-8">
        <button
          onClick={() => handleNavigation('received-products')}
          className="flex items-center justify-center gap-2 p-4 bg-white rounded-lg shadow-md hover:bg-gray-50 transition-colors"
        >
          <Box className="w-6 h-6 text-blue-600" />
          <span className="font-medium">Produtos Recebidos</span>
        </button>
      </div>

      {/* Recent Sales */}
      <div className="bg-white rounded-lg shadow-md p-6">
        <div className="flex items-center justify-between mb-6">
          <h3 className="text-lg font-semibold text-gray-900">
            Vendas Recentes
          </h3>
          <button className="text-blue-600 hover:text-blue-700 text-sm font-medium">
            Ver todas
          </button>
        </div>

        <div className="overflow-x-auto">
          <table className="w-full">
            <thead>
              <tr className="border-b border-gray-200">
                <th className="text-left py-3 px-4 text-sm font-medium text-gray-600">Produto</th>
                <th className="text-left py-3 px-4 text-sm font-medium text-gray-600">Valor</th>
                <th className="text-left py-3 px-4 text-sm font-medium text-gray-600">Data</th>
                <th className="text-left py-3 px-4 text-sm font-medium text-gray-600">Status</th>
              </tr>
            </thead>
            <tbody>
              {recentSales.map((sale, index) => (
                <tr key={index} className="border-b border-gray-100 last:border-0">
                  <td className="py-3 px-4">
                    <span className="text-gray-900">{sale.product}</span>
                  </td>
                  <td className="py-3 px-4">
                    <span className="text-gray-900 font-medium">{sale.amount}</span>
                  </td>
                  <td className="py-3 px-4">
                    <div className="flex items-center gap-2">
                      <Calendar className="w-4 h-4 text-gray-400" />
                      <span className="text-gray-600">
                        {new Date(sale.date).toLocaleDateString('pt-BR')}
                      </span>
                    </div>
                  </td>
                  <td className="py-3 px-4">
                    <span className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${
                      sale.status === 'Concluído'
                        ? 'bg-green-100 text-green-800'
                        : 'bg-yellow-100 text-yellow-800'
                    }`}>
                      {sale.status}
                    </span>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
}

export default DashboardNovo;