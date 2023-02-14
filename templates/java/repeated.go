package java

const repeatedConstTpl = `{{ renderConstants (.Elem "" "") }}`

const repeatedTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.GetIgnoreEmpty }}
			if ( !{{ accessor . }}.isEmpty() ) {
{{- end -}}
{{- if $r.GetMinItems }}
			io.AGPDev.pgv.RepeatedValidation.minItems("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMinItems }});
{{- end -}}
{{- if $r.GetMaxItems }}
			io.AGPDev.pgv.RepeatedValidation.maxItems("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMaxItems }});
{{- end -}}
{{- if $r.GetUnique }}
			io.AGPDev.pgv.RepeatedValidation.unique("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end }}
			io.AGPDev.pgv.RepeatedValidation.forEach({{ accessor . }}, item -> {
				{{ render (.Elem "item" "") }}
			});
{{- if $r.GetIgnoreEmpty }}
			}
{{- end -}}
`
