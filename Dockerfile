FROM centos:latest

COPY rpg-demo /app/rpg

CMD [ "/app/rpg" ]