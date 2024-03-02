#!/bin/bash

# Path to your .env file
ENV_FILE=".env"
# Path to your env-secret.yaml template
SECRET_TEMPLATE="env-secret.yaml.template"
# Path to the final env-secret.yaml
SECRET_YAML="templates/env-secret.yaml"

# Encode .env file contents to base64
ENCODED_ENV=$(base64 -i "$ENV_FILE" | tr -d '\n')

# Replace a placeholder in your env-secret.yaml.template with the encoded .env content
sed "s|REPLACE_WITH_ENV|$ENCODED_ENV|g" "$SECRET_TEMPLATE" > "$SECRET_YAML"

echo "Secret yaml generated at $SECRET_YAML"
