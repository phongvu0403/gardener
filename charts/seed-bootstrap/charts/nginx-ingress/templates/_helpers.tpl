{{- define "nginx-ingress.config.data" -}}
{{- range $k, $v := .Values.config }}
{{ $k }}: {{ quote $v }}
{{- end }}
{{- end -}}

{{- define "nginx-ingress.config.name" -}}
nginx-ingress-controller-{{ include "nginx-ingress.config.data" . | sha256sum | trunc 8 }}
{{- end }}