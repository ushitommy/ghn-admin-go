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
