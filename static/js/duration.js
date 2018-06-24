$(function () {
    $('#calculate').click(GetApi)
})

function GetApi() {
    var host = "http://localhost:3000/duration"
    var parameter = `?startDay=${$('#startDay').val()}` +
        `&startMonth=${$('#startMonth').val()}` +
        `&startYear=${$('#startYear').val()}` +
        `&endDay=${$('#endDay').val()}` + 
        `&endMonth=${$('#endMonth').val()}` + 
        `&endYear=${$('#endYear').val()}`
    var url = host + parameter
    $.getJSON(url, function (responseData) {
        $('#resultDay').html(responseData.days)
    })
}