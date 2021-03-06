

docker rm -f csi-provisioner

docker run --name csi-provisioner -d -e KUBECONFIG=/root/.kube/config -v /root/.kube:/root/.kube -v /tmp/csi:/csi:rw antmoveh/csi-provisioner:v2.1.0 \
--csi-address=unix:///csi/csi-provisioner.sock --v=5 --timeout=150s --leader-election=true --retry-interval-start=500ms \
--feature-gates=Topology=true --extra-create-metadata=true

docker rm -f csi-resizer

docker run --name csi-resizer -d -v /root/.kube:/root/.kube -v /tmp/csi:/csi:rw antmoveh/csi-resizer:v1.1.0 \
--csi-address=unix:///csi/csi-provisioner.sock --v=5 --timeout=150s --leader-election=true --retry-interval-start=500ms \
--handle-volume-inuse-error=false --kubeconfig=/root/.kube/config