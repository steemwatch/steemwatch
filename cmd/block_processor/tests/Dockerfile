FROM jwilder/dockerize

COPY build/tests_entrypoint /

ENTRYPOINT ["dockerize", "-wait", "tcp://stan:4222"]