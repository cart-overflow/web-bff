docker run -p 8001:8080  \
  -e URL=/spec/openapi.yml \
  -v $(pwd)/api:/usr/share/nginx/html/spec/ swaggerapi/swagger-ui
