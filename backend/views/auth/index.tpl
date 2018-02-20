{{ template "layouts/main.tpl" . }}

{{ define "content" }}

    <div class="auth-wrapper">
        <div class="form-wrapper">
            <form id="admin" action="/auth/admin/sign-in" method="post">
                <h2>Welcome, admin.</h2>
                <h2>Please, sign in.</h2>
                <div class="form-group">
                    <input name="public_key" type="text" class="form-control" placeholder="Enter public key">
                    {{if .errorMessage}}
                        <div class="invalid-feedback to-top">
                            {{.errorMessage}}
                        </div>
                    {{end}}
                    <input name="aes_key" type="text" class="form-control" placeholder="Enter password"
                           maxlength="32" size="32">
                    {{if .errorMessage2}}
                        <div class="invalid-feedback">
                            {{.errorMessage2}}
                        </div>
                    {{end}}
                </div>
                <div class="form-group">
                    <input type="submit" value="Sign in!" class="btn btn-primary" />
                </div>
            </form>
        </div>
    </div>

{{ end }}