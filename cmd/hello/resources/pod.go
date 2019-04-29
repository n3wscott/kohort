package resources

func Resource() string {
	return `
apiVersion: v1
kind: Pod
metadata:
  name: hello
  namespace: default
spec:
  containers:
    - image: github.com/n3wscott/kohort/cmd/hello
      imagePullPolicy: IfNotPresent
      name: hello-embedded
  restartPolicy: Never
`
}
