# Publr Site

Microservices Publishing Platform.

## Getting Started

### Prerequisites

- [golang](https://golang.org/doc/install)
- [bazel build](https://docs.bazel.build/versions/master/getting-started.html#installation)
- [skaffold](https://skaffold.dev/docs/getting-started/#installing-skaffold)
- [kubernetes cluster](https://kubernetes.io/docs/setup/)

### Quick Install

Clone this project

```bash
git clone https://github.com/prksu/publr.git
```

Deploy database

```bash
kubectl apply -f deployment/kubernetes/publr/database/
```

Deploy services

```bash
skaffold run --default-repo=YOUR_CONTIANER_REGISTRY
```

## License

This project is licensed under the Apache-2.0 License - see the [LICENSE](LICENSE) for details