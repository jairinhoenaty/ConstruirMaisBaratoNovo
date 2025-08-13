import React from 'react';
import { X, Smartphone } from 'lucide-react';

interface AppDownloadPopupProps {
  isOpen: boolean;
  onClose: () => void;
}

function AppDownloadPopup({ isOpen, onClose }: AppDownloadPopupProps) {
  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50">
      <div className="bg-white rounded-xl shadow-xl max-w-2xl w-full relative overflow-hidden">
        {/* Close Button */}
        <button
          onClick={onClose}
          className="absolute top-4 right-4 text-gray-500 hover:text-gray-700 transition-colors"
        >
          <X className="w-6 h-6" />
        </button>

        {/*
        <div className="grid md:grid-cols-2 gap-8 p-8">
          <div className="flex flex-col justify-center">
            <div className="flex items-center gap-2 mb-4">
              <Smartphone className="w-8 h-8 text-blue-600" />
              <h3 className="text-2xl font-bold text-gray-900">
                Profissional Já
              </h3>
            </div>
            
            <p className="text-gray-600 mb-6">
              Baixando o nosso aplicativo você tem acesso em tempo real aos profissionais na sua região e pode solicitar um profissional na mesma hora.
            </p>
            <div className="space-y-4">
              <button className="w-full bg-black text-white rounded-lg py-3 px-4 flex items-center justify-center gap-2 hover:bg-gray-900 transition-colors">
                <img src="https://upload.wikimedia.org/wikipedia/commons/f/fa/Apple_logo_black.svg" alt="App Store" className="w-5 h-5" />
                App Store
              </button>
              <button className="w-full bg-black text-white rounded-lg py-3 px-4 flex items-center justify-center gap-2 hover:bg-gray-900 transition-colors">
                <img src="https://upload.wikimedia.org/wikipedia/commons/d/d0/Google_Play_Arrow_logo.svg" alt="Play Store" className="w-5 h-5" />
                Play Store
              </button>
            </div>
          </div>
          

        
          <div className="relative h-[300px] md:h-full min-h-[300px] rounded-lg overflow-hidden">
            <img
              src="https://images.unsplash.com/photo-1605152276897-4f618f831968"
              alt="App Interface"
              className="absolute inset-0 w-full h-full object-cover"
            />
            <div className="absolute inset-0 bg-gradient-to-t from-black/50 to-transparent"></div>
          </div>
      </div>
        */}
      </div>
    </div>
  );
}

export default AppDownloadPopup;