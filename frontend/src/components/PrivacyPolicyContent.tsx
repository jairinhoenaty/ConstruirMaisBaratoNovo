import type { ReactNode } from "react";

// Documento jurídico vigente, usado pela página /privacy e pelo modal do
// cadastro. As cláusulas 11 e 12 dos Termos trazem o preço da taxa: se ele
// mudar, elas mudam junto — ver a linha `solicitation` da tabela `plans`.

const DocumentTitle = ({ children }: { children: ReactNode }) => (
  <h2 className="text-2xl font-bold text-gray-900 mt-10 mb-4 first:mt-0">
    {children}
  </h2>
);

const Section = ({ children }: { children: ReactNode }) => (
  <h3 className="text-xl font-bold text-gray-900 mt-8 mb-3">{children}</h3>
);

const SubSection = ({ children }: { children: ReactNode }) => (
  <h4 className="text-lg font-semibold text-gray-800 mt-5 mb-2">{children}</h4>
);

const P = ({ children }: { children: ReactNode }) => (
  <p className="mb-3 leading-relaxed">{children}</p>
);

const List = ({ items }: { items: string[] }) => (
  <ul className="list-disc pl-6 mb-4 space-y-1">
    {items.map((item) => (
      <li key={item}>{item}</li>
    ))}
  </ul>
);

const Email = () => (
  <a
    href="mailto:hassisconecta@gmail.com"
    className="text-blue-600 hover:text-blue-800 underline"
  >
    hassisconecta@gmail.com
  </a>
);

const Divider = () => <hr className="my-10 border-gray-200" />;

