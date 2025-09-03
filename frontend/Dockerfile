# Build da aplicação - faz tudo, mas retorna erro, foi preferível subir buildado: -complexidade, -manutenção, +praticidade
# //# FROM node:20-alpine AS build

# WORKDIR /app
# COPY package.json package-lock.json* ./
# RUN npm install
# COPY . .
# RUN npm run build

# # Servir com nginx
# FROM nginx:alpine
# COPY --from=build /app/dist /usr/share/nginx/html
# COPY nginx.conf /etc/nginx/conf.d/default.conf

# EXPOSE 80
# CMD ["nginx", "-g", "daemon off;"]
# //

FROM nginx:alpine

# Copiar arquivos já buildados localmente
COPY dist /usr/share/nginx/html

# Copiar configuração do nginx
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]