package main

import (
	"bytes"
	"fmt"
	"text/template"

	"k8s.io/client-go/pkg/api"
)

func main() {
	t, _ := template.New("service").Parse(`
    apiVersion: v1
    kind: Service
    metadata:
      name: {{ .Name }}-subdomain
      labels:
        app: {{ .Name }}
    spec:
      selector:
        app: {{ .Name }}
      clusterIP: None
      ports:
      - name: dummy # Actually, no port is needed.
        port: 1234
        targetPort: 1234`)

	buf := new(bytes.Buffer)
	t.Execute(&buf, map[string]string{
		"Name": "hoo",
	})

	decode := api.Codecs.UniversalDeserializer().Decode
	object := decode(buf.Bytes(), nil, nil)
	fmt.Println(object)
}
