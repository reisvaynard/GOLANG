<!DOCTYPE html>
<html lang="en-US">
    <head>
        <title>Assignment GOLANG 3</title>
        <meta charset="UTF-8" />
    </head>
    <body>
        <h2>Edit Name and City</h2>  
    <form method="POST" action="updateadmin">
      <input type="hidden" name="uid" value="{{ .Id }}" />
      <label> User Name </label><input type="text" name="username" value="{{ .Username }}"  /><br />
      <label> Password </label><input type="password" name="password" value="{{ .Password }}"  /><br />
      <label> Name </label><input type="text" name="name" value="{{ .Name }}"  /><br />
      <label> Email </label><input type="text" name="email" value="{{ .Email }}"  /><br />
      <label> Priviledge </label><input type="text" name="priviledge" value="{{ .Priviledge }}"  /><br />
      <input type="submit" value="Save user" />
    </form><br />
    </body>
</html>