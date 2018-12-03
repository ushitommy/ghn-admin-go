<!DOCTYPE html>
<html>
<body>
    <h3>ジョブ一覧</h3>
    ==============<br>
    <table border="1">
        <tr>
            <th></th><th>ID</th><th>分</th><th>時</th><th>日</th><th>月</th><th>曜日</th><th>しゃべる内容</th>
        </tr>
        {{range .}}
        <tr>
            <td><input type="button" value="edit" /></td><td>{{ .ID }}</td><td>{{ .Hour }}</td><td>{{ .Min }}</td><td>{{ .Date }}</td><td>{{ .Month }}</td><td>{{ .Days }}</td><td>{{ .Text }}</td>
        </tr>
        {{end}}
    </table>
    ==============<br>
    <br><br>
    <form>
    <input type="button" onclick="location.href='../'" value="Topにもどる" />
</body>
</html>
