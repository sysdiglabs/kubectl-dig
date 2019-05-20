# kubectl dig

<table style="width: 100%; border-style: none;"><tr>
<td style="width: 140px; text-align: center;"><a href="https://aka.ms/vscode-remote/download/extension"><img width="128px" src="docs/img/logo.png" alt="kubectl dig logo"/></a></td>
<td>
<strong>kubectl dig</strong><br />
<i>A is a simple, intuitive, and fully customizable UI to dig into your kubernetes clusters</i><br>

```
kubectl dig <node>
```
</td>
</tr></table>


## Install


```bash
go get -u github.com/leodido/kubectl-dig/cmd/kubectl-dig
```

## Usage

There's only one thing to do, provide the node name!

```
kubectl dig <node>
```

You just identify the node you want to dig in with `kubectl get nodes` and then
provide it to the dig command!

```
kubectl dig ip-180-12-0-152.ec2.internal
```

---

[![Analytics](https://ga-beacon.appspot.com/UA-49657176-1/kubectl-dig?flat)](https://github.com/igrigorik/ga-beacon)

