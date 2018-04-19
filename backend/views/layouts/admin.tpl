{{ template "layouts/main.tpl" . }}

{{ define "header" }}
      {{ template "partials/header.tpl" }}
{{ end }}

{{ define "content" }}
    <aside class="admin-sidebar">
        {{ template "partials/aside.html" }}
    </aside>
    <article>
        {{ block "layout-content" . }}{{ end }}
    </article>
{{ end }}