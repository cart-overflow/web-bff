PORT=8001

echo "🔨 Swagger will be avaliable on localhost:${PORT}"

docker run -p ${PORT}:8080  \
  -e URL=/spec/openapi.yml \
  -v $(pwd)/api:/usr/share/nginx/html/spec/ swaggerapi/swagger-ui
