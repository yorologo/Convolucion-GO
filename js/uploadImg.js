$(document).ready(function () {
    $('input[type="file"]').change(function (e) {
        var fileName = e.target.files[0].name;
        var fileExtension = fileName.split('.').pop()

        switch (fileExtension) {
            case 'jpg':
            case 'jpeg':
            case 'jpe':
            case 'jfif':
                $('#inputFileName').removeClass("text-danger")
                document.getElementById('inputFileName').innerHTML = fileName
                $('#btn').removeClass("disabled")
                document.getElementById('btn').disabled = false
                break;

            default:
                if (!$('#inputFileName').hasClass("text-danger")) {
                    $('#inputFileName').addClass("text-danger")
                    document.getElementById('inputFileName').innerHTML =
                        '<b>No se eligió archivo</b>'
                }
                if (!$('#btn').hasClass("disabled")) {
                    $('#btn').addClass("disabled")
                    document.getElementById('btn').disabled = true
                }
                $('#myModal').modal()
                break;
        }
    })
})