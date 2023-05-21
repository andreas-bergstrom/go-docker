## What is this?

Simple Go HTTP-server with multi-stage Dockerfile

## Deploy to Railway

1. Install the Railway CLI
   https://docs.railway.app/develop/cli#installation

2. Login to Railway

```bash
railway login --browserless
```

3. Create a project for this source code

```bash
railway init
```

4. Deploy the image

```bash
railway up --detach
```
