const apiUrlV1 = `http://127.0.0.1:18080`

export async function fetchGetGame(gameId) {
    const res = await fetch(`${apiUrlV1}/game/load/${gameId}`)
    return await res.json();

}

export async function fetchCreateGame() {
    const res = await fetch(`${apiUrlV1}/game/create`);
    return await res.json();

}