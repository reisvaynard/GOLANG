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
        <td>Edit</td>
        <td>Delete</td>
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
        <td><a href="/editadmin?id={{ .Id }}">Edit</a></td>
        <td><a href="/delete?id={{ .Id }}">Delete</a></td>
      </tr>
    {{ end }}
      <tr>
        <td>
        <a href="/insert">Tambah User</a>
        </td>
        <a href="/login">Keluar</a>
      </tr>
       </tbody>
    </table>
    </body>
</html>
