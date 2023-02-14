package java

const stringConstTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.In }}
		private final {{ javaTypeFor . }}[] {{ constantName . "In" }} = new {{ javaTypeFor . }}[]{
			{{- range $r.In -}}
				"{{- sprintf "%v" . -}}",
			{{- end -}}
		};
{{- end -}}
{{- if $r.NotIn }}
		private final {{ javaTypeFor . }}[] {{ constantName . "NotIn" }} = new {{ javaTypeFor . }}[]{
			{{- range $r.NotIn -}}
				"{{- sprintf "%v" . -}}",
			{{- end -}}
		};
{{- end -}}
{{- if $r.Pattern }}
		com.google.re2j.Pattern {{ constantName . "Pattern" }} = com.google.re2j.Pattern.compile({{ javaStringEscape $r.GetPattern }});
{{- end -}}`

const stringTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.GetIgnoreEmpty }}
			if ( !{{ accessor . }}.isEmpty() ) {
{{- end -}}
{{- if $r.Const }}
			io.AGPDev.pgv.ConstantValidation.constant("{{ $f.FullyQualifiedName }}", {{ accessor . }}, "{{ $r.GetConst }}");
{{- end -}}
{{- if $r.In }}
			io.AGPDev.pgv.CollectiveValidation.in("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "In" }});
{{- end -}}
{{- if $r.NotIn }}
			io.AGPDev.pgv.CollectiveValidation.notIn("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "NotIn" }});
{{- end -}}
{{- if $r.Len }}
			io.AGPDev.pgv.StringValidation.length("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetLen }});
{{- end -}}
{{- if $r.MinLen }}
			io.AGPDev.pgv.StringValidation.minLength("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMinLen }});
{{- end -}}
{{- if $r.MaxLen }}
			io.AGPDev.pgv.StringValidation.maxLength("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMaxLen }});
{{- end -}}
{{- if $r.LenBytes }}
			io.AGPDev.pgv.StringValidation.lenBytes("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetLenBytes }});
{{- end -}}
{{- if $r.MinBytes }}
			io.AGPDev.pgv.StringValidation.minBytes("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMinBytes }});
{{- end -}}
{{- if $r.MaxBytes }}
			io.AGPDev.pgv.StringValidation.maxBytes("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMaxBytes }});
{{- end -}}
{{- if $r.Pattern }}
			io.AGPDev.pgv.StringValidation.pattern("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "Pattern" }});
{{- end -}}
{{- if $r.Prefix }}
			io.AGPDev.pgv.StringValidation.prefix("{{ $f.FullyQualifiedName }}", {{ accessor . }}, "{{ $r.GetPrefix }}");
{{- end -}}
{{- if $r.Contains }}
			io.AGPDev.pgv.StringValidation.contains("{{ $f.FullyQualifiedName }}", {{ accessor . }}, "{{ $r.GetContains }}");
{{- end -}}
{{- if $r.NotContains }}
			io.AGPDev.pgv.StringValidation.notContains("{{ $f.FullyQualifiedName }}", {{ accessor . }}, "{{ $r.GetNotContains }}");
{{- end -}}
{{- if $r.Suffix }}
			io.AGPDev.pgv.StringValidation.suffix("{{ $f.FullyQualifiedName }}", {{ accessor . }}, "{{ $r.GetSuffix }}");
{{- end -}}
{{- if $r.GetEmail }}
			io.AGPDev.pgv.StringValidation.email("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetAddress }}
			io.AGPDev.pgv.StringValidation.address("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetHostname }}
			io.AGPDev.pgv.StringValidation.hostName("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetIp }}
			io.AGPDev.pgv.StringValidation.ip("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetIpv4 }}
			io.AGPDev.pgv.StringValidation.ipv4("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetIpv6 }}
			io.AGPDev.pgv.StringValidation.ipv6("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetUri }}
			io.AGPDev.pgv.StringValidation.uri("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetUriRef }}
			io.AGPDev.pgv.StringValidation.uriRef("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetUuid }}
			io.AGPDev.pgv.StringValidation.uuid("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetIgnoreEmpty }}
			}
{{- end -}}
`
