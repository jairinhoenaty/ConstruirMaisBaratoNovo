import React from "react";
import { ArrowLeft } from "lucide-react";
import { useNavigate } from "react-router-dom";

function PrivacyPolicy() {
  const navigate = useNavigate();
  return (
    <div className="min-h-screen bg-gray-50 py-12">
      <div className="max-w-4xl mx-auto px-4">
        <button
          onClick={() => navigate("/")}
          className="flex items-center gap-2 text-gray-600 hover:text-blue-600 mb-8 transition-colors"
        >
          <ArrowLeft className="w-5 h-5" />
          Voltar para página inicial
        </button>

        <div className="bg-white rounded-xl shadow-xl p-8">
          <h1 className="text-3xl font-bold text-gray-900 mb-8">
            Política de Privacidade
          </h1>

          <div className="prose prose-lg max-w-none text-gray-600 space-y-6">
            <p>
              Sua privacidade é muito importante para nós! Esta Política de
              Privacidade esclarece como é feito o tratamento dos seus dados
              pessoais a partir da nossa ferramenta. Assim, prezamos pela
              transparência entre nossa equipe e você, nosso usuário,
              fortalecendo nossa parceria e relação de confiança. Nesse sentido,
              gostaríamos de tranquilizá-los, pois estamos totalmente adequados
              à Lei Geral de Proteção de Dados do Brasil – LGPD (Lei n°
              13.709/2018), conforme podem conferir os termos abaixo
              estipulados.
            </p>

            <h2 className="text-2xl font-bold text-gray-900 mt-8">
              Quem somos?
            </h2>
            <p>
              Mais que um site, a C + B é uma plataforma online que busca reunir
              prestadores de serviços e clientes de uma forma rápida e barata,
              facilitando o encontro entre profissional e sua obra.
            </p>
            <p>
              O nosso contato é realizado por meio do e-mail:
              atendimento@construirmaisbarato.com.br
            </p>
            <p>
              Nós temos também um responsável pela proteção de dados, portanto,
              quaisquer dúvidas ou solicitações sobre o uso de seus dados
              pessoais devem ser encaminhadas para o nosso encarregado de dados:
            </p>
            <p>Jairo Assis lgpd@construirmaisbarato.com.br (14) 98835-0791</p>

            <h2 className="text-2xl font-bold text-gray-900 mt-8">
              COMO USAMOS OS SEUS DADOS:
            </h2>
            <p>
              Nosso site pode ser utilizado para áreas como construção, pintura,
              elétrica e reparos hidráulicos. Podem oferecer serviços em nosso
              site profissionais com CNPJ, MEI ou autônomos. Os usuários
              (cliente final) poderão ser pessoas jurídicas ou físicas. Ao fazer
              o cadastro em nossa plataforma (site/aplicativo), coletaremos
              algumas informações que serão fornecidas exclusivamente pelo
              usuário. Todavia, esclarecemos que essas informações são
              basicamente cadastrais, como as seguintes informações: nome,
              e-mail, CPF, endereço e telefone. Quando solicitado o endereço,
              este se refere ao local da prestação de serviço a ser realizado.
              Menores de idade não poderão utilizar nossos serviços. Ressaltamos
              que a exclusão dos dados de nossa ferramenta é perfeitamente
              possível.
            </p>
            <p>
              Usamos essas informações exclusivamente para a funcionalidade de
              nosso sistema. Também podemos lhe enviar e-mails. Faremos isso com
              base em nosso interesse legítimo em fornecer informações precisas
              e um serviço de qualidade. Caso não queira receber nossos e-mails,
              basta realizar o descadastramento em nosso site.
            </p>
            <p>
              Suas informações são armazenadas em nosso servidor e será tratada
              apenas em decorrência da nossa prestação de serviços. Não
              comercial
            </p>

            <h2 className="text-2xl font-bold text-gray-900 mt-8">COOKIES</h2>
            <p>
              Quando você usa nosso site para navegar em nossos serviços, vários
              cookies são usados por nós e por terceiros para permitir que o
              site funcione, para coletar informações úteis sobre os visitantes,
              ajudando a tornar sua experiência de usuário melhor.
            </p>
            <p>
              Alguns dos cookies que usamos são estritamente necessários para o
              funcionamento do nosso site, e não pedimos o seu consentimento
              para colocá-los no seu computador. No entanto, para os cookies que
              são úteis, mas não estritamente necessários, pediremos sempre o
              seu consentimento antes de os colocar.
            </p>

            <h2 className="text-2xl font-bold text-gray-900 mt-8">
              Do Compartilhamento
            </h2>
            <p>
              Seus dados são armazenados em nosso banco de dados, mas não serão
              compartilhados com terceiros, a não ser nos casos previstos em
              Lei.
            </p>

            <h2 className="text-2xl font-bold text-gray-900 mt-8">
              Dos Serviços
            </h2>
            <p>
              A função da nossa plataforma é facilitar o encontro entre
              profissionais e clientes, meramente informativo e consultivo, no
              estilo "páginas amarelas" das listas telefônicas. Toda e qualquer
              negociação realizada entre as partes é de responsabilidade delas.
              Nosso site NÃO se responsabiliza por defeitos na prestação dos
              serviços contratados pelo usuário.
            </p>

            <h2 className="text-2xl font-bold text-gray-900 mt-8">
              Do armazenamento e segurança
            </h2>
            <p>
              Utilizamos técnicas e softwares seguros e renomados para o
              armazenamento de todas as informações que transitam pelo site.
              Assim, garantimos a utilização de medidas técnicas e
              administrativas aptas a proteger os dados pessoais de acessos não
              autorizados e de situações acidentais ou ilícitas de destruição,
              perda, alteração, comunicação ou difusão de seus dados.
            </p>

            <h2 className="text-2xl font-bold text-gray-900 mt-8">
              Seus direitos como titular de dados
            </h2>
            <p>
              Por lei, qualquer indivíduo poderá nos perguntar quais são as
              informações que temos sobre ele em nosso banco de dados, além de
              ser garantido o direito de correção, se as informações estiverem
              imprecisas, por meio do e-mail lgpd@construirmaisbarato.com.br. Se
              solicitarmos o seu consentimento para processar seus dados, você
              poderá retirar esse consentimento a qualquer momento, bem como
              solicitar a exclusão de dados. Caso queira enviar uma solicitação
              sobre a utilização de seus dados pessoais (informações, correções
              e exclusão), use o endereço eletrônico fornecido nesta política.
            </p>

            <h2 className="text-2xl font-bold text-gray-900 mt-8">
              Atualizações para esta política de privacidade
            </h2>
            <p>
              Revisamos regularmente e, se apropriado, atualizaremos esta
              política de privacidade de tempos em tempos, e conforme nossos
              serviços e uso de dados sejam alterados. Se, eventualmente,
              usarmos seus dados pessoais de uma forma que não identificada ou
              descrita anteriormente, entraremos em contato para fornecer
              informações sobre isso e, se necessário, solicitar o seu
              consentimento.
            </p>
          </div>
        </div>
      </div>
    </div>
  );
}

export default PrivacyPolicy;
