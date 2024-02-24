const apiUrlV1 = `http://127.0.0.1:18080`

export async function fetchGetGame(gameId) {
    return await fetch(`${apiUrlV1}/game/load/${gameId}`)
        .then(response => response.json())
        .then(data => {
            return data;
        })
        .catch(error => console.log(error));
}

export async function fetchCreateGame(gameId) {
    return await fetch(`${apiUrlV1}/game/create`)
        .then(response => response.json())
        .then(data => {
            return data;
        })
        .catch(error => console.log(error));
}