'use client';

import {useEffect, useRef, useState} from "react";
import {drawBlock} from "@/common/graphics";
import {fetchGetGame} from "@/common/game-api";

function render(ctx, game) {
    ctx.clearRect(0, 0, ctx.canvas.width, ctx.canvas.height);

    game.board.blocks.forEach(
        block => drawBlock(ctx, block, game.board)
    )
}

export default function Game({ canvasWidth, canvasHeight }) {
    const canvasRef = useRef(null);
    const [gameId, setGameId] = useState("659e32832f2d328e97de49e3");
    let [game, setGame] = useState({});

    useEffect(() => {
        async function fetchData() {
            const data = await fetchGetGame(gameId);
            setGame(data);
        }
        fetchData();
    }, [gameId]);

    useEffect(() => {
        const canvas = canvasRef.current;
        const ctx = canvas.getContext('2d');

        if (game.data !== undefined) {
            render(ctx, game.data.game);
        }

    }, [game]);

    return (
        <canvas ref={canvasRef} className="bg-neutral-400" width={canvasWidth} height={canvasHeight}/>
    )
}
