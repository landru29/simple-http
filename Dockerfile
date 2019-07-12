FROM debian:stretch

COPY ./simple-http /bin/.

CMD simple-http