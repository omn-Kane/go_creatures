var api_location = 'http://daniel.hbonow.com:8080'

function breedWith(session, creature1ID, creature2ID) {
    $.post(api_location + '/breedWith?Session=' + session + "&Creature1ID=" + creature1ID + "&Creature2ID=" + creature2ID,
        function(data) {
            console.log('Yay, Success', data, data == "true")
        });
}

function setInstruction(session, creatureID, action) {
    $.post(api_location + '/setAction?Session=' + session + "&CreatureID=" + creatureID + "&Action=" + action,
        function(data) {
            if (data !== "false") document.getElementById("creature_action_" + creatureID).innerHTML = data;
        });
}
