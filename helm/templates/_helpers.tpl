{{/*
Expand the name of the chart.
*/}}
{{- define "phoros.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "phoros.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "phoros.labels" -}}
helm.sh/chart: {{ include "phoros.chart" . }}
{{ include "phoros.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "phoros.selectorLabels" -}}
app.kubernetes.io/name: {{ include "phoros.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

