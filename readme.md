# Podman Kube Secrets
a simple utility program for the correct generation of podman secrets when
using podman play kube

## Usage
| flag           | description                                                                                                                        |
|----------------|------------------------------------------------------------------------------------------------------------------------------------|
| -h, --help     | show this help message and exit                                                                                                    |
| -n ,--name     |the name of the secret in the medadata of the created kubernetes secret. This is also used as the name for the data element itself. |
| -d, --data     | the secret value                                                                                                                   |
| -b64, --base64 | if specified, the generated result will be encoded in base64                                                                       |

create a new secret using [podman-secret-create](https://docs.podman.io/en/latest/markdown/podman-secret-create.1.html).
note the usage of --base64 to encode the secret in base64:
```bash
podman-kube-secrets --base64 -n my-secret -d GEHEIM | podman secret create my-secret -
```
create a new secret using [podman-kube-play](https://docs.podman.io/en/latest/markdown/podman-kube-play.1.html): 
```bash
python3 podman-kube-secrets.py -n my-secret -d GEHEIM | podman kube play -
```

## ContainerizeIT
to avoid installation of python on the target system, podman-kube-secrets.py is also published as a container image based on the official [python image](https://hub.docker.com/_/python/)

```bash
podman run --rm ghcr.io/mwildt/podman-kube-secrets:main -n my-secret -d GEHEIM | podman kube play -
```





