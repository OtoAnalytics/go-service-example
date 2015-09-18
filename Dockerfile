FROM centurylink/ca-certs
WORKDIR /app
# copy binary into image
COPY go-service-example /app/
ENTRYPOINT ["./go-service-example"]
