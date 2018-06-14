$(function () {
    $('#calculate').click(getAPI)
})

function getAPI() {
    var startDay = $('#startDay').val()
    var startMonth = $('#startMonth').val()
    var startYear = $('#startYear').val()
    var endDay = $('#endDay').val()
    var endMonth = $('#endMonth').val()
    var endYear = $('#endYear').val()
    var host = "http://localhost:3000/duration"
    var param = `?startDay=${startDay}&startMonth=${startMonth}&startYear=${startYear}` +
        `&endDay=${endDay}&endMonth=${endMonth}&endYear=${endYear}`
    var url = host + param
    $.getJSON(url, function (responseData) {
        $('#resultDay').html(responseData.days)
    })
} 