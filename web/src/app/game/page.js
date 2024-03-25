'use client';

import Header from "@/components/Header";
import Game from "@/components/Game";
import {useSearchParams} from "next/navigation";
import {Suspense}  from "react";
import {fetchCreateGame} from "@/common/game-api";

export default function GamePage() {
    function GameSetup() {
        const searchParams = useSearchParams();
        const gameId = searchParams.get('id');

        if (!gameId) {
            fetchCreateGame().then(
                (data) => {
                    console.log('Game created:', data);
                }
            );
        }
        return <Game id={gameId} canvasWidth={336} canvasHeight={670}/>
    }

    return (
        <>
            <Header/>
            <main className="flex relative justify-center">
                <div></div>
                <div>
                    <Suspense fallback={<div>Loading...</div>}>
                        <GameSetup/>
                    </Suspense>
                </div>
                <div></div>
            </main>
        </>
    )
}