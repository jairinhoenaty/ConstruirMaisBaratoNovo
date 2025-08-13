// Mock data for states and professionals
export const states = [
  { id: "AC", name: "Acre" },
  { id: "AL", name: "Alagoas" },
  { id: "AP", name: "Amapá" },
  { id: "AM", name: "Amazonas" },
  { id: "BA", name: "Bahia" },
  { id: "CE", name: "Ceará" },
  { id: "DF", name: "Distrito Federal" },
  { id: "ES", name: "Espírito Santo" },
  { id: "GO", name: "Goiás" },
  { id: "MA", name: "Maranhão" },
  { id: "MT", name: "Mato Grosso" },
  { id: "MS", name: "Mato Grosso do Sul" },
  { id: "MG", name: "Minas Gerais" },
  { id: "PA", name: "Pará" },
  { id: "PB", name: "Paraíba" },
  { id: "PR", name: "Paraná" },
  { id: "PE", name: "Pernambuco" },
  { id: "PI", name: "Piauí" },
  { id: "RJ", name: "Rio de Janeiro" },
  { id: "RN", name: "Rio Grande do Norte" },
  { id: "RS", name: "Rio Grande do Sul" },
  { id: "RO", name: "Rondônia" },
  { id: "RR", name: "Roraima" },
  { id: "SC", name: "Santa Catarina" },
  { id: "SP", name: "São Paulo" },
  { id: "SE", name: "Sergipe" },
  { id: "TO", name: "Tocantins" },
];

export const carouselSections = [
  { id: "home", label: "Página Inicial", type: "H", route: "/" },
  { id: "marketplace", label: "Marketplace", type: "M", route: "/marketplace" },
  {
    id: "search",
    label: "Busca de Profissionais",
    type: "S",
    route: "/search",
  },
  {
    id: "privacy",
    label: "Política de Privacidade",
    type: "P",
    route: "/privacy",
  },
  //{ id: "dashboard", label: "Dashboard", type: "D", route: "/dashboard" },
  { id: "login", label: "Login", type: "L", route: "/login" },
  { id: "register", label: "Cadastre-se", type: "R", route: "/register" },
  {
    id: "profissoes_cidade",
    label: "Profissões/Cidade",
    type: "C",
    route: "/search-result",
  },
  {
    id: "professional-panel",
    label: "Área de trabalho",
    type: "A",
    route: "/professional-panel",
  },
  {
    id: "regions",
    label: "Regiões",
    type: "U",
    route: "",
  },
  {
    id: "before-search",
    label: "Modal Busca",
    type: "B",
    route: "",
  },
];


