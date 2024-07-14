"use client";

import Game from "@/components/Game";
import Navigation from "@/components/Navigation";
import { useSearchParams } from "next/navigation";
import { Suspense } from "react";

export default function GamePage() {
  function GameSetup() {
    const searchParams = useSearchParams();
    let gameId = searchParams.get("id");

    if (!gameId) {
      return <div>Not sure how you got here but you were looking to create a gamne probably!</div>
    }

    return <Game id={gameId} canvasWidth={336} canvasHeight={670} />;
  }

  return (
    <div className="min-h-screen bg-base-200">
      <Navigation />

      <div className="grid grid-cols-3 gap-4 place-content-evenly h-48">
        <div>
          <h1 className="text-xl">Controls</h1>
          <p>Arrow keys: move the cursor around</p>
          <p>Space: Switch blocks</p>
        </div>
        <div>
          <Suspense fallback={<div>Loading...</div>}>
            <GameSetup />
          </Suspense>
        </div>
        <div></div>
      </div>
    </div>
  );
}
