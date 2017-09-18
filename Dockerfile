# this is a simple app to test a webap
FROM ubuntu
ADD simplehttp /simplehttp
EXPOSE 8080
ENTRYPOINT ["/simplehttp"]