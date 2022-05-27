export DDL_RPC_PORT=443
export DDL_RPC_TLS_CERTIFICATE=$(cat ../certs/cert.crt)
export DDL_RPC_TLS_KEY=$(cat ../certs/cert.key)
DDL_RPC_CLIENT_AUTHENTICATION=NoClientCert
