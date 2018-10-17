# Riot - A CLI application which introduces chaos in Kubernetes cluster using open-source Chaos Engineering tools

Riot is a small app written in Go which can be used to add chaos to the kubernetes cluster.

## Prerequisites
1. Kubernetes cluster up and running.
2. `kubectl`, if running locally.
3. `Docker`, if running a docker image.

## Installation
1. Clone the repository.
2. Run `make all` for creating a binary.
3. Modify `<your-docker-username>` in `Makefile` and run `make docker` for building a docker image

## Getting Started
1. Clone the repository and build the binary/docker image.
2. Modify the values for chaos experiments for supported tools. Manifests and related configuration can be found in `manifests` directory.
3. Run `riot install chaos-tool`.
4. While using docker image, run: `docker run -it -v /path/to/kube/config.yml:/root/.kube/config ishantanu16/riot-k8s install/remove/update`. On prompting for kubeconfig path, enter `/root/.kube/config`.

## Supported Operations

1. `install`: Initiates chaos in the cluster.
2. `remove`: Removes chaos from the cluster.
3. `update`: Updates the attributes of the ongoing chaos.

## Contributions

Contribtions are welcome for adding more Chaos Engineering tools which can help in adding chaos to kubernetes clusters.

## Credits

To all the creators of awesome chaos engineering tools:
1. [Pumba](https://github.com/alexei-led/pumba) (@alexei-led)
2. [kube-monkey](https://github.com/asobti/kube-monkey) (@asobti)
3. [pod-reaper](https://github.com/target/pod-reaper) (@target)
4. [chaoskube](https://github.com/linki/chaoskube) (@linki)
