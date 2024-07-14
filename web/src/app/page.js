"use client";

import Navigation from "@/components/Navigation";
import Stats from "@/components/Stats";
import CreateGameButton from "@/components/CreateGameButton";


export default function Page() {
  return (
    <div className="min-h-screen bg-base-200">
      <Navigation />

      <div className="grid grid-cols-3 gap-4 place-content-evenly h-48">
        <div>
          <h1 className="text-xl">Message from the dev</h1>
          <div className="chat chat-start">
            <div className="chat-bubble chat-bubble-primary">This is very much still in development. 7/2024</div>
          </div>
        </div>
        <div>
          <Stats />
        </div>
        <div>
          <CreateGameButton />
        </div>
      </div>

    </div>
  );
}
