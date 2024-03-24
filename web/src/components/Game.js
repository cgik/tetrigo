'use client';

import {useEffect, useMemo, useRef, useState} from "react";
import {renderCanvas} from "@/common/graphics";
import {fetchGetGame} from "@/common/game-api";

export default function Game({ id, canvasWidth, canvasHeight }) {
    const canvasRef = useRef(null);
    // const [gameId, setGameId] = useState("659e32832f2d328e97de49e3");
    let [game, setGame] = useState({});

    const gameId = useMemo(() => {
        return id;
    }, [id]);

    useEffect(() => {
        fetchGetGame(gameId).then(
            data => {
                setGame(data);
            }
        );
    }, [gameId]);

    useEffect(() => {
        const canvas = canvasRef.current;
        const ctx = canvas.getContext('2d');

        if (game.data !== undefined) {
            renderCanvas(ctx, game.data.game);
        }

    }, [game]);

    return (
        <canvas ref={canvasRef} className="bg-neutral-400" width={canvasWidth} height={canvasHeight}/>
    )
}