import { FloatingWhatsApp } from "react-floating-whatsapp";

const Footer = () => {
  const chatboxStyle: React.CSSProperties = {
    color: "Black",
    bottom: 200,
  };
  const buttonStyle: React.CSSProperties = {
    right: 25,
  };
  const mainStyle: React.CSSProperties = {
    //right: 100,
    //bottom: 120,
    bottom: 400,
  };
  return (
    <div className="footer bg-blue-600 text-white text-center small">
      <span className="block sm:inline ml-2">
        construirmaisbarato.com.br - CNPJ: 59.887.176/0001-50
      </span>
      <FloatingWhatsApp
        phoneNumber="+551499166-5023"
        accountName={"Construir Mais Barato"}
        /*        onClick={() => {
          console.log("TESTE");
      }}
      */
        chatMessage="Olá! 🤝 Como podemos ajudar?"
        statusMessage="Geralmente responde em 1 hora"
        placeholder="Digite uma mensagem..."
        chatboxStyle={chatboxStyle}
        buttonStyle={buttonStyle}
        notificationStyle={mainStyle}
      />
    </div>
  );
};

export default Footer;
