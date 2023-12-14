# ext-builder

Join external builder instance from outside the production (cluster) deployment

## Usage

1. Set `./id_ed25519` (ns default private key for fetching repository) and config `./config.yaml` accordingly
    - Fetch configuration from production if necessary
2. Set secret information in `.env` file according to `compose.yaml`
3. `ssh -L 0.0.0.0:3306:private.kmbk.tokyotech.org:3306 -L 0.0.0.0:10000:10.43.193.98:10000 c1-203`
    - Do this if some ports are unreachable from outside the deployment stack
4. `docker compose pull`
5. `docker compose up -d`
