# Imagem de Origem
FROM node:20-alpine
# Diretório de trabalho(é onde a aplicação ficará dentro do container).
WORKDIR /app
# Adicionando `/app/node_modules/.bin` para o $PATH
ENV PATH /app/node_modules/.bin:$PATH
# Instalando dependências da aplicação e armazenando em cache.
COPY . .
RUN npm install
RUN npm install react-scripts@5.0.1 -g
# Inicializa a aplicação
CMD ["npm", "start"]