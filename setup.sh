set -e

cd backend && go mod tidy && cd ..
cd web && bun install && cd ..