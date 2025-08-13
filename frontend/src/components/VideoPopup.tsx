import React from 'react';
import { X } from 'lucide-react';

interface VideoPopupProps {
  isOpen: boolean;
  onClose: () => void;
  url: string;
}

function VideoPopup({ isOpen, onClose,url }: VideoPopupProps) {
  if (!isOpen) return null;
  return (
    <div className="fixed inset-0 bg-black bg-opacity-75 flex items-center justify-center z-50 p-4">
      <div className="bg-white rounded-xl shadow-xl w-full max-w-4xl relative">
        <button
          onClick={onClose}
          className="absolute -top-12 right-0 text-white hover:text-gray-200 transition-colors"
        >
          <X className="w-8 h-8" />
        </button>
        
        <div className="aspect-video w-full">
          <iframe
            className="w-full h-full rounded-xl"
            src={url}
            //src="https://www.youtube.com/embed/EpOykD8vDRU"//?autoplay=1&mute=1"
            title="Quem Somos - Construir Mais Barato"
            allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
            allowFullScreen
          ></iframe>
        </div>
      </div>
    </div>
  );
}

export default VideoPopup;