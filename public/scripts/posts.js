var api_location = 'http://localhost:8080'

function endDay(session) {
    $.post(api_location + '/endDay', JSON.stringify({"Session": session}));
}
