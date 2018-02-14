<!DOCTYPE html>

<html>
    <head>
        <title>Wizebit</title>
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
        <link rel="stylesheet" href="/static/css/app.css">
        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
        {{ block "css" . }}{{ end }}
    </head>
    <body>
        {{ block "header" . }}{{ end }}
        <main>
            {{ block "content" . }}{{ end }}
        </main>

        {{ block "footer" . }}{{ end }}

        {{/*<script type="text/javascript" src="http://code.jquery.com/jquery-2.0.3.min.js"></script>*/}}
        {{/*<script src="http://netdna.bootstrapcdn.com/bootstrap/3.0.3/js/bootstrap.min.js"></script>*/}}
        {{ block "js" . }}{{ end }}
    </body>
</html>
