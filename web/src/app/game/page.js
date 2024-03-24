'use client';

import Header from "@/components/Header";
import Game from "@/components/Game";
import {useSearchParams} from "next/navigation";
import {fetchCreateGame} from "@/common/game-api";

export default function GamePage() {
    // TODO: Find game or create it and inject it into the Game component
    const searchParams = useSearchParams();
    const gameId = searchParams.get('id');

    if (!gameId) {
        fetchCreateGame().then(
            (data) => {
                console.log('Game created:', data);
            }
        );
    }

    return (
        <>
            <Header/>
            <main className="flex relative justify-center">
                <div></div>
                <div>
                    <Game id={gameId} canvasWidth={336} canvasHeight={670}/>
                </div>
                <div></div>
            </main>
        </>
    )
}