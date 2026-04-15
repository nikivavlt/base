# ── Build stage ────────────────────────────────────────────
FROM node:20 AS builder

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm ci

COPY . .

# Bake env vars into the static build
ARG VITE_API_URL
ARG VITE_AUTH_URL
ENV VITE_API_URL=$VITE_API_URL
ENV VITE_AUTH_URL=$VITE_AUTH_URL

RUN npm run build

# ── Run stage ──────────────────────────────────────────────
FROM nginx:alpine

COPY --from=builder /app/dist /usr/share/nginx/html

# SPA routing — all paths serve index.html
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80