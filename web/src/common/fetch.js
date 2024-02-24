export async function fetchGame(gameId) {
    const response = await fetch(`http://127.0.0.1:18080/game/load/${gameId}`);

    if (!response.ok) {
        throw new Error(response.statusText);
    }

    return response.json();
}