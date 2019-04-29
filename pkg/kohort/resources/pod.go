package resources

import (
	"bytes"
	"log"
	"text/template"
)

type PodSpec struct {
	Name      string
	Namespace string
	Image     string
}

func Pod(ps *PodSpec) string {
	// Define a template.
	const podTemplate = `
apiVersion: v1
kind: Pod
metadata:
  name: {{.Name}}
  namespace: {{.Namespace}}
spec:
  containers:
    - image: {{.Image}}
      imagePullPolicy: IfNotPresent
      name: user-container
  restartPolicy: Never
`

	t := template.Must(template.New("pod_tmpl").Parse(podTemplate))

	buf := new(bytes.Buffer)

	err := t.Execute(buf, ps)
	if err != nil {
		log.Println("executing template:", err)
	}
	return buf.String()
}
