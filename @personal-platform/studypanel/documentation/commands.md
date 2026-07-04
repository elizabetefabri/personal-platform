# Commands used for StudyPanel frontend

This file lists the commands used during development for installing libraries, running the app, building, testing and deploying. Use as a reference to reproduce the environment and steps.

## Environment (assumptions)

- Node.js >= 18
- npm >= 9
- Angular CLI (optional global)

## Angular project creation

ng new frontend --routing --style=scss --standalone

## Local development server

npm install
npm run start

# or using Angular CLI

ng serve

## Install dependencies added during implementation

# Install basic-ftp (deployment script)

npm install basic-ftp --save-dev

# If you need TypeScript node runner for deploy script

npm install -D ts-node typescript @types/node

# Or use npx: npx ts-node deploy.ts

## Build

# development build

npm run build

# production build

npm run build:prod

## Deploy (FTP)

# Example (configured in package.json):

npm run deploy

# which runs:

# 1. npm run build:prod

# 2. npx ts-node deploy.ts

## Tests

# Unit tests (Vitest or configured runner)

npm run test

# E2E tests (if configured)

npm run e2e

## Linting and formatting

# Lint (if configured)

npm run lint

# Prettier

npx prettier --check "\*_/_.{ts,tsx,scss,html,md,json}"

## Useful Git commands used during implementation

git add -A
git commit -m "commit message"
git push origin main

## Notes

- The deploy script reads credentials from `.env.deployment` which must be created from `.env.deployment.example` and must NOT be committed.
- Use `npx ts-node deploy.ts` to run the TypeScript deploy script without globally installing ts-node.
