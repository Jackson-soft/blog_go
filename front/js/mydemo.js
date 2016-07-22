var URL = "http://127.0.0.1:8080"
function UserLogin() {
    $.ajax({
        url: URL + "/userlogin/login",
        type: "POST",
        data: {
            mobile: "12345678901",
            password: "123456"
        }
    }).done(function(data){
        alert(data);
    })
}
