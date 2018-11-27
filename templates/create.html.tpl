<!DOCTYPE html>
<html>
<body>
    {{.}}<br><br>
    <h3>新規登録</h3>
    ==============<br>
    ***ここに登録フォームを出したい***<br>
    <b>実行条件：</b>
    <form>
    曜日：
    <input type="checkbox" name="day" value="sun" checked="checked">日曜日
    <input type="checkbox" name="day" value="mon">月曜日
    <input type="checkbox" name="day" value="tue">火曜日
    <input type="checkbox" name="day" value="wed">水曜日
    <input type="checkbox" name="day" value="thu">木曜日
    <input type="checkbox" name="day" value="fri">金曜日
    <input type="checkbox" name="day" value="stu">土曜日<br>
    時間：
    <input type="time" name="time"><br>
    しゃべらせる内容：
    <input type="text" size="120" name="text"><br>
    </form>

    <form>
    <input type="submit" value="追加">
    </form>
    ==============<br>
<br><br>
    <h3>ジョブ一覧</h3>
    ==============<br>
    ここにジョブ一覧を出したい<br>
    [Table]<br>
    |No|実行条件（時間と曜日）|実行内容|<br>
    ==============<br>
    <br><br>
    <form>
    <input type="button" onclick="location.href='../'" value="Topにもどる">
</body>
</html>