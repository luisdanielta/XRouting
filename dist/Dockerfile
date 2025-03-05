FROM taluisdaniel/xrouting:server AS server
FROM taluisdaniel/xrouting:web AS web
FROM golang:1.24 AS builder

COPY --from=server /app .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /xrouting

FROM node:20

WORKDIR /app
COPY --from=builder xrouting .
COPY --from=web /app .

RUN chmod +x xrouting
EXPOSE 8000
EXPOSE 4173

RUN npm i -E
ENV NODE_OPTIONS="--max-old-space-size=4096"
RUN npm run build

RUN apt-get update && apt-get install -y tini

ENV AWS_ACCESS_KEY_ID=""
ENV AWS_SECRET_ACCESS_KEY=""
ENV AWS_REGION=""

CMD ["sh", "-c", "/usr/bin/tini -- sh -c './xrouting & npm run preview'"]