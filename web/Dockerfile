FROM oven/bun:1-alpine AS base

FROM base AS deps
WORKDIR /app
COPY package.json bun.lock* ./
RUN bun install --frozen-lockfile --production

FROM base AS builder
WORKDIR /app
COPY package.json bun.lock* ./
RUN bun install --frozen-lockfile
COPY . .
ENV NEXT_TELEMETRY_DISABLED=1
ENV NODE_ENV=production
RUN bun run build

FROM base AS runner
WORKDIR /app

ENV NEXT_TELEMETRY_DISABLED=1
ENV NODE_ENV=production

RUN addgroup -g 1001 -S nodejs && \
    adduser -S nextjs -u 1001

# Copy the standalone output
COPY --from=builder --chown=nextjs:nodejs /app/.next/standalone ./

# Copy static files to the correct location within standalone
COPY --from=builder --chown=nextjs:nodejs /app/.next/static ./.next/static

# Copy public files to the standalone directory
COPY --from=builder --chown=nextjs:nodejs /app/public ./public

# Copy public and static files into the standalone directory structure
RUN cp -r public .next/standalone/ && cp -r .next/static .next/standalone/.next/

USER nextjs

EXPOSE 3000
ENV PORT=3000

CMD ["bun", "server.js"]
