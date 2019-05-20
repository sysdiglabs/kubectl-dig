# kubectl dig

<table style="width: 100%; border-style: none;"><tr>
<td style="width: 140px; text-align: center;"><img width="128px" src="docs/img/logo.png" alt="kubectl dig logo"/></td>
<td>
<strong>kubectl dig</strong><br />
<i>A is a simple, intuitive, and fully customizable UI to dig into your kubernetes clusters</i><br>

```
kubectl dig <node>
```
</td>
</tr></table>

[![asciicast](https://asciinema.org/a/czA06tSyEjKpusDooKZe3KQry.svg)](https://asciinema.org/a/czA06tSyEjKpusDooKZe3KQry)

## Install


```bash
go get -u github.com/leodido/kubectl-dig/cmd/kubectl-dig
```

## Usage

### Just dig
There's only one thing to do, provide the node name!

```
kubectl dig <node>
```

You just identify the node you want to dig in with `kubectl get nodes` and then
provide it to the dig command!

```
kubectl dig ip-180-12-0-152.ec2.internal
```

### dig + cluster metadata

By default, `kubectl dig` shows only information about the local node, if you want to dig from that node to the whole cluster you have to provide a **service account** that can read resources.

You can create a `dig-viewer` service account with:

```bash
kubectl apply -f https://github.com/leodido/kubectl-dig/raw/develop/docs/setup/read-serviceaccount.yml
```

Then you just use it with `kubectl dig`.

```bash
kubectl dig --serviceaccount dig-viewer 127.0.0.1
```

At this point you have access to the fancy cluster metadata, press `F2` and look for the `K8s` views!


---

[![Analytics](https://ga-beacon.appspot.com/UA-49657176-1/kubectl-dig?flat)](https://github.com/igrigorik/ga-beacon)

