FROM alpine:latest

RUN apk add --update curl redis && apk add --no-cache bash && rm -rf /var/cache/apk/*

WORKDIR /app

COPY . .

RUN chmod +x create-rule.sh

RUN chmod +x wait-for-it.sh

RUN chmod +x redis-init.sh

CMD ["sh", "-c", "./redis-init.sh ; ./wait-for-it.sh ${KUIPER_IP}:${KUIPER_PORT} --strict -- ./create-rule.sh"]
