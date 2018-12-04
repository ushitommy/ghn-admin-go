<!DOCTYPE html>
<html>
<body>
<script type="text/javascript">
    function autoCheck(d, id) {
        var dlist =['mon', 'tue', 'wed', 'thu', 'fri', 'sat', 'sun'];
        for ( var i = 0; i < dlist.length; i++) {
            r = d.indexOf(dlist[i]);
            if (r !== -1) {
                document.getElementById(dlist[i] + id).checked = true;
            } else {
                document.getElementById(dlist[i] + id).checked = false;
            }
        }
    }
</script>
    <h3>編集</h3>
    ==============<br>
    <table border="1">
        <tr>
            <th></th><th>ID</th><th>分</th><th>時</th><th>月</th><th>日</th><th>曜日</th><th>しゃべる内容</th>
        </tr>
        {{range .}}
        <form action="/edit" method="POST">
        <tr>
            <td><input type="submit" value="更新" /></td>
            <td><input type="hidden" name="id" value="{{ .ID }}" />{{ .ID }}</td>
            <td><input type="text" size="5" name="hour" pattern="(0[0-9]|1[0-9]|2[0-3])" value="{{ .Hour }}" required /></td>
            <td><input type="text" size="5" name="min" pattern="([0-5][0-9])" value="{{ .Min }}" required /></td>
            <td><input type="hidden" name="month" value="{{ .Month }}" /> {{ .Month }}</td>
            <td><input type="hidden" name="date" value="{{ .Date }}" />{{ .Date }}</td>
            <td><input type="checkbox" id="mon{{ .ID }}" name="days[]" value="mon" />月
                <input type="checkbox" id="tue{{ .ID }}" name="days[]" value="tue" />火
                <input type="checkbox" id="wed{{ .ID }}" name="days[]" value="wed" />水
                <input type="checkbox" id="thu{{ .ID }}" name="days[]" value="thu" />木
                <input type="checkbox" id="fri{{ .ID }}" name="days[]" value="fri" />金
                <input type="checkbox" id="sat{{ .ID }}" name="days[]" value="sat" />土
                <input type="checkbox" id="sun{{ .ID }}" name="days[]" value="sun" />日</td>
            <td><input type="text" size="120" name="text" value="{{ .Text }}" required /></td>
        </tr>
        <script type="text/javascript">
           autoCheck({{ .Days }}, {{ .ID }});
        </script>
        </form>
        {{end}}
    </table>
    ==============<br>
    <br><br>
    <form>
    <input type="button" onclick="location.href='../'" value="Topにもどる" />
</body>
</html>