function PrivacyPolicyContent() {
  return (
    <div className="text-gray-600">
      <p className="text-sm italic text-gray-500 mb-6">
        Última atualização: 13/09/2026
      </p>

      <P>
        A Hassis Conecta valoriza a privacidade, a transparência e a segurança
        das informações de seus usuários.
      </P>
      <P>
        Esta página reúne, em um único local, dois documentos independentes:
      </P>
      <P>
        <strong>POLÍTICA DE PRIVACIDADE</strong> — explica como a Hassis Conecta
        coleta, utiliza, armazena, protege e compartilha dados pessoais.
      </P>
      <P>
        <strong>TERMOS DE USO</strong> — estabelecem as regras para utilização
        da plataforma Hassis Conecta, bem como os direitos, deveres e
        responsabilidades dos clientes, profissionais, lojistas e da própria
        plataforma.
      </P>
      <P>
        Ao utilizar a Hassis Conecta, o usuário declara que leu e compreendeu as
        condições aplicáveis à utilização da plataforma.
      </P>

      <Divider />
      <DocumentTitle>POLÍTICA DE PRIVACIDADE</DocumentTitle>

      <Section>1. QUEM SOMOS</Section>
      <P>
        A Hassis Conecta é uma plataforma tecnológica destinada a facilitar a
        conexão entre clientes, profissionais e lojistas.
      </P>
      <P>
        A plataforma pode permitir, entre outras funcionalidades, que usuários
        encontrem profissionais, solicitem serviços, recebam propostas ou
        orçamentos, consultem produtos e preços, encontrem estabelecimentos e
        profissionais próximos, troquem informações e iniciem contatos para
        negociação direta.
      </P>
      <P>
        A Hassis Conecta atua como plataforma de conexão e tecnologia e não
        executa diretamente os serviços realizados pelos profissionais
        cadastrados, salvo quando expressamente informado em determinada
        funcionalidade.
      </P>
      <P>
        Para assuntos relacionados à privacidade e proteção de dados, o contato
        atualmente disponibilizado é:
      </P>
      <P>
        <strong>Responsável/Canal de Privacidade:</strong> Jairo Assis
        <br />
        <strong>E-mail:</strong> <Email />
        <br />
        <strong>Telefone:</strong> 14 99166-5023
      </P>

      <Section>2. COMPROMISSO COM A PRIVACIDADE</Section>
      <P>
        A Hassis Conecta respeita a privacidade dos seus usuários e busca tratar
        os dados pessoais de acordo com a Lei nº 13.709/2018 — Lei Geral de
        Proteção de Dados Pessoais (LGPD), além das demais normas aplicáveis.
      </P>
      <P>
        Esta Política tem como objetivo explicar, de forma transparente, quais
        informações podem ser coletadas, para quais finalidades podem ser
        utilizadas, em quais situações podem ser compartilhadas e quais direitos
        o titular possui.
      </P>

      <Section>3. QUEM PODE UTILIZAR A HASSIS CONECTA</Section>
      <P>
        A plataforma poderá ser utilizada por diferentes categorias de usuários,
        incluindo:
      </P>
      <List
        items={[
          "clientes;",
          "profissionais autônomos;",
          "profissionais cadastrados como MEI;",
          "empresas e prestadores de serviços;",
          "lojistas e estabelecimentos comerciais;",
          "representantes de empresas ou estabelecimentos;",
          "outros usuários autorizados conforme as funcionalidades disponibilizadas pela plataforma.",
        ]}
      />
      <P>
        A Hassis Conecta poderá estabelecer requisitos específicos de cadastro
        para cada categoria.
      </P>
      <P>
        Menores de idade não estão autorizados a realizar cadastro ou utilizar a
        plataforma por conta própria, salvo nas hipóteses legalmente permitidas
        e mediante as condições necessárias.
      </P>

      <Section>4. DADOS PESSOAIS QUE PODEM SER COLETADOS</Section>
      <P>
        Dependendo da forma de utilização da plataforma e da funcionalidade
        acessada, poderão ser coletados dados como:
      </P>

      <SubSection>4.1. Dados de identificação</SubSection>
      <List
        items={[
          "nome;",
          "CPF;",
          "CNPJ;",
          "MEI, quando aplicável;",
          "informações necessárias para identificação do usuário;",
          "foto de perfil;",
          "informações profissionais.",
        ]}
      />

      <SubSection>4.2. Dados de contato</SubSection>
      <List
        items={[
          "número de telefone;",
          "WhatsApp;",
          "endereço de e-mail;",
          "outras informações de contato fornecidas pelo usuário.",
        ]}
      />

      <SubSection>4.3. Dados de localização</SubSection>
      <P>
        Quando necessário para o funcionamento das funcionalidades da
        plataforma, poderão ser tratados dados de localização, incluindo
        localização aproximada ou localização fornecida pelo próprio usuário.
      </P>
      <P>A localização poderá ser utilizada, por exemplo, para:</P>
      <List
        items={[
          "identificar profissionais próximos;",
          "apresentar serviços disponíveis em determinada região;",
          "permitir que o cliente solicite atendimento em determinado local;",
          "facilitar a aproximação entre cliente e profissional;",
          "acompanhar o deslocamento do profissional até o endereço informado pelo cliente, quando a funcionalidade estiver ativa;",
          "apresentar estabelecimentos ou lojistas próximos;",
          "melhorar a experiência de utilização da plataforma;",
          "prevenir fraudes e atividades incompatíveis com o funcionamento do sistema.",
        ]}
      />
      <P>
        Durante um atendimento em andamento, a localização aproximada do
        profissional poderá ser apresentada ao cliente correspondente, pelo
        tempo necessário ao deslocamento, para que ele acompanhe a distância que
        falta até o local combinado.
      </P>
      <P>
        Quando a funcionalidade permitir, o usuário poderá informar ou alterar
        manualmente o local em que deseja solicitar determinado serviço.
      </P>
      <P>
        A utilização da localização observará as configurações disponíveis no
        dispositivo e as permissões necessárias.
      </P>

      <SubSection>4.4. Dados profissionais</SubSection>
      <P>Para profissionais, poderão ser coletadas informações como:</P>
      <List
        items={[
          "profissão;",
          "especialidades;",
          "descrição profissional;",
          "cidade e região de atuação;",
          "endereço, quando necessário;",
          "experiência e informações profissionais fornecidas pelo próprio usuário;",
          "fotos;",
          "vídeos;",
          "informações sobre MEI ou CNPJ;",
          "documentos e informações utilizados em processos de cadastro ou verificação, quando aplicável;",
          "informações necessárias para apresentação do profissional aos clientes.",
        ]}
      />

      <SubSection>4.5. Dados de utilização da plataforma</SubSection>
      <P>
        Poderão ser registrados dados relacionados à utilização da Hassis
        Conecta, como:
      </P>
      <List
        items={[
          "páginas ou funcionalidades acessadas;",
          "solicitações realizadas;",
          "interações dentro da plataforma;",
          "informações técnicas do dispositivo;",
          "endereço IP;",
          "informações de navegador;",
          "sistema operacional;",
          "data e horário de acesso;",
          "registros de segurança;",
          "informações necessárias para diagnóstico de erros e melhoria da plataforma.",
        ]}
      />

      <Section>5. COMO UTILIZAMOS OS DADOS</Section>
      <P>
        Os dados pessoais poderão ser tratados para finalidades relacionadas ao
        funcionamento da Hassis Conecta, incluindo:
      </P>
      <List
        items={[
          "criação e manutenção de contas;",
          "identificação dos usuários;",
          "funcionamento da plataforma;",
          "conexão entre clientes e profissionais;",
          "conexão entre clientes e lojistas;",
          "apresentação de profissionais e estabelecimentos;",
          "envio e recebimento de solicitações;",
          "envio de propostas ou orçamentos;",
          "comunicação entre usuários;",
          "funcionamento de ferramentas de chat, quando disponíveis;",
          "utilização de funcionalidades de localização;",
          "segurança da plataforma;",
          "prevenção e identificação de fraudes;",
          "atendimento ao usuário;",
          "cumprimento de obrigações legais;",
          "exercício regular de direitos;",
          "cumprimento de contratos;",
          "cobrança de valores relacionados aos serviços da plataforma;",
          "processamento de assinaturas ou planos;",
          "melhoria da plataforma;",
          "análise de utilização e desempenho;",
          "envio de comunicações relacionadas à plataforma;",
          "divulgação de informações e novidades, quando permitido;",
          "cumprimento de determinações de autoridades competentes.",
        ]}
      />
      <P>
        O tratamento de dados será realizado de acordo com as bases legais
        aplicáveis previstas na LGPD.
      </P>

      <Section>6. COMPARTILHAMENTO DE DADOS ENTRE USUÁRIOS</Section>
      <P>
        A própria finalidade da Hassis Conecta pode exigir que determinadas
        informações sejam apresentadas a outros usuários.
      </P>
      <P>
        Por exemplo, quando um cliente solicita um profissional, determinadas
        informações fornecidas pelo cliente poderão ser disponibilizadas aos
        profissionais participantes daquela solicitação, na medida necessária
        para que possam avaliar o pedido e entrar em contato.
      </P>
      <P>
        Da mesma forma, informações profissionais cadastradas poderão ser
        apresentadas aos clientes para que estes possam conhecer e avaliar as
        opções disponíveis.
      </P>
      <P>
        As informações exibidas dependerão da funcionalidade utilizada e da
        necessidade para o funcionamento da plataforma.
      </P>
      <P>
        A Hassis Conecta busca limitar o compartilhamento às informações
        necessárias para a finalidade correspondente.
      </P>

      <Section>7. COMPARTILHAMENTO COM PRESTADORES DE TECNOLOGIA</Section>
      <P>
        A Hassis Conecta poderá utilizar fornecedores e prestadores de serviços
        necessários ao funcionamento da plataforma, como serviços de:
      </P>
      <List
        items={[
          "hospedagem;",
          "armazenamento;",
          "infraestrutura tecnológica;",
          "segurança;",
          "autenticação;",
          "envio de mensagens;",
          "processamento de pagamentos, quando aplicável;",
          "análise técnica;",
          "manutenção;",
          "suporte;",
          "ferramentas necessárias ao funcionamento do site ou aplicativo.",
        ]}
      />
      <P>
        Quando houver tratamento de dados por terceiros em nome da Hassis
        Conecta, serão adotadas medidas contratuais e técnicas compatíveis com a
        proteção dos dados pessoais.
      </P>
      <P>
        A Hassis Conecta não comercializa dados pessoais dos usuários como
        produto.
      </P>

      <Section>8. COMPARTILHAMENTO POR EXIGÊNCIA LEGAL</Section>
      <P>Os dados poderão ser compartilhados quando necessário para:</P>
      <List
        items={[
          "cumprimento de obrigação legal ou regulatória;",
          "atendimento de determinação judicial;",
          "atendimento de solicitação de autoridade competente;",
          "exercício regular de direitos;",
          "prevenção de fraudes;",
          "investigação de atividades ilícitas;",
          "proteção da segurança da plataforma, usuários ou terceiros.",
        ]}
      />

      <Section>9. COOKIES E TECNOLOGIAS SEMELHANTES</Section>
      <P>
        A Hassis Conecta poderá utilizar cookies e tecnologias semelhantes para
        permitir o funcionamento adequado da plataforma, melhorar a experiência
        do usuário, compreender a utilização dos serviços, realizar análises
        técnicas e, quando aplicável, personalizar determinadas funcionalidades.
      </P>
      <P>
        Cookies estritamente necessários poderão ser utilizados para o
        funcionamento da plataforma.
      </P>
      <P>
        Para cookies que dependam de consentimento, quando aplicável, serão
        disponibilizadas as opções correspondentes ao usuário.
      </P>
      <P>
        O usuário poderá também administrar determinadas preferências de cookies
        por meio das configurações disponibilizadas na plataforma ou no próprio
        navegador.
      </P>

      <Section>10. SEGURANÇA DAS INFORMAÇÕES</Section>
      <P>
        A Hassis Conecta adota medidas técnicas e administrativas razoáveis e
        compatíveis com sua estrutura e com os riscos envolvidos para proteger
        os dados pessoais contra acessos não autorizados e situações acidentais
        ou ilícitas.
      </P>
      <P>
        Essas medidas podem incluir controles de acesso, autenticação, proteção
        de infraestrutura, monitoramento, procedimentos internos e outras
        medidas de segurança aplicáveis.
      </P>
      <P>
        Entretanto, nenhum sistema conectado à internet pode ser considerado
        absolutamente invulnerável.
      </P>
      <P>
        Por isso, a Hassis Conecta não pode garantir que nenhum incidente de
        segurança ocorrerá, mas buscará adotar medidas preventivas e responder
        adequadamente a eventuais incidentes, observando a legislação aplicável.
      </P>

      <Section>11. RESPONSABILIDADE DO USUÁRIO PELA SEGURANÇA DA CONTA</Section>
      <P>
        O usuário é responsável por manter em segurança suas credenciais de
        acesso e por não compartilhá-las indevidamente.
      </P>
      <P>
        Caso suspeite de acesso não autorizado à sua conta, o usuário deverá
        comunicar a Hassis Conecta pelos canais disponíveis.
      </P>
      <P>
        A Hassis Conecta poderá adotar medidas de segurança, inclusive bloqueio
        ou suspensão temporária da conta, quando houver indícios de acesso
        indevido, fraude ou utilização incompatível com os Termos de Uso.
      </P>

      <Section>12. ARMAZENAMENTO E CONSERVAÇÃO DOS DADOS</Section>
      <P>
        Os dados pessoais poderão ser armazenados enquanto forem necessários
        para cumprir as finalidades descritas nesta Política, manter a relação
        com o usuário, cumprir obrigações legais, exercer direitos ou atender
        outras hipóteses legalmente permitidas.
      </P>
      <P>
        Após o encerramento da relação com o usuário, determinados dados poderão
        permanecer armazenados pelo período necessário para cumprimento de
        obrigações legais, regulatórias, prevenção de fraudes, exercício regular
        de direitos ou outras finalidades permitidas pela legislação.
      </P>
      <P>
        Quando não houver mais necessidade legal ou legítima de conservação, os
        dados poderão ser eliminados, anonimizados ou submetidos a tratamento
        compatível com a legislação.
      </P>

      <Section>13. DIREITOS DOS TITULARES</Section>
      <P>
        Nos termos da LGPD e observadas as hipóteses legais aplicáveis, o
        titular poderá solicitar, entre outros:
      </P>
      <List
        items={[
          "confirmação da existência de tratamento;",
          "acesso aos dados;",
          "correção de dados incompletos, inexatos ou desatualizados;",
          "anonimização, bloqueio ou eliminação de dados desnecessários, excessivos ou tratados em desconformidade;",
          "portabilidade, quando aplicável e regulamentada;",
          "informações sobre compartilhamento;",
          "informações sobre as consequências de eventual negativa de consentimento;",
          "revogação do consentimento, quando o tratamento estiver baseado nessa hipótese;",
          "revisão de decisões tomadas unicamente com base em tratamento automatizado, quando aplicável;",
          "demais direitos previstos na legislação.",
        ]}
      />
      <P>
        O exercício de determinado direito poderá estar sujeito às limitações
        previstas na própria LGPD e em outras normas aplicáveis.
      </P>

      <Section>14. COMO SOLICITAR SEUS DIREITOS</Section>
      <P>
        Solicitações relacionadas aos dados pessoais deverão ser encaminhadas
        para: <Email />
      </P>
      <P>
        A Hassis Conecta poderá solicitar informações adicionais para confirmar
        a identidade do solicitante e evitar que dados pessoais sejam fornecidos
        indevidamente a terceiros.
      </P>
      <P>
        As solicitações serão analisadas e respondidas dentro dos prazos e
        condições previstos na legislação aplicável.
      </P>

      <Section>15. DADOS DE PROFISSIONAIS E VERIFICAÇÃO</Section>
      <P>
        Quando a Hassis Conecta realizar procedimentos de conferência ou
        verificação de informações de profissionais, essa atividade terá
        finalidade de segurança e confiabilidade da plataforma.
      </P>
      <P>
        A realização de uma verificação não significa que a Hassis Conecta
        esteja garantindo:
      </P>
      <List
        items={[
          "idoneidade absoluta do profissional;",
          "ausência de antecedentes;",
          "inexistência de condutas futuras inadequadas;",
          "qualidade absoluta do serviço;",
          "cumprimento de prazo;",
          "cumprimento do orçamento;",
          "existência de seguro;",
          "capacidade técnica absoluta;",
          "ausência de qualquer risco.",
        ]}
      />
      <P>
        O usuário deverá utilizar seu próprio julgamento antes de contratar
        qualquer profissional.
      </P>

      <Section>16. COMUNICAÇÕES</Section>
      <P>
        A Hassis Conecta poderá enviar comunicações necessárias ao funcionamento
        da conta, incluindo informações sobre:
      </P>
      <List
        items={[
          "cadastro;",
          "solicitações;",
          "propostas;",
          "mensagens;",
          "segurança;",
          "alterações de serviço;",
          "pagamentos;",
          "assinaturas;",
          "atualizações importantes.",
        ]}
      />
      <P>
        Também poderão ser enviadas comunicações informativas ou promocionais
        quando houver base legal adequada.
      </P>
      <P>
        Quando aplicável, o usuário poderá solicitar o cancelamento de
        comunicações promocionais.
      </P>

      <Section>17. TRANSFERÊNCIA INTERNACIONAL DE DADOS</Section>
      <P>
        Determinados fornecedores de tecnologia utilizados pela plataforma
        poderão estar localizados no Brasil ou em outros países.
      </P>
      <P>
        Quando houver transferência internacional de dados pessoais, a Hassis
        Conecta adotará as medidas e mecanismos exigidos pela legislação
        aplicável, especialmente pelas regras da LGPD e regulamentações da ANPD.
      </P>

      <Section>18. INCIDENTES DE SEGURANÇA</Section>
      <P>
        Caso ocorra incidente de segurança envolvendo dados pessoais e sejam
        preenchidos os requisitos legais para comunicação, a Hassis Conecta
        adotará as providências cabíveis perante os titulares e/ou a Autoridade
        Nacional de Proteção de Dados, conforme a legislação aplicável.
      </P>

      <Section>19. ALTERAÇÕES DESTA POLÍTICA</Section>
      <P>
        Esta Política de Privacidade poderá ser atualizada periodicamente para
        refletir mudanças na plataforma, nas funcionalidades, nas práticas de
        tratamento de dados, na legislação ou nas orientações das autoridades
        competentes.
      </P>
      <P>A versão atualizada será disponibilizada nesta página.</P>
      <P>
        Quando uma alteração exigir comunicação específica ou novo
        consentimento, a Hassis Conecta adotará as medidas necessárias.
      </P>

      <Section>20. CONTATO</Section>
      <P>
        Para dúvidas, solicitações ou assuntos relacionados à privacidade:
      </P>
      <P>
        <strong>Hassis Conecta</strong>
        <br />
        <strong>E-mail:</strong> <Email />
        <br />
        <strong>Telefone:</strong> 14 99166-5023
      </P>

      <Divider />
      <DocumentTitle>TERMOS DE USO DA HASSIS CONECTA</DocumentTitle>

      <Section>1. ACEITAÇÃO DOS TERMOS</Section>
      <P>
        Ao criar uma conta, acessar ou utilizar a Hassis Conecta, o usuário
        declara que leu, compreendeu e concorda com estes Termos de Uso e com a
        Política de Privacidade.
      </P>
      <P>
        Caso não concorde com qualquer disposição destes Termos, o usuário
        deverá deixar de utilizar a plataforma.
      </P>

      <Section>2. O QUE É A HASSIS CONECTA</Section>
      <P>
        A Hassis Conecta é uma plataforma tecnológica que facilita a conexão
        entre clientes, profissionais e lojistas.
      </P>
      <P>A plataforma pode permitir que o usuário:</P>
      <List
        items={[
          "encontre profissionais;",
          "solicite serviços;",
          "receba propostas;",
          "compare profissionais;",
          "entre em contato com profissionais;",
          "consulte produtos ou preços;",
          "encontre lojas e lojistas;",
          "solicite orçamentos;",
          "utilize ferramentas de localização;",
          "utilize comunicação por chat ou outros canais disponibilizados;",
          "utilize outras funcionalidades disponibilizadas pela plataforma.",
        ]}
      />
      <P>
        A Hassis Conecta não é, por regra, a responsável pela execução dos
        serviços realizados pelos profissionais cadastrados.
      </P>

      <Section>3. INDEPENDÊNCIA DOS PROFISSIONAIS</Section>
      <P>
        Os profissionais cadastrados na Hassis Conecta atuam de forma
        independente.
      </P>
      <P>
        O profissional não é empregado, funcionário, representante, sócio,
        preposto ou subordinado da Hassis Conecta.
      </P>
      <P>
        O cadastro ou apresentação de um profissional na plataforma não cria
        vínculo empregatício, societário, de representação comercial, franquia
        ou qualquer outra relação de subordinação entre o profissional e a
        Hassis Conecta.
      </P>
      <P>
        O profissional é responsável por sua própria atividade, inclusive quanto
        às obrigações:
      </P>
      <List
        items={[
          "civis;",
          "comerciais;",
          "tributárias;",
          "trabalhistas;",
          "previdenciárias;",
          "administrativas;",
          "profissionais;",
          "regulatórias.",
        ]}
      />

      <Section>4. CONTRATAÇÃO DO SERVIÇO</Section>
      <P>
        A contratação realizada entre cliente e profissional é feita diretamente
        entre as partes.
      </P>
      <P>Cabe ao cliente e ao profissional negociar e definir diretamente:</P>
      <List
        items={[
          "preço;",
          "forma de pagamento;",
          "prazo;",
          "materiais;",
          "condições de execução;",
          "escopo do serviço;",
          "garantias;",
          "alterações no serviço;",
          "cancelamento;",
          "demais condições da contratação.",
        ]}
      />
      <P>
        A Hassis Conecta poderá disponibilizar ferramentas para facilitar essa
        comunicação, mas isso não significa que a plataforma seja parte do
        contrato de prestação de serviços celebrado entre cliente e
        profissional.
      </P>

      <Section>5. RESPONSABILIDADE DO PROFISSIONAL</Section>
      <P>O profissional é exclusivamente responsável pelos serviços que prestar.</P>
      <P>Isso inclui, entre outros:</P>
      <List
        items={[
          "qualidade do serviço;",
          "execução;",
          "prazo;",
          "preço informado;",
          "materiais utilizados;",
          "danos causados;",
          "informações fornecidas;",
          "comportamento durante o atendimento;",
          "obrigações legais;",
          "tributos;",
          "documentação profissional;",
          "cumprimento das normas aplicáveis à sua atividade.",
        ]}
      />
      <P>
        Caso um profissional pratique qualquer ato ilícito, irregular ou
        criminoso, sua responsabilidade será própria e pessoal, conforme a
        legislação aplicável.
      </P>

      <Section>6. RESPONSABILIDADE DO CLIENTE</Section>
      <P>
        O cliente é responsável pelas informações que fornecer e pelas decisões
        tomadas durante a contratação.
      </P>
      <P>O cliente deverá:</P>
      <List
        items={[
          "fornecer informações verdadeiras;",
          "descrever corretamente o serviço desejado;",
          "informar adequadamente o local da prestação;",
          "avaliar os profissionais antes da contratação;",
          "negociar diretamente as condições do serviço;",
          "verificar preços, prazos e materiais;",
          "adotar cuidados razoáveis de segurança;",
          "não fornecer desnecessariamente senhas, cartões, documentos ou informações confidenciais ao profissional;",
          "comunicar problemas relevantes pelos canais disponíveis.",
        ]}
      />
      <P>
        A decisão de permitir que um profissional entre em residência, empresa
        ou qualquer outro local é exclusivamente do cliente.
      </P>

      <Section>7. RESPONSABILIDADE DA HASSIS CONECTA</Section>
      <P>
        A Hassis Conecta fornece tecnologia para facilitar a aproximação entre
        usuários.
      </P>
      <P>A plataforma não garante que:</P>
      <List
        items={[
          "determinado profissional aceitará uma solicitação;",
          "haverá profissional disponível em determinada região;",
          "o profissional chegará no horário combinado;",
          "o serviço será executado conforme o esperado;",
          "o orçamento será mantido;",
          "o profissional não cancelará;",
          "o cliente pagará corretamente;",
          "o serviço não apresentará defeitos;",
          "não ocorrerão conflitos entre as partes;",
          "nenhum usuário praticará ato ilícito.",
        ]}
      />
      <P>
        A Hassis Conecta não possui controle permanente sobre a conduta de todos
        os usuários fora da plataforma.
      </P>
      <P>
        A Hassis Conecta também não executa, por si só, os serviços contratados
        diretamente entre cliente e profissional.
      </P>
      <P>
        Isso não exclui eventuais responsabilidades que possam decorrer da
        própria atuação da Hassis Conecta quando houver participação, falha,
        omissão ou responsabilidade legalmente imputável à plataforma.
      </P>

      <Section>8. ATOS ILÍCITOS PRATICADOS POR USUÁRIOS</Section>
      <P>A Hassis Conecta não autoriza qualquer prática ilícita por usuários.</P>
      <P>
        É expressamente proibida a utilização da plataforma para práticas como:
      </P>
      <List
        items={[
          "fraude;",
          "golpe;",
          "furto;",
          "roubo;",
          "ameaça;",
          "violência;",
          "assédio;",
          "discriminação;",
          "perseguição;",
          "dano ao patrimônio;",
          "falsidade;",
          "uso indevido de documentos;",
          "obtenção ilícita de vantagem;",
          "invasão de contas;",
          "qualquer outra conduta contrária à legislação.",
        ]}
      />
      <P>
        Quando tomar conhecimento de situações dessa natureza, a Hassis Conecta
        poderá adotar medidas compatíveis com o caso, incluindo suspensão ou
        encerramento da conta e, quando legalmente cabível, colaboração com
        autoridades competentes.
      </P>

      <Section>9. CADASTRO</Section>
      <P>
        Para utilizar determinadas funcionalidades, o usuário deverá realizar
        cadastro.
      </P>
      <P>
        O usuário declara que fornecerá informações verdadeiras, completas e
        atualizadas.
      </P>
      <P>É proibido:</P>
      <List
        items={[
          "criar conta em nome de outra pessoa sem autorização;",
          "utilizar documentos de terceiros;",
          "fornecer informações falsas;",
          "criar múltiplas contas com finalidade fraudulenta;",
          "compartilhar indevidamente credenciais;",
          "utilizar a conta para finalidade ilícita.",
        ]}
      />
      <P>
        A Hassis Conecta poderá solicitar informações adicionais para segurança,
        validação ou funcionamento da plataforma.
      </P>

      <Section>10. PROFISSIONAL PREMIUM</Section>
      <P>
        A Hassis Conecta poderá disponibilizar planos pagos destinados aos
        profissionais, incluindo modalidade PREMIUM.
      </P>
      <P>
        Os benefícios, valores, condições e funcionalidades de cada plano serão
        aqueles apresentados no momento da contratação.
      </P>
      <P>
        Quando disponibilizado, o plano PREMIUM poderá oferecer benefícios como
        maior exposição, prioridade ou participação em determinadas
        funcionalidades da plataforma.
      </P>
      <P>
        A contratação do plano PREMIUM não garante quantidade mínima de
        clientes, serviços, propostas ou faturamento.
      </P>

      <Section>11. TAXAS DA PLATAFORMA</Section>
      <P>
        A Hassis Conecta poderá cobrar valores pela utilização de determinadas
        funcionalidades.
      </P>
      <P>
        Quando aplicável, os valores serão apresentados ao usuário antes da
        confirmação da operação.
      </P>
      <P>
        Atualmente, determinadas funcionalidades poderão envolver uma taxa de
        utilização cujo valor cheio é de <strong>R$ 9,90</strong> e que, em
        condição promocional de lançamento, é cobrada por{" "}
        <strong>R$ 4,95</strong>, quando expressamente indicada na plataforma.
      </P>
      <P>
        O valor efetivamente devido é sempre aquele exibido na tela de pagamento
        antes da confirmação da operação.
      </P>
      <P>
        A taxa cobrada pela plataforma não se confunde com o valor cobrado pelo
        profissional pelo serviço prestado.
      </P>
      <P>
        O preço do serviço contratado pelo cliente será negociado diretamente
        com o profissional, salvo quando uma funcionalidade específica informar
        de forma diferente.
      </P>

      <Section>12. CONDIÇÕES PROMOCIONAIS DA TAXA</Section>
      <P>
        A taxa de utilização poderá ser objeto de campanhas promocionais, como o
        preço de lançamento indicado no item anterior, em que a Hassis Conecta
        assume parte do valor e o usuário paga somente a diferença.
      </P>
      <P>
        Nessa modalidade não há devolução, reembolso ou estorno posterior: o
        desconto já está aplicado no valor cobrado, de forma que o usuário paga
        apenas o valor promocional exibido no momento do pagamento.
      </P>
      <P>
        Caso a plataforma venha a disponibilizar campanha ou funcionalidade que
        preveja devolução parcial de determinada taxa, as condições
        correspondentes serão apresentadas ao usuário no momento da utilização.
      </P>
      <P>
        A condição promocional não representa desconto sobre o preço do serviço
        do profissional e não altera a responsabilidade das partes pela
        contratação.
      </P>
      <P>
        As condições promocionais poderão ser alteradas ou encerradas a qualquer
        momento, conforme as regras informadas na própria campanha, passando a
        ser exibido o valor então vigente.
      </P>

      <Section>13. PAGAMENTOS ENTRE CLIENTE E PROFISSIONAL</Section>
      <P>
        Salvo quando a própria plataforma disponibilizar sistema específico de
        pagamento, os valores referentes à prestação do serviço serão negociados
        e pagos diretamente entre cliente e profissional.
      </P>
      <P>
        A Hassis Conecta não define o preço do serviço do profissional e não
        garante a realização do pagamento entre as partes.
      </P>

      <Section>14. PROFISSIONAL JÁ</Section>
      <P>
        A Hassis Conecta poderá disponibilizar funcionalidades como o{" "}
        <strong>PROFISSIONAL JÁ</strong>, destinadas a facilitar a localização e
        conexão entre clientes e profissionais disponíveis.
      </P>
      <P>
        A disponibilidade de profissionais dependerá da quantidade de
        profissionais cadastrados, localização, disponibilidade, conexão,
        aceitação das solicitações e demais fatores técnicos e operacionais.
      </P>
      <P>
        A Hassis Conecta não garante que sempre haverá profissionais disponíveis
        para determinada solicitação.
      </P>

      <Section>15. LOCALIZAÇÃO</Section>
      <P>
        Algumas funcionalidades poderão utilizar localização aproximada ou
        informações fornecidas pelo usuário para identificar profissionais e
        serviços próximos.
      </P>
      <P>
        Durante um atendimento em andamento, a plataforma poderá apresentar ao
        cliente a posição aproximada do profissional e a distância restante até
        o local combinado, enquanto durar o deslocamento.
      </P>
      <P>
        A localização poderá ser alterada pelo usuário quando a funcionalidade
        permitir, inclusive para solicitar atendimento em outra região.
      </P>
      <P>A utilização de localização estará sujeita à Política de Privacidade.</P>

      <Section>16. CHAT E COMUNICAÇÃO</Section>
      <P>
        Quando disponibilizado, o sistema de comunicação poderá permitir contato
        entre clientes, profissionais e/ou lojistas.
      </P>
      <P>
        Os usuários são responsáveis pelo conteúdo das mensagens que enviarem.
      </P>
      <P>É proibido utilizar os canais de comunicação para:</P>
      <List
        items={[
          "ameaças;",
          "assédio;",
          "spam;",
          "golpes;",
          "conteúdo ilegal;",
          "conteúdo discriminatório;",
          "divulgação maliciosa;",
          "obtenção indevida de dados;",
          "qualquer finalidade contrária aos presentes Termos.",
        ]}
      />
      <P>
        A Hassis Conecta poderá adotar medidas de segurança e moderação quando
        necessário e permitido pela legislação.
      </P>

      <Section>17. LOJISTAS</Section>
      <P>A Hassis Conecta poderá permitir o cadastro de lojas e lojistas.</P>
      <P>
        Os lojistas poderão apresentar produtos, informações, preços, condições
        comerciais e outras informações.
      </P>
      <P>
        As informações comerciais são de responsabilidade do respectivo lojista.
      </P>
      <P>
        A disponibilidade, preço, estoque, prazo de entrega, qualidade e
        condições de venda deverão ser confirmados diretamente com o
        estabelecimento, salvo quando a plataforma informar expressamente que
        determinada transação será realizada pela própria Hassis Conecta.
      </P>

      <Section>18. ORÇAMENTOS E PREÇOS</Section>
      <P>
        Informações de preços e orçamentos apresentadas na plataforma podem
        sofrer alterações.
      </P>
      <P>
        Um orçamento ou preço apresentado por profissional ou lojista não será
        considerado obrigação da Hassis Conecta, salvo quando a própria
        plataforma for expressamente parte da operação.
      </P>
      <P>
        O usuário deverá confirmar as condições comerciais antes de concluir uma
        contratação.
      </P>

      <Section>19. AVALIAÇÕES E CONTEÚDO</Section>
      <P>
        Quando a plataforma disponibilizar avaliações, comentários, fotos,
        vídeos ou outros conteúdos, o usuário deverá publicar somente
        informações verdadeiras e lícitas.
      </P>
      <P>É proibido publicar:</P>
      <List
        items={[
          "conteúdo falso ou fraudulento;",
          "ofensas;",
          "ameaças;",
          "informações pessoais de terceiros sem autorização;",
          "conteúdo discriminatório;",
          "conteúdo ilegal;",
          "material que viole direitos autorais;",
          "conteúdo destinado a prejudicar deliberadamente outro usuário.",
        ]}
      />
      <P>
        A Hassis Conecta poderá remover ou restringir conteúdos que violem estes
        Termos ou a legislação aplicável.
      </P>

      <Section>20. PROPRIEDADE INTELECTUAL</Section>
      <P>
        A marca Hassis Conecta, seu nome, identidade visual, logotipos, layouts,
        textos, sistemas, funcionalidades, elementos gráficos e demais conteúdos
        próprios da plataforma são protegidos pela legislação aplicável.
      </P>
      <P>
        O usuário não poderá copiar, reproduzir, modificar, distribuir,
        comercializar ou explorar esses elementos sem autorização.
      </P>
      <P>
        O cadastro ou utilização da plataforma não concede ao usuário qualquer
        direito de propriedade sobre a tecnologia ou marca Hassis Conecta.
      </P>

      <Section>21. SUSPENSÃO E ENCERRAMENTO DE CONTAS</Section>
      <P>
        A Hassis Conecta poderá suspender, restringir ou encerrar uma conta
        quando houver, entre outras situações:
      </P>
      <List
        items={[
          "violação destes Termos;",
          "informações falsas;",
          "fraude;",
          "uso indevido da plataforma;",
          "comportamento abusivo;",
          "risco à segurança;",
          "prática ilícita;",
          "tentativa de prejudicar a plataforma ou outros usuários;",
          "descumprimento de obrigações financeiras;",
          "utilização incompatível com a finalidade da plataforma.",
        ]}
      />
      <P>
        Quando aplicável, a Hassis Conecta poderá solicitar esclarecimentos
        antes de adotar medidas definitivas.
      </P>

      <Section>22. DISPONIBILIDADE DA PLATAFORMA</Section>
      <P>
        A Hassis Conecta busca manter seus serviços disponíveis, mas poderá
        ocorrer indisponibilidade temporária em razão de:
      </P>
      <List
        items={[
          "manutenção;",
          "atualizações;",
          "falhas técnicas;",
          "problemas de internet;",
          "problemas de fornecedores;",
          "eventos externos;",
          "ataques cibernéticos;",
          "força maior;",
          "outras situações fora do controle razoável da plataforma.",
        ]}
      />
      <P>
        A indisponibilidade temporária não significa necessariamente falha
        definitiva na prestação do serviço.
      </P>

      <Section>23. LINKS E SERVIÇOS DE TERCEIROS</Section>
      <P>
        A plataforma poderá utilizar ou disponibilizar links, serviços, sistemas
        ou integrações de terceiros.
      </P>
      <P>
        Quando o usuário acessar serviço externo, poderá estar sujeito aos termos
        e políticas daquele terceiro.
      </P>
      <P>
        A Hassis Conecta não controla necessariamente as práticas de terceiros
        independentes.
      </P>

      <Section>24. PROTEÇÃO DE DADOS</Section>
      <P>
        O tratamento de dados pessoais realizado pela Hassis Conecta está
        descrito na Política de Privacidade apresentada nesta mesma página.
      </P>
      <P>
        Ao utilizar a plataforma, o usuário declara estar ciente da existência
        da Política de Privacidade e das regras nela estabelecidas.
      </P>

      <Section>25. ALTERAÇÕES DOS TERMOS</Section>
      <P>
        Estes Termos poderão ser atualizados para refletir alterações na
        plataforma, criação de novas funcionalidades, mudanças comerciais,
        legislação ou necessidades de segurança.
      </P>
      <P>A versão atualizada será disponibilizada na plataforma.</P>
      <P>
        Quando uma alteração exigir manifestação específica do usuário, serão
        adotadas as medidas necessárias.
      </P>

      <Section>26. COMUNICAÇÕES OFICIAIS</Section>
      <P>
        As comunicações relacionadas à conta poderão ser realizadas por e-mail,
        telefone, WhatsApp, notificações dentro da plataforma ou outros canais
        disponibilizados pelo usuário.
      </P>
      <P>
        É responsabilidade do usuário manter seus dados de contato atualizados.
      </P>

      <Section>27. LEGISLAÇÃO APLICÁVEL</Section>
      <P>
        Estes Termos serão interpretados de acordo com as leis da República
        Federativa do Brasil, observadas as normas obrigatórias aplicáveis à
        relação existente entre as partes.
      </P>
      <P>
        Nada nestes Termos pretende afastar direitos que sejam garantidos ao
        consumidor ou ao usuário por normas de caráter obrigatório.
      </P>

      <Section>28. DISPOSIÇÕES FINAIS</Section>
      <P>
        Caso alguma disposição destes Termos seja considerada inválida, ilegal
        ou inexigível, as demais disposições permanecerão válidas naquilo que
        for juridicamente possível.
      </P>
      <P>
        A eventual tolerância da Hassis Conecta em relação ao descumprimento de
        determinada regra não significa renúncia ao direito de exigir seu
        cumprimento posteriormente.
      </P>
      <P>
        Estes Termos devem ser interpretados em conjunto com a Política de
        Privacidade e demais regras específicas eventualmente apresentadas para
        determinadas funcionalidades.
      </P>

      <Section>29. CONTATO</Section>
      <P>
        Para dúvidas, solicitações ou informações relacionadas à plataforma:
      </P>
      <P>
        <strong>Hassis Conecta</strong>
        <br />
        <strong>E-mail:</strong> <Email />
        <br />
        <strong>Telefone:</strong> 14 99166-5023
      </P>

      <Divider />
      <DocumentTitle>DECLARAÇÃO DE CIÊNCIA</DocumentTitle>
      <P>
        Ao criar uma conta ou utilizar as funcionalidades da Hassis Conecta, o
        usuário declara que teve acesso à <strong>Política de Privacidade</strong>{" "}
        e aos <strong>Termos de Uso</strong> e que compreendeu as regras
        aplicáveis à utilização da plataforma.
      </P>
      <p className="mt-8 text-center font-semibold text-gray-900">
        HASSIS CONECTA
        <br />
        <span className="font-normal text-gray-500">
          Conectando você a quem resolve.
        </span>
      </p>
    </div>
  );
}

export default PrivacyPolicyContent;
