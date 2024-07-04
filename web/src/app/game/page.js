"use client";

import Game from "@/components/Game";
import Navigation from "@/components/Navigation";
import { useSearchParams } from "next/navigation";
import { Suspense } from "react";
import { fetchCreateGame } from "@/common/game-api";

export default function GamePage() {
  function GameSetup() {
    const searchParams = useSearchParams();
    const gameId = searchParams.get("id");

    if (!gameId) {
      fetchCreateGame().then((data) => {
        console.log("Game created:", data);
      });
    }
    return <Game id={gameId} canvasWidth={336} canvasHeight={670} />;
  }

  return (
    <div className="min-h-screen bg-base-200">
      <Navigation />
      <div className="flex justify-center">
        <Suspense fallback={<div>Loading...</div>}>
          <GameSetup />
        </Suspense>
      </div>
    </div>
  );
}
