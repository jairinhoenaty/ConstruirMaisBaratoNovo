import { Building2 } from "lucide-react";

function CondoManagerTag() {
  return (
    <span
      className="inline-flex items-center gap-1 rounded-full bg-amber-100 px-2 py-0.5 text-xs font-medium text-amber-800"
      title="Síndico(a) ou representante de condomínio"
    >
      <Building2 className="w-3 h-3" />
      Síndico
    </span>
  );
}

export default CondoManagerTag;
