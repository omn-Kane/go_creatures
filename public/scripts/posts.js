var api_location = 'http://daniel.hbonow.com:8080'

function breedWith(session, creature1ID, creature2ID) {
    $.post(api_location + '/breedWith?Session=' + session + "&Creature1ID=" + creature1ID + "&Creature2ID=" +creature2ID,
        function(data) {
            console.log('Yay, Success', data, data == "true")
        });
}
