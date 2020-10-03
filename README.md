# Node-Role-Controller

## Introduction/Synopsis

Due to changes introduced in k8s 1.16, 

*"Nodes are not permitted to assert their own role labels. Node roles are typically used to identify privileged or control plane types of nodes, and allowing nodes to label themselves into that pool allows a compromised node to trivially attract workloads."* 
[Kubelet label migration for NodeRestriction Admission Controller](https://github.com/kubernetes/kubernetes/issues/84912#issuecomment-551362981)

This simple controller overcomes this by assigning a role from existing labels to a node. 

## Usage

```
apiVersion: operators.node.role.controller.io/v1alpha1
kind: NodeRoleController
metadata:
  name: node-role-controller
  namespace: operations
spec:
  roles:
    - name: "ingress"
      label: "node.kubernetes.io/ingress"
      value: ""
    - name: "worker"
      label: "node.kubernetes.io/worker"
      value: ""
    - name: "worker-nat"
      label: "node.kubernetes.io/worker"
      value: "nat"
    - name: "worker-nat"
      label: "node.kubernetes.io/worker-nat"
      value: ""
```

## Deployment 
```
helm3 upgrade node-role-controller ./chart/node-role-controller --install --debug --wait --namespace="operations" 
```

## Todo
- If labels do not exist a predefined role of `worker` should be assigned to a node.
- ~In cluster Deployment Strategy~ 
- Assign multiple Node Roles???
## Research

- https://github.com/weaveworks/eksctl/issues/2197
- https://github.com/kubernetes/kubernetes/issues/84912#issuecomment-551362981
- https://github.com/kubernetes/enhancements/blob/master/keps/sig-auth/0000-20170814-bounding-self-labeling-kubelets.md#proposal

