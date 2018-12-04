<!DOCTYPE html>
<html>
<body>
    <h3>新規登録</h3>
    ==============<br>
    <b>実行条件：</b>
       <form action="/create" method="POST">
    曜日：
    <input type="checkbox" name="days[]" value="mon" />月
    <input type="checkbox" name="days[]" value="tue" />火
    <input type="checkbox" name="days[]" value="wed" />水
    <input type="checkbox" name="days[]" value="thu" />木
    <input type="checkbox" name="days[]" value="fri" />金
    <input type="checkbox" name="days[]" value="sat" />土
    <input type="checkbox" name="days[]" value="sun" />日<br>
    <input type="hidden" name="date" value="*" />
    <input type="hidden" name="month" value="*" />
    時間：
    <input type="text" size="5" name="hour" pattern="(0[0-9]|1[0-9]|2[0-3])" required />時
    <input type="text" size="5" name="min" pattern="([0-5][0-9])" required />分<br>
    しゃべらせる内容：
    <input type="text" size="120" name="text" required /><br>
    <input type="submit" value="追加" />
    </form>
    ==============<br>
<br><br>
    <h3>ジョブ一覧</h3>
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
    <br><br>
    <form>
    <input type="button" onclick="location.href='../'" value="Topにもどる" />
    </form>
</body>
</html>
