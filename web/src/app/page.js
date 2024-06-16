"use client";

import Navigation from "@/components/Navigation";
import Game from "@/components/Game";
import Stats from "@/components/Stats";
import Script from "next/script";

// async function wasmLoad() {
//     const goWasm = new window.Go();
//     const result = await WebAssembly.instantiateStreaming(
//         fetch("static/tetrigo.wasm"),
//         goWasm.importObject
//     );
//
//     goWasm.run(result.instance);
// }

export default function Page() {
  return (
    <div className="min-h-screen bg-base-200">
      {/*<Script src="static/wasm_exec.js"*/}
      {/*        onLoad={wasmLoad}*/}
      {/*/>*/}

      <Navigation />

      <div className="flex justify-center">
        <Stats />
      </div>
    </div>
  );
}
