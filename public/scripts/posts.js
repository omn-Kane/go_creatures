function breedWith(session, creature1ID, creature2ID) {
    $.post('/breedWith?Session=' + session + "&Creature1ID=" + creature1ID + "&Creature2ID=" + creature2ID,
        function(data) {
            console.log('Yay, Success', data, data == "true")
        });
}

function setInstruction(session, season, creatureID, action) {
    $.post('/setAction?Session=' + session + "&Season=" + season + "&CreatureID=" + creatureID + "&Action=" + action,
        function(data) {
            if (data !== "false") document.getElementById("creature_action_" + creatureID).innerHTML = data;
        });
}
