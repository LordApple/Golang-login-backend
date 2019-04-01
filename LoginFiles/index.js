function login(){
    var uname = document.getElementById("field1").value;
    var pswd = document.getElementById("field2").value;

    if (pswd.length != 0 && uname.length != 0){
        location.href = document.location + "handleLogin" + "?u=" + uname + "&p=" + pswd 
    }
}