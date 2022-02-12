$.ajax({
    url: '/admin/list',
    type: "get",
    dataType: "json",

    success: function(data) {
        drawTable(data);
    }
});

function drawTable(data) {
    for (var i = 0; i < data.length; i++) {
        drawRow(data[i]);
    }
}

function drawRow(rowData) {
    var row = $("<tr />")
    var editFunction = "<button type=\"button\" class=\"open-editEvents btn btn-success btn-sm\" data-id=\"" + rowData.ID + "\" data-name=\"" + rowData.Name + "\" data-email=\"" + rowData.Email + "\" data-role=\"" + rowData.Role + "\" data-verify=\"" + rowData.Verify + "\" data-toggle=\"modal\" data-target=\"#editModal\">Edit</button>"
    var deleteFunction = "<button type=\"button\" class=\"open-delEvents btn btn-danger btn-sm\" data-id=\"" + rowData.ID + "\" data-name=\"" + rowData.Name + "\" data-email=\"" + rowData.Email + "\" data-toggle=\"modal\" data-target=\"#deleteConfirmModal\">Delete</button>"

    $(document).on("click", ".open-editEvents", function() {
        var eventId = $(this).data('id');
        $('#idHolder').html(eventId);
        var eventName = $(this).data('name');
        $('#nameHolder').html(eventName);
        $('#namePlace').attr('placeholder', eventName);
        var eventEmail = $(this).data('email');
        $('#emailHolder').html(eventEmail);
        $('#emailPlace').attr('placeholder', eventEmail);
        var eventRole = $(this).data('role');
        $('#roleHolder').html(eventRole);
        $('#rolePlace').attr('placeholder', eventRole);
        var eventVerify = $(this).data('verify');
        $('#verifyHolder').html(eventVerify);
        $('#verifyPlace').attr('placeholder', eventVerify);

        $(document).unbind("click").on("click", ".confirm-edit", function() {
            var newName = $("#namePlace").val();
            if (newName == "") { newName = eventName; }
            var newEmail = $("#emailPlace").val();
            if (newEmail == "") { newEmail = eventEmail; }
            var newRole = $("#rolePlace").val();
            var newVerify = $("#verifyPlace").val();
            var request = "{" +
                "\"email\":" + "\"" + newEmail + "\"" + "," +
                "\"name\":" + "\"" + newName + "\"" + "," +
                "\"role\":" + "\"" + newRole + "\"" + "," +
                "\"verify\":" + "\"" + newVerify + "\"" +
                "}";
            $.ajax({
                url: '/admin/' + eventId,
                type: "put",
                data: request
            });
            location.reload();
        });

        $(document).on("click", ".close-edit", function() {
            $("#editModal").modal("toggle")
            location.reload();
        });

    });


    $(document).on("click", ".open-delEvents", function() {
        var eventId = $(this).data('id');
        $('#idHolder').html(eventId);
        var eventName = $(this).data('name');
        $('#nameHolder').html(eventName);
        var eventEmail = $(this).data('email');
        $('#emailHolder').html(eventEmail);

        $(document).unbind("click").on("click", ".confirm-delete", function() {
            $.ajax({
                url: '/admin/' + eventId,
                type: "delete",
            });
            location.reload();
        });
        $(document).on("click", ".close-delete", function() {
            $("#deleteConfirmModal").modal("toggle")
            location.reload();
        });
    });


    $("#personDataTable").append(row);
    row.append($("<td>" + rowData.ID + "</td>"));
    row.append($("<td>" + rowData.Name + "</td>"));
    row.append($("<td>" + rowData.Email + "</td>"));
    row.append($("<td>" + rowData.Role + "</td>"));
    row.append($("<td>" + rowData.Verify + "</td>"));
    row.append($("<td>" + editFunction + " " + deleteFunction + "</td>"));
}