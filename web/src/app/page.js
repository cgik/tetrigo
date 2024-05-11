'use client';

import Header from '@/components/Header'
import Game from '@/components/Game'
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
        <>
            {/*<Script src="static/wasm_exec.js"*/}
            {/*        onLoad={wasmLoad}*/}
            {/*/>*/}

            <Header/>
            <p>Home</p>
        </>
    )
}
