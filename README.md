# go local k8s pipeline

Functionality includes:

- walks directory for changes in docker-compose files, Dockerfiles, k8s, and .git (pushes)
- rebuilds docker image on change and reapplies k8s files
