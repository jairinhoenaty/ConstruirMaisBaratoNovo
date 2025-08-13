import React from 'react';
import { MessageSquare, User, Mail, Phone, MapPin, Package, Calendar, ArrowRight } from 'lucide-react';

interface ProductMessage {
  id: string;
  productName: string;
  clientName: string;
  email: string;
  phone: string;
  location: string;
  message: string;
  date: string;
}

function ProductMessages() {
  // Mock data for product messages
  const messages: ProductMessage[] = [
    {
      id: '1',
      productName: 'Kit Ferramentas Profissional',
      clientName: 'Maria Silva',
      email: 'maria.silva@email.com',
      phone: '(11) 99999-9999',
      location: 'São Paulo, SP',
      message: 'Gostaria de saber mais detalhes sobre o Kit Ferramentas Profissional. Qual o prazo de entrega para minha região?',
      date: '2024-03-15 14:30'
    },
    {
      id: '2',
      productName: 'Furadeira de Impacto',
      clientName: 'João Santos',
      email: 'joao.santos@email.com',
      phone: '(21) 98888-8888',
      location: 'Rio de Janeiro, RJ',
      message: 'Tenho interesse na Furadeira de Impacto. Aceita pagamento parcelado?',
      date: '2024-03-14 16:45'
    },
    {
      id: '3',
      productName: 'Serra Circular',
      clientName: 'Ana Oliveira',
      email: 'ana.oliveira@email.com',
      phone: '(31) 97777-7777',
      location: 'Belo Horizonte, MG',
      message: 'Qual a garantia da Serra Circular? Faz entrega em Belo Horizonte?',
      date: '2024-03-14 09:15'
    }
  ];

  return (
    <div className="bg-white rounded-lg shadow-md p-8">
      <div className="flex items-center gap-3 mb-8">
        <MessageSquare className="w-8 h-8 text-blue-600" />
        <h2 className="text-2xl font-bold text-gray-900">
          Mensagens sobre Produtos
        </h2>
      </div>

      <div className="space-y-6">
        {messages.map((message) => (
          <div
            key={message.id}
            className="border border-gray-200 rounded-lg p-6 hover:border-blue-500 transition-colors"
          >
            <div className="flex justify-between items-start mb-4">
              <div className="flex items-center gap-2">
                <Package className="w-5 h-5 text-blue-600" />
                <h3 className="text-lg font-semibold text-blue-600">
                  {message.productName}
                </h3>
              </div>
              <div className="flex items-center gap-2 text-sm text-gray-500">
                <Calendar className="w-4 h-4" />
                <span>{new Date(message.date).toLocaleString('pt-BR')}</span>
              </div>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
              <div className="space-y-3">
                <div className="flex items-center gap-2">
                  <User className="w-5 h-5 text-gray-400" />
                  <div>
                    <span className="text-sm text-gray-500">Cliente:</span>
                    <p className="text-gray-900">{message.clientName}</p>
                  </div>
                </div>

                <div className="flex items-center gap-2">
                  <Mail className="w-5 h-5 text-gray-400" />
                  <div>
                    <span className="text-sm text-gray-500">Email:</span>
                    <p className="text-gray-900">{message.email}</p>
                  </div>
                </div>
              </div>

              <div className="space-y-3">
                <div className="flex items-center gap-2">
                  <Phone className="w-5 h-5 text-gray-400" />
                  <div>
                    <span className="text-sm text-gray-500">WhatsApp:</span>
                    <p className="text-gray-900">{message.phone}</p>
                  </div>
                </div>

                <div className="flex items-center gap-2">
                  <MapPin className="w-5 h-5 text-gray-400" />
                  <div>
                    <span className="text-sm text-gray-500">Localização:</span>
                    <p className="text-gray-900">{message.location}</p>
                  </div>
                </div>
              </div>
            </div>

            <div className="border-t border-gray-100 pt-4">
              <span className="text-sm text-gray-500">Mensagem:</span>
              <p className="mt-2 text-gray-900">{message.message}</p>
            </div>

            <div className="flex justify-end gap-3 mt-4">
              <a
                href={`https://wa.me/${message.phone.replace(/\D/g, '')}`}
                target="_blank"
                rel="noopener noreferrer"
                className="flex items-center gap-2 px-4 py-2 text-green-600 hover:bg-green-50 rounded-lg transition-colors"
              >
                <Phone className="w-5 h-5" />
                Responder via WhatsApp
              </a>
              <a
                href={`mailto:${message.email}`}
                className="flex items-center gap-2 px-4 py-2 text-blue-600 hover:bg-blue-50 rounded-lg transition-colors"
              >
                <Mail className="w-5 h-5" />
                Responder via Email
              </a>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}

export default ProductMessages;