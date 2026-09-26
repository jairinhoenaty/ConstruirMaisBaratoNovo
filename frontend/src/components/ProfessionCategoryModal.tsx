import { HardHat, X } from "lucide-react";
import { IProfessionCategory } from "../interfaces/IProfessionCategory";

interface ProfessionCategoryModalProps {
  categories: IProfessionCategory[];
  loading: boolean;
  error: string;
  onSelect: (category: IProfessionCategory) => void;
  onClose: () => void;
}

// A lista vem da página: ela também alimenta o select de categoria, e buscar
// duas vezes deixaria os dois controles fora de sincronia.
function ProfessionCategoryModal({
  categories,
  loading,
  error,
  onSelect,
  onClose,
}: ProfessionCategoryModalProps) {
  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
      <div className="bg-white rounded-xl shadow-2xl w-full max-w-2xl max-h-[85vh] overflow-y-auto">
        <div className="flex items-center justify-between border-b border-gray-200 px-6 py-4">
          <h2 className="text-xl font-bold text-gray-900">
            O que você precisa?
          </h2>
          <button
            onClick={onClose}
            className="p-1 text-gray-400 hover:text-gray-700"
            aria-label="Fechar"
          >
            <X className="w-5 h-5" />
          </button>
        </div>

        <div className="p-6">
          {loading && <p className="text-gray-500">Carregando categorias...</p>}
          {!loading && error && <p className="text-red-600">{error}</p>}
          {!loading && !error && categories.length === 0 && (
            <p className="text-gray-500">Nenhuma categoria disponível.</p>
          )}

          <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
            {categories.map((category) => (
              <button
                key={category.id}
                onClick={() => onSelect(category)}
                className="flex items-center gap-4 p-4 text-left border border-gray-200 rounded-lg hover:border-blue-500 hover:bg-blue-50 transition-colors"
              >
                <div className="w-12 h-12 shrink-0 rounded-lg bg-blue-100 overflow-hidden flex items-center justify-center">
                  {category.icon ? (
                    <img
                      src={category.icon}
                      alt=""
                      className="w-full h-full object-cover"
                    />
                  ) : (
                    <HardHat className="w-6 h-6 text-blue-600" />
                  )}
                </div>
                <div className="min-w-0">
                  <p className="font-medium text-gray-900">{category.name}</p>
                  <p className="text-sm text-gray-500 truncate">
                    {category.description ||
                      `${category.profession_count} profissão(ões)`}
                  </p>
                  {category.client_travels && (
                    <p className="text-xs text-orange-600 mt-1">
                      Você vai até o profissional
                    </p>
                  )}
                </div>
              </button>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
}

export default ProfessionCategoryModal;
