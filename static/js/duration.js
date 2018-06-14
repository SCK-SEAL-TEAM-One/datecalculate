$(function () {
    $('#calculate').click(GetApi)
})

function GetApi() {
    var host = "http://localhost:3000/duration"
    var parameter = `?startday=${$('#startDay').val()}` +
        `&startmonth=${$('#startMonth').val()}` +
        `&startyear=${$('#startYear').val()}` +
        `&endday=${$('#endDay').val()}` + 
        `&endmonth=${$('#endMonth').val()}` + 
        `&endyear=${$('#endYear').val()}`
    var url = host + parameter
    $.getJSON(url, function (responseData) {
        $('#resultDay').html(responseData.days)
    })
}