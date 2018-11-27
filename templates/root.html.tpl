<!DOCTYPE html>
<html>
<body>
    
    <h3>メニュー</h3>
    <form>
        <input type="button" onclick="location.href='../edit'" value="編集" /> <br>
        <input type="button" onclick="location.href='../create'" value="新規登録" />
    </form>
    <br><br>
    <h3>ジョブ一覧</h3>
    ==============<br>
    <table border="1">
        <tr>
            <th>ID</th><th>分</th><th>時</th><th>日</th><th>月</th><th>曜日</th><th>しゃべる内容</th>
        </tr>
        <tr>
            <td>{{ .ID }}</td><td>{{ .Hour }}</td><td>{{ .Min }}</td><td>{{ .Date }}</td><td>{{ .Month }}</td><td>{{ .Day }}</td><td>{{ .Text }}</td>
        </tr>
    </table>
    ==============<br>
    <br><br>
</body>
</html>