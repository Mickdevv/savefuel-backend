mkdir -p ../certs
cd ../certs

openssl req -x509 -newkey rsa:4096 -nodes -keyout server.key -out server.crt -days 3650

# On your host
sudo chown 999:999 server.key server.crt

chmod 600 server.key
chmod 644 server.crt