export const carouselImages = {
  home: [
    /*{
      url: 'https://images.unsplash.com/photo-1503387762-592deb58ef4e',
      title: 'Arquitetura Moderna',
    },
    {
      url: 'https://images.unsplash.com/photo-1541976590-713941681591',
      title: 'Projetos Residenciais',
    },
    {
      url: 'https://images.unsplash.com/photo-1466096115517-bceecbfb6fde',
      title: 'Construção Civil',
    },
    {
      url: 'https://images.unsplash.com/photo-1581094794329-c8112a89af12',
      title: 'Acabamentos Premium',
    },
    {
      url: 'https://images.unsplash.com/photo-1504307651254-35680f356dfd',
      title: 'Reformas Completas',
    },*/
    {
      url: "/images/Untitled.jpeg",
      title: "Anuncie Aqui",
    },
    {
      url: "/images/banner2.jpeg",
      title: "Se Você viu",
    },
    {
      url: "/images/banner3.jpeg",
      title: "Seus Clientes",
    },
    {
      url: "/images/banner4.jpeg",
      title: "Também vão ver",
    },
  ],
  marketplace: [
    /*  {
        url: 'https://images.unsplash.com/photo-1607472586893-edb57bdc0e39',
        title: 'Materiais de Construção',
      },
      {
        url: 'https://images.unsplash.com/photo-1581235720704-06d3acfcb36f',
        title: 'Ferramentas Profissionais',
      },
      {
        url: 'https://images.unsplash.com/photo-1588854337115-1c67d9247e4d',
        title: 'Acabamentos',
      },
      {
        url: 'https://images.unsplash.com/photo-1584622650111-993a426fbf0a',
        title: 'Tintas e Revestimentos',
      },
      {
        url: 'https://images.unsplash.com/photo-1517646287270-a5a9ca602e5c',
        title: 'Equipamentos',
      },*/
    {
      url: "/images/Untitled.jpeg",
      title: "Anuncie Aqui",
    },
    {
      url: "/images/banner2.jpeg",
      title: "Se Você viu",
    },
    {
      url: "/images/banner3.jpeg",
      title: "Seus Clientes",
    },
    {
      url: "/images/banner4.jpeg",
      title: "Também vão ver",
    },
  ],
  search: [
    /*
      {
        url: 'https://images.unsplash.com/photo-1504917595217-d4dc5ebe6122',
        title: 'Profissionais Qualificados',
      },
      {
        url: 'https://images.unsplash.com/photo-1581092580497-e0d23cbdf1dc',
        title: 'Arquitetos',
      },
      {
        url: 'https://images.unsplash.com/photo-1590274853856-f22d5ee3d228',
        title: 'Engenheiros',
      },
      {
        url: 'https://images.unsplash.com/photo-1626885930974-4b69aa21bbf9',
        title: 'Empreiteiros',
      },
      {
        url: 'https://images.unsplash.com/photo-1527192491265-7e15c55b1ed2',
        title: 'Designers de Interiores',
      },*/
    {
      url: "/images/Untitled.jpeg",
      title: "Anuncie Aqui",
    },
    {
      url: "/images/banner2.jpeg",
      title: "Se Você viu",
    },
    {
      url: "/images/banner3.jpeg",
      title: "Seus Clientes",
    },
    {
      url: "/images/banner4.jpeg",
      title: "Também vão ver",
    },
  ],
  privacy: [
    /*
    {
      url: 'https://images.unsplash.com/photo-1562813733-b31f71025d54',
      title: 'Segurança de Dados',
    },
    {
      url: 'https://images.unsplash.com/photo-1563013544-824ae1b704d3',
      title: 'Privacidade Garantida',
    },
    {
      url: 'https://images.unsplash.com/photo-1451187580459-43490279c0fa',
      title: 'Tecnologia Avançada',
    },
    {
      url: 'https://images.unsplash.com/photo-1504639725590-34d0984388bd',
      title: 'Proteção Digital',
    },
    {
      url: 'https://images.unsplash.com/photo-1563986768494-4dee2763ff3f',
      title: 'Compromisso com Segurança',
    },*/
    {
      url: "/images/Untitled.jpeg",
      title: "Anuncie Aqui",
    },
    {
      url: "/images/banner2.jpeg",
      title: "Se Você viu",
    },
    {
      url: "/images/banner3.jpeg",
      title: "Seus Clientes",
    },
    {
      url: "/images/banner4.jpeg",
      title: "Também vão ver",
    },
  ],
  dashboard: [
    /*
    {
      url: 'https://images.unsplash.com/photo-1460925895917-afdab827c52f',
      title: 'Gestão de Projetos',
    },
    {
      url: 'https://images.unsplash.com/photo-1507925921958-8a62f3d1a50d',
      title: 'Análise de Dados',
    },
    {
      url: 'https://images.unsplash.com/photo-1551288049-bebda4e38f71',
      title: 'Controle Financeiro',
    },
    {
      url: 'https://images.unsplash.com/photo-1454165804606-c3d57bc86b40',
      title: 'Relatórios',
    },
    {
      url: 'https://images.unsplash.com/photo-1531403009284-440f080d1e12',
      title: 'Desempenho',
    },
    */
    {
      url: "/images/Untitled.jpeg",
      title: "Anuncie Aqui",
    },
    {
      url: "/images/banner2.jpeg",
      title: "Se Você viu",
    },
    {
      url: "/images/banner3.jpeg",
      title: "Seus Clientes",
    },
    {
      url: "/images/banner4.jpeg",
      title: "Também vão ver",
    },
  ],
  login: [
    /*
    {
      url: 'https://images.unsplash.com/photo-1497366754035-f200968a6e72',
      title: 'Área do Profissional',
    },
    {
      url: 'https://images.unsplash.com/photo-1497366811353-6870744d04b2',
      title: 'Portal do Cliente',
    },
    {
      url: 'https://images.unsplash.com/photo-1497366216548-37526070297c',
      title: 'Acesso Seguro',
    },
    {
      url: 'https://images.unsplash.com/photo-1497366412874-3415097a27e7',
      title: 'Área Restrita',
    },
    {
      url: 'https://images.unsplash.com/photo-1486406146926-c627a92ad1ab',
      title: 'Cadastro',
    },
    */
    {
      url: "/images/Untitled.jpeg",
      title: "Anuncie Aqui",
    },
    {
      url: "/images/banner2.jpeg",
      title: "Se Você viu",
    },
    {
      url: "/images/banner3.jpeg",
      title: "Seus Clientes",
    },
    {
      url: "/images/banner4.jpeg",
      title: "Também vão ver",
    },
  ],
  profissoes_cidade: [
    /*{
      url: 'https://images.unsplash.com/photo-1497366754035-f200968a6e72',
      title: 'Área do Profissional',
    },
    {
      url: 'https://images.unsplash.com/photo-1497366811353-6870744d04b2',
      title: 'Portal do Cliente',
    },
    {
      url: 'https://images.unsplash.com/photo-1497366216548-37526070297c',
      title: 'Acesso Seguro',
    },
    {
      url: 'https://images.unsplash.com/photo-1497366412874-3415097a27e7',
      title: 'Área Restrita',
    },
    {
      url: 'https://images.unsplash.com/photo-1486406146926-c627a92ad1ab',
      title: 'Cadastro',
    },
    */
    {
      url: "/images/Untitled.jpeg",
      title: "Anuncie Aqui",
    },
    {
      url: "/images/banner2.jpeg",
      title: "Se Você viu",
    },
    {
      url: "/images/banner3.jpeg",
      title: "Seus Clientes",
    },
    {
      url: "/images/banner4.jpeg",
      title: "Também vão ver",
    },
  ],
};
