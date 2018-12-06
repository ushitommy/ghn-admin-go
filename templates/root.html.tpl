<!DOCTYPE html>
<html>
<body>
    <h2>Google Homeをしゃべらせるやつを管理するやつ</h2>
    <h3>メニュー</h3>
    ==============<br>
    <form>
        <input type="button" onclick="location.href='../edit'" value="編集" /> <br>
        <input type="button" onclick="location.href='../create'" value="新規登録" />
    </form>
    ==============<br>
    <br><br>
    <h3>ジョブ一覧</h3>
    ==============<br>
    <table border="1">
        <tr>
            <th>ID</th><th>分</th><th>時</th><th>日</th><th>月</th><th>曜日</th><th>しゃべる内容</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{ .ID }}</td><td>{{ .Hour }}</td><td>{{ .Min }}</td><td>{{ .Date }}</td><td>{{ .Month }}</td><td>{{ .Days }}</td><td>{{ .Text }}</td>
        </tr>
        {{end}}
    </table>
    ==============<br>
    <br><br>
</body>
</html>
