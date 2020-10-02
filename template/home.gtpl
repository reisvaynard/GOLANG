<!DOCTYPE html>
<html lang="en-US">
    <head>
        <title>Assignment GOLANG 3</title>
        <meta charset="UTF-8" />
    </head>
    <body>
    <h2> Registered </h2>
    <table border="1">
      <thead>
      <tr>
        <td>ID</td>
        <td>Username</td>
        <td>Name</td>
        <td>Email</td>
        <td>Priviledge</td>
      </tr>
       </thead>
       <tbody>
    {{ range . }}
      <tr>
        <td>{{ .Id }}</td>
        <td> {{ .Username }} </td>
        <td>{{ .Name }} </td>
        <td>{{ .Email }} </td>
        <td>{{ .Priviledge }} </td> 
      </tr>
    {{ end }}
      <tr>
        <a href="/login">Keluar</a>
      </tr>
      <tr>
      <a href="/edit?id={{ .Id }}">Update Data</a>
      </tr>
       </tbody>
    </table>
    </body>
</html>

