const apiUrlV1 = process.env.NODE_ENV == "production" ? `/api` : `http://localhost:18080`

export async function fetchGetGame(gameId) {
    const res = await fetch(`${apiUrlV1}/game/load/${gameId}`)
    return await res.json();

}

export async function fetchCreateGame() {
    const res = await fetch(`${apiUrlV1}/game/create`);
    return await res.json();

}