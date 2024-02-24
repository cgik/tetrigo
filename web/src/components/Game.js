'use client';

import {useContext, useEffect, useRef, useState} from "react";
import {drawBlock} from "@/common/graphics";
import {fetchGame} from "@/common/fetch";

function render(ctx, game) {
    ctx.clearRect(0, 0, 336, 670);
    console.log(game.data);
    // game.board.forEach(
    //     block => drawBlock(ctx, block, game)
    // )

}

export default function Game({ canvasWidth, canvasHeight }) {
    const canvasRef = useRef(null);
    const [gameId, setGameId] = useState("659e32832f2d328e97de49e3");
    let [game, setGame] = useState({});

    useEffect(() => {
        const data= fetchGame(gameId);

        setGame(data);
    }, [gameId]);

    useEffect(() => {
        const canvas = canvasRef.current;
        const ctx = canvas.getContext('2d');

        render(ctx, game)
    }, [game]);

    return (
        <canvas ref={canvasRef} className="bg-neutral-400" width={canvasWidth} height={canvasHeight}/>
    )
}
