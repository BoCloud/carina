package e2e

import (
	"bocloud.com/cloudnative/carina/utils/log"
	"encoding/json"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"time"
)

var sample = `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: carina-sample
  namespace: carina
  labels:
    app: sample
spec:
  replicas: 1
  selector:
    matchLabels:
      app: sample
  template:
    metadata:
      labels:
        app: sample
    spec:
      containers:
        - name: web-server
          image: docker.io/library/nginx:latest
`

func normalDeployment() {
	label := "app=sample"
	It("default scheduler", func() {
		log.Info("Waiting for pod running")
		stdout, stderr, err := kubectlWithInput([]byte(sample), "apply", "-f", "-")
		Expect(err).ShouldNot(HaveOccurred(), "stdout=%s, stderr=%s", stdout, stderr)
		Eventually(func() error {
			stdout, stderr, err = kubectl("get", "pods", "-l", label, "-o", "json", "-n", NameSpace)
			if err != nil {
				log.Infof("get pod label %s, error %v", label, err)
				return err
			}
			var pods corev1.PodList
			err = json.Unmarshal(stdout, &pods)
			if err != nil {
				return fmt.Errorf("unmarshal error: stdout=%s", stdout)
			}

			for _, pod := range pods.Items {
				if pod.Name == "" {
					log.Infof("not found pod label %s", label)
					return fmt.Errorf("not found pod label %s", label)
				}
				By("pod scheduler validate")
				Expect(pod.Spec.SchedulerName).Should(Equal("default-scheduler"))

				if pod.Status.Phase != corev1.PodRunning {
					log.Infof("pod %s status %s", pod.Name, pod.Status.Phase)
					return fmt.Errorf("pod %s not running", pod.Name)
				}

				log.Infof("pod %s is running", pod.Name)
			}
			return nil
		}, 5*time.Minute, 10*time.Second).Should(Succeed())
	})
}

func deleteNormalDeployment() {
	It("delete normal scheduler pod", func() {
		deploymentName := "carina-sample"
		stdout, stderr, err := kubectl("delete", "deployment", deploymentName, "-n", NameSpace)
		Expect(err).ShouldNot(HaveOccurred(), "stdout=%s, stderr=%s", stdout, stderr)
		By("Waiting for pod delete")
		Eventually(func() error {
			stdout, stderr, err = kubectl("get", "deployment", deploymentName, "-n", NameSpace)
			if err != nil {
				log.Infof("delete deployment %s success %v", deploymentName, err)
				return err
			}
			return nil
		}).Should(HaveOccurred())
	})
}
