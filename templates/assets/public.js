$(document).on("click", ".submit", function() {
    var Name = $(".name").val();
    var Email = $(".email").val();
    var Password = $(".password").val();
    //validate input here


    $.ajax({
        type: "POST",
        url: "/register",
        // The key needs to match your method's input parameter (case-sensitive).
        data: JSON.stringify({ Markers: markers }),
        contentType: "application/json; charset=utf-8",
        dataType: "json",
        success: function(data) { alert(data); },
        error: function(errMsg) {
            alert(errMsg);
        }
    });

});