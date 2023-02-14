package java

const requiredTpl = `{{ $f := .Field }}
	{{- if .Rules.GetRequired }}
		if ({{ hasAccessor . }}) {
			io.AGPDev.pgv.RequiredValidation.required("{{ $f.FullyQualifiedName }}", {{ accessor . }});
		} else {
			io.AGPDev.pgv.RequiredValidation.required("{{ $f.FullyQualifiedName }}", null);
		};
	{{- end -}}
`
