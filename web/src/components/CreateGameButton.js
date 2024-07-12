'use client'

import { fetchCreateGame } from "@/common/game-api"
import { useRouter } from 'next/navigation'
import { useEffect } from "react"

export default function CreateGameButton() {
    const router = useRouter()
    const handleClick = async () => {
        fetchCreateGame()
        .then((response) => {
            router.push(`/game?id=${response.data.game.id}`)
        })
    }

    return (
        <button className="btn btn-wide" onClick={handleClick}>Create Game</button>
    )
}