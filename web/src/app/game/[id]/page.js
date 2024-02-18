import Header from "@/components/Header";
import Game from "@/components/Game";

export default function GamePage({ params }) {
    return (
        <>
            <Header/>
            <main className="flex relative justify-center border-2 border-red-500">
                <div className="border-2 border-teal-500">left</div>
                <div className="border-2 border-teal-500">
                    <Game canvasWidth={336} canvasHeight={670}/>
                </div>
                <div className="border-2 border-teal-500">right</div>
            </main>
        </>
    )
}