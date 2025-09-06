# AI Hands-on Project

This repository contains a hands-on tutorial for building a simple full-stack web application with a form that saves data to a PostgreSQL database using Next.js 15, Prisma, and related tooling.

## Prerequisites

- Node.js 20 (see `.nvmrc`)
- pnpm

## Getting Started

Install dependencies and run the development server:

```bash
pnpm install
pnpm dev
```

Start the database (PostgreSQL 16):

```bash
docker compose up -d db
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result. The page auto-updates as you edit source files in `src/`.

## API

- `GET /api/health` returns `{ ok: true, now }`

## Learn More

- [Next.js Documentation](https://nextjs.org/docs)
- [Prisma Documentation](https://www.prisma.io/docs)
