function Logout(){
        location.href = "/Logout"
}

function GetUser()
{
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.open( "GET", "http://localhost:8000/GetUser/", false );
    xmlHttp.send( null );
    return xmlHttp.responseText;
}