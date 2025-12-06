FROM node:20-alpine AS builder

ARG NUXT_BACKEND_URL
ENV NUXT_BACKEND_URL=${NUXT_BACKEND_URL}

WORKDIR /app

COPY package*.json ./
RUN npm ci

COPY . .
RUN npm run build

FROM node:20-alpine

WORKDIR /app

COPY --from=builder /app/.output ./.output

EXPOSE 3000

CMD ["node", ".output/server/index.mjs"]
