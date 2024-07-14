'use client'

import { fetchCreateGame } from "@/common/game-api"
import { useRouter } from 'next/navigation'

export default function CreateGameButton() {
    const router = useRouter()
    const handleClick = async () => {
        fetchCreateGame()
        .then((response) => {
            router.push(`/game?id=${response.data.game.id}`)
        })
    }

    return (
        <button className="btn btn-primary btn-wide" onClick={handleClick}>Create Game</button>
    )
}