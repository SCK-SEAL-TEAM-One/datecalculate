$(function() {
    $(".datepicker").datepicker({
        dateFormat:"d/m/yy"
    })
    $('#calculateDate').click(GetDurationAPI)
})

function GetDurationAPI() {
    var hostName    = "http://localhost:3000/duration"
    var parameter   = "?startdate=" + $("#fromDate").val() + "&enddate=" + $("#toDate").val()
    var urlAPI      = hostName + parameter
    console.log(urlAPI)
    $.getJSON(urlAPI, function( data ) {
        $("#from").html("From : "+data.form)
        $("#to").html("To : "+data.to)
        $("#resultDay").html(data.days)
        $("#years").html(data.years)
        $("#seconds").html(data.seconds)
        $("#minutes").html(data.minutes)
        $("#hours").html(data.hours)
        $("#days").html(data.days)
        $("#weeks").html(data.weeks)
        $("#ratioOfYear").html(data.ratioOfYear)
    });
}