# olric-cloud-plugin

Discover nodes in cloud environments. Uses [hashicorp/go-discover](https://github.com/hashicorp/go-discover) under the hood.

## Install

Get the code:

```
go get -u github.com/buraksezer/olric-cloud-plugin
```

## Usage

### Load as compiled plugin

#### Build

With a properly configured Go environment:

```
go build -buildmode=plugin -o olric-cloud-plugin.so 
```

If you want to strip debug symbols from the produced binary add `-ldflags="-s -w"` to `build` command.

#### Configuration

This plugin is only tested against Kubernetes.

##### Kubernetes

For the uses of `olricd`, the following block is going to loads and configures the plugin.

```
serviceDiscovery:
  provider: "k8s"
  path: "/path/to/olric-cloud-plugin.so"
  args: 'label_selector="run = olricd"'
```

Available configuration parameters for Kubernetes integration:

| Key | Description |
|-----|-------------|
| kubeconfig | Path to the kubeconfig file. |
| namespace | Namespace to search for pods (defaults to "default"). |
| label_selector | Label selector value to filter pods. |
| field_selector | Field selector value to filter pods. |
| host_network | "true" if pod host IP and ports should be used. |

Sample usage:

```
...
args: 'label_selector="run = olricd"'
...
```
From go-discover's documentation:

The **kubeconfig** file value will be searched in the following locations:
1. Use path from "kubeconfig" option if provided.
2. Use path from KUBECONFIG environment variable.
3. Use default path of $HOME/.kube/config

By default, the Pod IP is used to join. The "host_network" option may be set to use the Host IP. No port is used by default. 
Pods may set an annotation 'hashicorp/consul-auto-join-port' to a named port or an integer value. If the value matches a 
named port, that port will be used to join. Note that if "host_network" is set to true, then only pods that have a HostIP 
available will be selected. If a port annotation exists, then the port must be exposed via a HostPort as well, otherwise 
the pod will be ignored.

If you prefer to use Olric in embedded-member mode, the following code snipped will guide you:

```go
c := olric.New("lan")
...
cd := &discovery.CloudDiscovery{}
labelSelector := fmt.Sprintf("app.kubernetes.io/name=%s,app.kubernetes.io/instance=%s", name, instance)
c.ServiceDiscovery = map[string]interface{}{
  "plugin":   cd,
  "provider": "k8s",
  "args":     fmt.Sprintf("namespace=%s label_selector=\"%s\"", namespace, labelSelector),
}
```

See details about how you deploy your Olric-based application on Kubernetes:

* [olric-kubernetes repository](https://github.com/buraksezer/olric-kubernetes)
* [Olric documentation](https://github.com/buraksezer/olric#kubernetes)

## TODO

Test this plugin on the following cloud providers:

* Google Cloud
* DigitalOcean
* Amazon AWS
* Microsoft Azure

## Contributions

Please don't hesitate to fork the project and send a pull request or just e-mail me to ask questions and share ideas.

## License

The Apache License, Version 2.0 - see LICENSE for more details.
