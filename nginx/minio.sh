docker run -d -p 9000:9000 -p 9001:9001 \
  -e "MINIO_ROOT_USER=yale918" \
  -e "MINIO_ROOT_PASSWORD=12345678" \
  -v minio-data:/data \
  --network test_nginx \
  --name minio \
  minio/minio server /data --console-address ":9001"